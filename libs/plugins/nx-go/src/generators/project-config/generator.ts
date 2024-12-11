import {
  formatFiles,
  getProjects,
  Tree,
  ProjectConfiguration,
} from '@nx/devkit';
import { ProjectConfigGeneratorSchema } from './schema';
import { getContainerImageName } from '@platform/nx-container';
import {
  getTags,
  projectPathToProjectName,
  upsertProjectConfiguration,
} from '@platform/nx-devkit';
import { addProjectConfiguration } from 'nx/src/generators/utils/project-configuration';

export async function projectConfigGenerator(
  tree: Tree,
  options: ProjectConfigGeneratorSchema
) {
  const projectName = options.create
    ? projectPathToProjectName(options.project)
    : options.project;

  if (options.create) {
    if (!options.project.includes('/')) {
      throw new Error(
        'project must be specified in the form of <directory>/<name> when using create'
      );
    }

    const projectType = options.project.startsWith('libs')
      ? 'library'
      : 'application';
    addProjectConfiguration(
      tree,
      options.project,
      {
        name: projectName,
        root: options.project,
        sourceRoot: options.project,
        projectType,
      },
      true
    );
  }

  const config = getProjects(tree).get(projectName);

  if (!config) {
    throw new Error("Couldn't find project in workspace");
  }

  const updatedTargets = getConfigTargets(config);

  const defaultTags = getTags(config) ?? [];
  const existingUniqueTags =
    config.tags ??
    [].filter(
      (t) =>
        !defaultTags.some((dt) => {
          const [, ...parts] = dt.split(':').reverse();
          const prefix = parts.reverse().join(':') + ':';
          return t.startsWith(prefix);
        })
    );

  upsertProjectConfiguration(tree, projectName, {
    ...config,
    targets: { ...config.targets, ...updatedTargets },
    tags: [...defaultTags, ...(options.keepTags ? existingUniqueTags : [])],
  });

  await formatFiles(tree);
}
export const getFormatExecutorConfig = (projectRoot: string) => ({
  executor: 'nx:run-commands',
  options: {
    command: `./gradlew :${getApplicationPath(projectRoot)}:spotlessApply`,
    cwd: '',
  },
});
export const getContainerExecutorConfig = (projectRoot: string) => ({
  executor: 'nx:run-commands',
  options: {
    command: `./gradlew :${getApplicationPath(
      projectRoot
    )}:bootImage --imageName=${getContainerImageName(projectRoot)}`,
    cwd: '',
  },
});
export const getVersionExecutorConfig = (projectRoot: string) => ({
  outputs: ['{options.packageJson}'],
  executor: '@oeco/nx-version:bump',
  options: {
    packageJson: `${projectRoot}/package.json`,
  },
});

export const getContainerPublishExecutorConfig = (projectRoot: string) => ({
  executor: '@oeco/nx-container:push',
  options: {
    image: getContainerImageName(projectRoot),
    version: {
      path: `${projectRoot}/package.json`,
      key: 'version',
    },
    registries: ['platform-operations.registry.cpln.io', 'ghcr.io/platform'],
  },
});

export const getApplicationPath = (projectRoot: string) => {
  return projectRoot.replace(/\//g, ':');
};
export const GetBuildExecutorConfiguration = (projectRoot: string) => ({
  executor: 'nx:run-commands',
  options: {
    command: `./gradlew :${getApplicationPath(projectRoot)}:compileJava`,
    cwd: '',
  },
});
export const GetServeExecutorConfig = (projectRoot: string) => ({
  executor: 'nx:run-commands',
  options: {
    command: `./gradlew :${getApplicationPath(projectRoot)}:bootRun`,
    cwd: '',
  },
});
export const GetCleanExecutorConfig = (projectRoot: string) => ({
  executor: 'nx:run-commands',
  options: {
    command: `./gradlew :${getApplicationPath(projectRoot)}:clean`,
    cwd: '',
  },
});
export const GetLintExecutorConfiguration = (projectRoot: string) => ({
  executor: '@nx/eslint:eslint',
  outputs: ['{options.outputFile}'],
  options: {
    lintFilePatterns: [`${projectRoot}/**/*.yaml`],
  },
});

export const getConfigTargets = (config: ProjectConfiguration) => {
  let typeSpecificTasks = {};
  const isTest = config.name.endsWith('-test');
  switch (config.projectType) {
    case 'application': {
      typeSpecificTasks = {
        clean: GetCleanExecutorConfig(config.root),
        serve: GetServeExecutorConfig(config.root),
        container: getContainerExecutorConfig(config.root),
        publish: getContainerPublishExecutorConfig(config.root),
      };
      break;
    }
    case 'library': {
      typeSpecificTasks = {
        clean: GetCleanExecutorConfig(config.root),
      };
      break;
    }
  }
  return {
    build: GetBuildExecutorConfiguration(config.root),
    lint: GetLintExecutorConfiguration(config.root),
    format: getFormatExecutorConfig(config.root),
    version: getVersionExecutorConfig(config.root),
    ...typeSpecificTasks,
  };
};

export default projectConfigGenerator;
