import { dasherize } from '@nx/devkit/src/utils/string-utils';
import { Tree } from '@nx/devkit';
import type * as prompts from 'prompts';
import {
  getProtoDirectories,
  mapProtoRelativePath,
  protoByVersion,
  protoFromRelativePath,
  getProtoVersions,
  getProjectNames,
  protoFromProjectRoot,
  ProtoFileDetail,
} from './tree';
import { directories, getProjectConfig } from './nx';
import { dynamicPrompt, NormalizedSchema, normalizeOptions } from './prompt';
import {
  baseProtoDirectoryName,
  isOption,
  isSafeOption,
  isService,
  readProto,
  safeProtoDirToApiVisibility,
} from './proto';
import { PrefixCamelCaseKeys } from './types';
import * as path from 'path';
import { getRpcServices, SafeOptionsRpcService } from './proto-method';

export interface BaseProtoSchema {
  directory: string;
  version: string;
  relativePath: string;
  path: string;
  api: string;
  subApi: string;
  fileName: string;
}

export interface BaseMicroserviceSchema {
  name: string;
  suffix: string;
  directory: string;
  computedName: string;
}

export interface MicroServiceDetails {
  serviceName: string;
  rpcServices: Array<SafeOptionsRpcService>;
}

export type SelectFromProtoOverrides = PrefixCamelCaseKeys<
  BaseProtoSchema,
  'proto'
> &
  BaseMicroserviceSchema;

export type SelectFromProtoSchema = SelectFromProtoOverrides & NormalizedSchema;

export interface SelectProtoFromProjectName extends ProtoFileDetail {
  root: string;
  project: string;
}

export const getProtoFromProject = async (
  tree: Tree,
  prompts: prompts,
  overrides,
  dynamic
): Promise<SelectProtoFromProjectName> => {
  const projectNames = overrides?.project
    ? []
    : getProjectNames(tree, 'microservice');
  const { project } = await dynamicPrompt(
    prompts,
    dynamic,
    'project',
    projectNames
  );

  let { root } = getProjectConfig(tree, project);
  const { tags } = getProjectConfig(tree, project);
  if (tags !== undefined) {
    const result = tags.filter((tag) => tag.startsWith('multi-proto'));

    if (result.length > 0) {
      const pair = result[0].split(':');
      if (pair.length != 2) {
        throw new Error(
          `Use the proto tag only for multi-proto use cases. Also, pattern is multi-proto:{name-of-proto-file}`
        );
      }
      const p = pair[1];
      root = path.dirname(root) + '/' + p;
    }
  }

  const protoByFile = protoFromProjectRoot(tree, root);

  return { ...protoByFile, root, project };
};

export const selectFromProto = async (
  tree: Tree,
  prompts: prompts,
  overrides: Partial<SelectFromProtoOverrides>,
  dynamic,
  protoDirectoryPrefix = baseProtoDirectoryName
): Promise<SelectFromProtoSchema> => {
  // TODO : Support override of protoPath, if set then derive overrides for protoDirectory, protoVersion, protoRelativePath
  const safeProtoDirectories = overrides?.protoDirectory
    ? []
    : getProtoDirectories(tree, (d) => d.startsWith(protoDirectoryPrefix));

  const { protoDirectory } = await dynamicPrompt(
    prompts,
    dynamic,
    'protoDirectory',
    safeProtoDirectories
  );

  const protoVersions = overrides?.protoVersion
    ? []
    : getProtoVersions(tree, protoDirectory);
  const { protoVersion } = await dynamicPrompt(
    prompts,
    dynamic,
    'protoVersion',
    protoVersions
  );

  const protosByVersion =
    overrides?.protoApi && overrides?.protoSubApi && overrides?.protoPath
      ? []
      : protoByVersion(tree, protoDirectory, protoVersion);
  const { protoRelativePath } = await dynamicPrompt(
    prompts,
    dynamic,
    'protoRelativePath',
    protosByVersion.map(mapProtoRelativePath)
  );

  const allowUserInput =
    dynamic.properties.protoRelativePath['x-prompt']?.allowUserInput || false;

  const protoByFile = {
    ...(protosByVersion.find(
      (p) => p.protoRelativePath === protoRelativePath
    ) || allowUserInput
      ? protoFromRelativePath(protoRelativePath, protoDirectory)
      : {}),
  };

  const {
    protoApi,
    protoSubApi,
    protoFileName,
    protoPath,
    protoRelativeDirectory,
  } = {
    ...{
      protoApi: protoByFile.api,
      protoSubApi: protoByFile.subApi,
      protoFileName: protoByFile.fileName,
      protoPath: protoByFile.protoPath,
      protoRelativeDirectory: protoByFile.protoRelativeDirectory,
    },
    ...overrides,
  };

  if (!(protoApi && protoSubApi && protoPath)) {
    throw new Error(
      `Unable to find path info from proto file path: ${protoRelativePath}`
    );
  }

  const name = dasherize(protoSubApi);
  const [, directorySuffix] = protoDirectory.split('-');
  const parentDirectory = directorySuffix
    ? `${directories.microservice}-${directorySuffix}`
    : directories.microservice;
  const directory = dasherize(
    `${parentDirectory}/${safeProtoDirToApiVisibility[protoDirectory]}/${protoRelativeDirectory}`
  );
  const hasMultipleImplementations = readHasMultipleImplementationsFromProto(
    tree,
    protoPath
  );
  const { suffix } = overrides?.suffix
    ? overrides
    : hasMultipleImplementations
    ? await dynamicPrompt(prompts, dynamic, 'suffix')
    : { suffix: null };

  const computedName = suffix ? `${name}-${suffix}` : name;
  const normalizedOptions = normalizeOptions(tree, {
    name: computedName,
    directory,
    protoRelativeDirectory,
  });

  return {
    protoDirectory,
    protoVersion,
    protoRelativePath,
    protoPath,
    protoApi,
    protoSubApi,
    protoFileName,

    name,
    suffix,
    directory,
    computedName,

    ...normalizedOptions,
  };
};

export const readHasMultipleImplementationsFromProto = (
  tree: Tree,
  protoPath: string
) => {
  if (!tree.exists(protoPath)) {
    return false;
  }

  const {
    ast: { statements },
  } = readProto(tree, protoPath);

  const options = statements.filter(isOption);
  const safeOptions = options.filter(
    isSafeOption(/\.has_multiple_implementations$/)
  );
  if (safeOptions.length !== 1) {
    return false;
  }

  const [safeOption] = safeOptions;
  const protoText = tree.read(protoPath)?.toString();
  const optionText = protoText.substring(safeOption.start, safeOption.end);

  return optionText.endsWith('true;');
};

export const getPorts = () => {
  const grpcPort = '50000';
  const httpPort = '8080';
  return { grpcPort, httpPort };
};

export const readServicesFromProto = (
  tree: Tree,
  protoPath: string
): Array<MicroServiceDetails> => {
  const {
    ast: { statements },
  } = readProto(tree, protoPath);

  const services = statements.filter(isService);

  return services.map(
    ({
      serviceName: { text: serviceName },
      serviceBody: { statements: serviceStatements },
    }) => {
      const protoText = tree.read(protoPath)?.toString();
      const options = serviceStatements.filter(isOption);
      const safeOptions = options.filter(
        isSafeOption(/v\d+[a-zA-Z]*?\.service$/)
      );
      const rpcServices = getRpcServices(serviceStatements, protoText);
      if (!safeOptions?.length) {
        console.warn(`[WARN] Missing option: (platform.options.*.service)`);
        return { serviceName: serviceName, rpcServices: rpcServices };
      }

      // TODO : process values from all safeOptions
      // const [safeOption] = safeOptions;
      // const protoText = tree.read(protoPath)?.toString();
      // const optionText = protoText.substring(safeOption.start, safeOption.end);

      return { serviceName: serviceName, rpcServices: rpcServices };
    }
  );
};
