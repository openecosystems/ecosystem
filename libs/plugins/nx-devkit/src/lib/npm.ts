import { Tree, updateJson } from '@nx/devkit';
import { sortObjectKeys } from './sort';

export const addNpmScript = (tree: Tree, scripts: Record<string, string>) => {
  updateJson(tree, 'package.json', (npmPackage) => {
    const sortedScripts = sortObjectKeys(
      { ...npmPackage.scripts, ...scripts },
      npmPackageScriptsSortGroups
    );
    return sortObjectKeys(
      { ...npmPackage, scripts: sortedScripts },
      npmPackageSortGroups
    );
  });
};

export const npmPackageSortGroups = [
  'name',
  'description',
  'version',
  'private',
  'license',
  'engines',
  'schematics',
  'scripts',
  'dependencies',
  'devDependencies',
];
export const npmPackageScriptsSortGroups = [
  'start',
  'build',
  'test',
  'prepare',
  'generate',
  'm1',
];
