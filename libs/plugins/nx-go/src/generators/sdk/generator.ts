import {
  formatFiles,
  generateFiles,
  joinPathFragments,
  Tree,
} from '@nx/devkit';
import { SDKGeneratorSchema } from './schema';

export default async function (tree: Tree, options: SDKGeneratorSchema) {
  const { visibility, system, version } = options;
  let visibility_local = '';
  if (visibility == undefined || visibility == '' || visibility == null) {
    visibility_local = '';
  } else {
    if (visibility === 'poc' || visibility === 'internal')
      visibility_local = '-' + visibility;
  }

  const outputDirectory =
    'libs/public/go/sdk/' +
    system +
    '/' +
    version +
    '/' +
    system +
    '/';
  generateFiles(
    tree,
    joinPathFragments(__dirname, `./files`),
    outputDirectory,
    {
      visibility: visibility,
      system: system,
      version: version,
      visibility_local: visibility_local,
    }
  );
  await formatFiles(tree);
}
