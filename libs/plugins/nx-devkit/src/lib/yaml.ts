import { Tree } from '@nx/devkit';
import { readFile } from 'fs/promises';
import { parse, stringify } from 'yaml';
import {
  CreateNodeOptions,
  DocumentOptions,
  ParseOptions,
  SchemaOptions,
  ToStringOptions,
} from 'yaml/dist/options';

export const readYaml = (tree: Tree, filePath: string) => {
  const yamlText = tree.read(filePath)?.toString();
  return parse(yamlText);
};

export const readYamlFile = async (filePath: string) => {
  const yamlBuffer = await readFile(filePath);
  return parse(yamlBuffer.toString());
};

export type StringifyYamlOptions = DocumentOptions &
  SchemaOptions &
  ParseOptions &
  CreateNodeOptions &
  ToStringOptions;

export const writeYaml = (
  tree: Tree,
  filePath: string,
  value: unknown,
  options?: StringifyYamlOptions,
  replace?: (text: string) => string
) => {
  let yamlText = stringify(value, options);
  if (replace) {
    yamlText = replace(yamlText);
  }
  tree.write(filePath, yamlText);
};
