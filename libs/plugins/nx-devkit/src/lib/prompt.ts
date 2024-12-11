import { treeExtensions } from './tree';
import { getWorkspaceLayout, names, readJson, Tree } from '@nx/devkit';
import { PrefixCamelCaseKeys } from './types';
import { dasherize } from '@nx/devkit/src/utils/string-utils';
import { dirname } from 'path';

export const toChoices = (value) => ({ value });

export const toChoicesTitle = (title) => ({ title });

export const allowUserInput = function () {
  this.fallback = {
    title: this.input,
    description: `Executes ${this.input}`,
    value: this.input,
  };
  if (this.suggestions.length === 0) {
    this.value = this.fallback.value;
  }
};

export const onCancel = () => {
  console.log('Canceled by user, exiting...');
  process.exit();
  throw new Error('Canceled by user, exiting...');
};

export const dynamicPromptFromTree = async (
  tree,
  prompts,
  { properties, required },
  key
) => {
  const { extensionName } = properties[key]['x-prompt'];
  const value = treeExtensions[extensionName](tree);
  return dynamicPrompt(prompts, { properties, required }, key, value);
};

export const dynamicPrompt = async (
  prompts,
  { properties, required }: any,
  key,
  values: Array<unknown> = []
) => {
  const xPrompt = properties[key]['x-prompt'];

  if (values?.length === 1 && !xPrompt.preventDefault) {
    return { [key]: values[0] };
  }

  const options = {
    ...xPrompt,
    name: key,
    choices:
      xPrompt.type === 'autocomplete'
        ? values.map(toChoicesTitle)
        : xPrompt?.choices,
    onState: xPrompt?.allowUserInput ? allowUserInput : undefined,
    initial:
      values?.length === 1 && xPrompt.preventDefault ? values[0] : undefined,
  };

  const response = await prompts(options);
  if (required?.includes(key) && !response[key]) {
    throw new Error(`Missing required argument: ${key}`);
  }
  return response;
};

export const readSchema = <T extends object = { dynamic?: object }>(
  tree,
  dirName: string
): T => {
  const libsDir = getWorkspaceLayout(tree).libsDir;
  const pattern = new RegExp('^.*?(' + libsDir + '/plugins/.*)', 'gi');
  const path = `${dirName.replace(pattern, '$1')}/schema.json`;
  return readJson<T>(tree, path);
};

export interface BaseProjectSchema {
  name: string;
  root: string;
  directory: string;
}

export interface BaseGeneratorSchema {
  name: string;
  tags?: string;
  directory?: string;
  protoRelativeDirectory?: string;
}

export type NormalizedSchema<T = BaseGeneratorSchema> = T &
  PrefixCamelCaseKeys<BaseProjectSchema, 'project'> & {
    parsedTags: string[];
  };

export const normalizeOptions = (
  tree: Tree,
  options: BaseGeneratorSchema,
  type: 'application' | 'library' = 'application'
): NormalizedSchema => {
  const name = names(options.name).fileName;
  const projectPath = dirname(options.protoRelativeDirectory);
  const projectDirectory = options.directory
    ? `${names(options.directory).name}/${name}`
    : projectPath;
  const projectName = dasherize(projectDirectory).replace(
    new RegExp('/', 'g'),
    '-'
  );
  const w = getWorkspaceLayout(tree);
  const projectRoot = `${
    w[type === 'application' ? 'appsDir' : 'libsDir']
  }/${projectDirectory}`;
  const parsedTags = options.tags
    ? options.tags.split(',').map((s) => s.trim())
    : [];

  return {
    ...options,
    projectName,
    projectRoot,
    projectDirectory,
    parsedTags,
  };
};
