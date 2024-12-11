import {
  addProjectConfiguration,
  getProjects,
  JsonParseOptions,
  JsonSerializeOptions,
  ProjectConfiguration,
  Tree,
  updateJson,
  updateProjectConfiguration,
  writeJson,
} from '@nx/devkit';
import { downloadBuffer } from './fs';
import { sortObjectKeys } from './sort';
import { join } from 'path';

export const upsertProjectConfiguration = (
  tree: Tree,
  project: string,
  configuration: Partial<ProjectConfiguration>
) => {
  const projectExists = getProjects(tree).has(project);
  const action = projectExists
    ? updateProjectConfiguration
    : addProjectConfiguration;

  const sortedConfiguration = sortObjectKeys(
    {
      ...configuration,
      targets: sortObjectKeys(configuration.targets, nxProjectTargets),
    },
    nxProjectGroups
  ) as ProjectConfiguration;
  return action(tree, project, sortedConfiguration);
};

export const upsertJson = <T extends object, U extends object = T>(
  tree: Tree,
  path: string,
  updater: (value: T) => U,
  options?: JsonParseOptions & JsonSerializeOptions
): void => {
  const exists = tree.exists(path);
  if (exists) {
    return updateJson(tree, path, updater, options);
  }
  const data = updater(null);
  writeJson(tree, path, data, options);
};

export interface ProjectExpansion {
  projectName: string;
  projectRoot: string;
  projectExists: boolean;
  projectConfig: ProjectConfiguration;
}

export const expandProjectInput = (
  tree: Tree,
  projectNameOrPath: string
): ProjectExpansion => {
  try {
    const projectConfig = getProjectConfig(tree, projectNameOrPath);
    const { root: projectRoot } = projectConfig;
    return {
      projectName: projectNameOrPath,
      projectRoot,
      projectExists: true,
      projectConfig,
    };
  } catch (err) {
    const projectRoot = projectNameOrPath;
    const projectName = projectPathToProjectName(projectNameOrPath);
    return {
      projectName,
      projectRoot,
      projectExists: false,
      projectConfig: null,
    };
  }
};

export const projectPathToProjectName = (projectPath: string) => {
  return projectPath
    .replace(/^(libs|apps|tools\/(generators|executors))\//gi, '')
    .replace(/\//gi, '-');
};

export const nxProjectGroups = [
  '$schema',
  'projectType',
  'sourceRoot',
  'implicitDependencies',
  'targets',
  'tags',
];

export const nxProjectTargets = [
  'breaking',
  'generate',
  'build',
  'build-storybook',
  'pre-build',
  'serve',
  'storybook',
  'test',
  'test-storybook',
  'e2e',
  'eslint',
  'lint',
  'format',
  'clean',
  'container',
  'package',
  'publish',
  'chromatic',
  'prepare',
  'version',
];

export const getProjectAssetsRoot = (tree: Tree, projectName: string) => {
  const { root, targets } = getProjectConfig(tree, projectName);
  let assetsRoot = null;

  switch (targets?.build?.executor) {
    case '@nx/webpack:webpack':
    case '@nx/esbuild:esbuild':
    case '@nx/node:webpack': {
      assetsRoot = `${root}/src/assets`;
      break;
    }
    case '@nx-dotnet/core:build': {
      assetsRoot = `${root}/Assets`;
      break;
    }
    case '@nxrocks/nx-spring-boot:build': {
      assetsRoot = `${root}/src/main/resources`;
      break;
    }
    case '@nx-go/nx-go:build': {
      assetsRoot = `${root}/assets`;
      break;
    }
    case 'nx:run-commands': {
      if (targets.build?.options?.command?.startsWith('go')) {
        assetsRoot = `${root}/assets`;
        break;
      }
      if (targets.build?.options?.command?.includes('gradle')) {
        assetsRoot = root;
        break;
      }
    }
    // eslint-disable-next-line no-fallthrough
    default: {
      throw new Error(
        `Unable to get asset root for ${root}. Unimplemented executor ${targets?.build?.executor}`
      );
    }
  }
  return assetsRoot;
};

export const getProjectConfig = (tree: Tree, projectName: string) => {
  const projectConfigs = getProjects(tree);

  const hasConfig = projectConfigs.has(projectName);
  if (!hasConfig) {
    throw new Error(`Unable to get project config for ${projectName}.`);
  }
  return projectConfigs.get(projectName);
};

export const getProjectSrcRoot = (tree: Tree, projectName: string) => {
  // NOTE : This value is stored on the project config so this might be duplicate.
  const { root, targets } = getProjectConfig(tree, projectName);
  let srcRoot = null;

  switch (targets?.build?.executor) {
    case '@nx/webpack:webpack':
    case '@nx/node:webpack': {
      srcRoot = `${root}/src`;
      break;
    }
    case '@nx-go/nx-go:build':
    case '@nx-dotnet/core:build': {
      srcRoot = `${root}`;
      break;
    }
    case '@nxrocks/nx-spring-boot:build': {
      srcRoot = `${root}/src`;
      break;
    }
    default: {
      throw new Error(
        `Unable to get src root for ${root}. Unimplemented executor ${targets?.build?.executor}`
      );
    }
  }
  return srcRoot;
};

export const getProjectLocalRoot = (tree: Tree, projectName: string) => {
  const { targets, root } = getProjectConfig(tree, projectName);
  let localRoot = null;

  switch (targets?.build?.executor) {
    case '@nxrocks/nx-spring-boot:build':
    case '@nx/esbuild:esbuild':
    case '@nx/webpack:webpack':
    case '@nx/node:webpack':
    case '@nx-go/nx-go:build':
    case '@nx-dotnet/core:build': {
      localRoot = `${root}/local`;
      break;
    }
    case 'nx:run-commands': {
      if (targets.build?.options?.command?.startsWith('go')) {
        localRoot = `${root}/local`;
        break;
      }
      if (targets.build?.options?.command?.includes('gradle')) {
        localRoot = root;
        break;
      }
    }
    // eslint-disable-next-line no-fallthrough
    default: {
      throw new Error(
        `Unable to get local root for ${root}. Unimplemented executor ${targets?.build?.executor}`
      );
    }
  }
  return localRoot;
};

export const terraformModulesRoot = 'infrastructure/modules/terraform';

export const getProjectTerraformRoot = (tree: Tree, projectName: string) => {
  const { root } = getProjectConfig(tree, projectName);
  return join(terraformModulesRoot, 'platform', root.replace('apps/', ''));
};

export const pathToUnderscore = (path: string): string => {
  return path.replace(/\//gi, '_');
};

export const getTags = (config: ProjectConfiguration) => {
  const { appType, api, language, framework, system } =
    decodeProjectConfiguration(config);
  const isTest = config.name.endsWith('-test');

  return [
    `platform:platform`,
    appType ? `app:type:${appType}` : null,
    system ? `system:${system}` : null,
    api?.visibility ? `api:visibility:${api.visibility}` : null,
    api?.version ? `api:version:${api.version}` : null,
    language ? `language:${language}` : null,
    framework ? `framework:${framework}` : null,
  ].filter((v) => !!v);
};

export const decodeProjectConfiguration = (config: ProjectConfiguration) => {
  const [, subType, ...dirs] = config.root.split('/');
  let framework: string = frameworks.some((f) => f === subType)
    ? subType
    : null;
  let language: string = null;
  let appType = null;

  switch (config?.targets?.build?.executor) {
    case '@nxrocks/nx-spring-boot:build':
      language = 'java';
      framework = 'spring-boot';
      break;
    case '@nx/esbuild:esbuild':
    case '@nx/webpack:webpack':
    case '@nx/node:webpack':
      language = 'typeScript';
      break;
    case '@nx-go/nx-go:build':
      language = 'golang';
      break;
    case '@nx-dotnet/core:build':
      language = 'c#';
      framework = 'dotnet';
      break;
    case 'nx:run-commands': {
      if (config?.targets.build?.options?.command?.startsWith('go')) {
        language = 'golang';
        break;
      }
      if (config?.targets.build?.options?.command?.includes('gradle')) {
        language = 'java';
        break;
      }
    }
  }

  switch (config.projectType) {
    case 'library':
      break;
    case 'application':
      if (config.name.endsWith('-test')) {
        appType = 'test';
        break;
      }
      appType = subType;
      break;
  }

  const api =
    appType === 'workloads'
      ? {
          visibility: dirs[0],
          version: dirs[2],
        }
      : null;

  const system = appType === 'workloads' ? dirs[1] : null;

  return {
    projectType: config.projectType,
    appType,
    language,
    framework,
    api,
    system,
  };
};

/**
 * Downloads a file from a URL to a tree and returns the output path.
 * See downloadBuffer for more details.
 *
 * @param tree The tree to write the file to
 * @param url The URL to download the file from
 * @param path The directory to download the file to within the tree
 * @param fileName The optional name (without extension) of the file to place in the outputPath
 * @param fileExtension The optional extension of the file to place in the outputPath.
 * @returns The path to the downloaded file within the tree
 * @throws If the download fails
 */
export async function download(
  tree: Tree,
  url: string,
  path: string,
  fileName?: string,
  fileExtension?: string
): Promise<string> {
  const [buffer, outputPath] = await downloadBuffer(
    url,
    path,
    fileName,
    fileExtension
  );

  await tree.write(outputPath, buffer);
  return outputPath;
}

export const directories = {
  angular: 'web',
  react: 'web',
  microservice: 'workloads',
  plugin: 'plugins',
} as const;

export const languages = [
  'c#',
  'golang',
  'java',
  'typeScript',
  'python',
  'swift',
  'rust',
  'objectiveC',
] as const;
export const frameworks = [
  'dotnet',
  'go',
  'node',
  'java',
  'python',
  'swift',
] as const;

export type framework = keyof typeof frameworks;
export type language = keyof typeof languages;

export type Directory = keyof typeof directories;
