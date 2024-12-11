import { query } from '@phenomnomnominal/tsquery/dist/src/query';
import { getProjects, getWorkspaceLayout, Tree } from '@nx/devkit';
import { dasherize, underscore } from '@nx/devkit/src/utils/string-utils';
import { join } from 'path';
import { directories } from './nx';
import { KeysWithValueType } from './types';
import { apiVisibilityToSafeProtoDir, serviceToProtoName } from './proto';

export const libsPluginsRoot = 'libs/plugins';
export const protoRoot = 'proto';

const languagesPath = `${libsPluginsRoot}/protoc-gen-paltform/languages`;

export const readJsonBom = <T extends object = any>(
  tree: Tree,
  path: string
): T => {
  const fileContents = tree.read(path).toString();
  const contentsWithoutBom = fileContents.replace(/^\ufeff/, '');
  return JSON.parse(contentsWithoutBom);
};

export const writeJsonBom = <T extends object = object>(
  tree: Tree,
  path: string,
  value: T
): void => {
  const fileContents = JSON.stringify(value, null, 2);
  const bom = /^\ufeff/.test(fileContents) ? '\ufeff' : '';
  tree.write(path, `${bom}${fileContents}`);
};

export const updateJsonBom = async <
  T extends object = any,
  U extends object = T
>(
  tree: Tree,
  path: string,
  updater: (value: T) => U | Promise<U>
): Promise<void> => {
  const value = readJsonBom(tree, path);
  const updatedValue = await updater(value);
  writeJsonBom(tree, path, updatedValue);
};

export const getProtocGenSafeLanguages = (tree: Tree) => {
  return tree.children(languagesPath);
};

export const getProtocGenSafePlugins = (tree: Tree) => {
  return tree
    .children(languagesPath)
    .flatMap((language) =>
      tree
        .children(`${languagesPath}/${language}/plugins`)
        .map((type) => `${language}/${type}`)
    )
    .map(dasherize);
};

export const getProtocGenPlugins = (tree) => {
  const languages = getProtocGenSafeLanguages(tree);
  const plugins = getProtocGenSafePlugins(tree);
  return [...languages, ...plugins.map((p) => `platform/${p}`)];
};

export const getProtocGenSafeLanguageExt = (tree: Tree, language: string) => {
  const languageFile = tree.read(
    `${languagesPath}/${language}/language_${language}.go`
  );
  const languageText = languageFile?.toString();
  const [matchingNode] = query(
    languageText,
    'ExpressionStatement:has(Identifier[name=FileExtension]) + Block > ReturnStatement > StringLiteral'
  );
  return matchingNode.getText().replace(/"/gi, '');
};

export const getProtoVersions = (tree, protoDirectory) => {
  const apisPath = `${protoRoot}/${protoDirectory}/platform`;
  const safeApiVersions = tree
    .children(apisPath)
    .flatMap((a) =>
      tree
        .children(`${apisPath}/${a}`)
        .filter((k) => !tree.isFile(`${apisPath}/${a}/${k}`))
    )
    .filter((v) => v.startsWith('v'));
  return [...new Set(safeApiVersions)];
};

export interface ProtoFileDetail {
  protoRoot: string;
  directory: string;
  packageName: string;
  api: string;
  subApi: string;
  version: string;
  fileName: string;
  protoRelativePath: string;
  protoRelativeDirectory: string;
  protoPath: string;
}

export const protoByVersion = (
  tree: Tree,
  directory: string,
  version: string,
  packageName = 'platform'
): Array<ProtoFileDetail> => {
  const apisPath = `${protoRoot}/${directory}/${packageName}`;
  return tree.children(apisPath).flatMap((api) => {
    const apiPath = `${apisPath}/${api}/${version}`;
    return tree
      .children(apiPath)
      .filter((fileName) => fileName.endsWith('.proto'))
      .map((fileName) => {
        const subApi = fileName.replace('.proto', '');
        const protoRelativePath = `${api}/${version}/${fileName}`;
        const protoPath = `${apiPath}/${fileName}`;
        const protoRelativeDirectory = `${api}/${version}`;
        // TODO : remove proto prefix on root, relativePath, & path???
        return {
          protoRoot,
          directory,
          packageName,
          api,
          subApi,
          version,
          fileName,
          protoRelativePath,
          protoPath,
          protoRelativeDirectory,
        };
      });
  });
};

export const protoFromProjectRoot = (
  tree: Tree,
  root: string,
  packageName = 'platform'
): ProtoFileDetail => {
  const { appsDir } = getWorkspaceLayout(tree);
  const appsMicroservicePath = join(appsDir, directories.microservice);
  if (!root.startsWith(appsMicroservicePath)) {
    throw new Error('Unsupported project path. Must be a microservice.');
  }
  const [apiVisibility, system, version, ...serviceDirs] = root
    .replace(appsDir + '/', '')
    .replace(directories.microservice + '/', '')
    .split('/');

  const [service, ...serviceParentDirs] = serviceDirs.reverse();
  const directory = apiVisibilityToSafeProtoDir[apiVisibility];
  const parentDirectory = `proto/${directory}/${packageName}`;
  const protoName = serviceToProtoName(service);
  const protoDirs = underscore(
    `${system}/${version}${serviceParentDirs?.join('/') ?? ''}`
  );
  const relativePath = `${protoDirs}/${protoName}`;
  const protoPath = `${parentDirectory}/${relativePath}`;

  if (!tree.exists(protoPath)) {
    throw new Error(
      `No proto file in ${parentDirectory} matched version: ${version}, fileName: ${protoName}`
    );
  }

  const protoFile = protoFromRelativePath(relativePath, directory);

  return { ...protoFile };
};

export const protoFromRelativePath = (
  protoRelativePath: string,
  directory: string,
  packageName = 'platform'
): ProtoFileDetail => {
  if (!protoRelativePath.endsWith('.proto')) {
    throw new Error('invalid proto extension');
  }
  // TODO : add support for folder pre nested proto file. Example foo/bar/audit/v2/admin/audit_admin.proto
  const [api, version, subApi] = underscore(
    protoRelativePath.replace('.proto', '')
  ).split('/');
  const fileName = `${subApi}.proto`;
  const protoPath = `${protoRoot}/${directory}/${packageName}/${api}/${version}/${fileName}`;
  const protoRelativeDirectory = `${api}/${version}`;
  return {
    protoRoot,
    directory,
    packageName,
    api,
    subApi,
    version,
    fileName,
    protoRelativePath,
    protoPath,
    protoRelativeDirectory,
  };
};

export const mapProtoRelativePath = ({ protoRelativePath }) =>
  protoRelativePath;

export const getProtoRoots = (tree: Tree) => {
  return tree
    .children(protoRoot)
    .flatMap((parent) =>
      tree
        .children(`${protoRoot}/${parent}`)
        .filter((child) => !tree.isFile(`${protoRoot}/${parent}/${child}`))
        .map((child) => `${parent}/${child}`)
    )
    .reverse();
};

export const getProtoDirectories = (tree: Tree, filter) => {
  return tree
    .children(protoRoot)
    .filter((d) => !tree.isFile(`${protoRoot}/${d}`))
    .filter(filter);
};

export const getMatchingFilePaths = (
  tree: Tree,
  directory: string,
  filter: (path: string) => boolean,
  filePaths = []
) => {
  if (tree.exists(directory)) {
    for (const entry of tree.children(directory)) {
      const entryPath = join(directory, entry);
      if (!tree.isFile(entryPath)) {
        getMatchingFilePaths(tree, entryPath, filter, filePaths);
      } else {
        if (filter(entryPath)) {
          filePaths.push(entryPath);
        }
      }
    }
  }

  return filePaths;
};

export const getProjectNames = (
  tree: Tree,
  type?: KeysWithValueType<typeof directories, string>
) => {
  const projectConfigs = getProjects(tree);
  const projectNames = Array.from(projectConfigs.keys());
  return type
    ? projectNames.filter((k) => k?.startsWith(directories[type]))
    : projectNames;
};

export const treeExtensions: { [key: string]: (tree: Tree) => Array<unknown> } =
  {
    getProtocGenSafeLanguages,
    getProtoRoots,
  };

export type treeExtension = keyof typeof treeExtensions;
