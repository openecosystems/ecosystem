import { createTreeWithEmptyWorkspace } from '@nx/devkit/testing';
import {
  addProjectConfiguration,
  readProjectConfiguration,
  Tree,
} from '@nx/devkit';
import generator from './generator';
import { SDKGeneratorSchema } from './schema';

jest.mock('@nxrocks/nx-spring-boot', () => ({
  projectGenerator: jest.fn(),
}));

describe('library generator', () => {
  let appTree: Tree;

  const options: SDKGeneratorSchema = { system: 'test', version: 'v1alpha' };

  beforeEach(() => {
    appTree = createTreeWithEmptyWorkspace({ layout: 'apps-libs' });

  });


});
