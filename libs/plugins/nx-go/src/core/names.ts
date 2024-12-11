import { underscore } from '@nx/devkit/src/utils/string-utils';

export const groupId = 'health.platform.service';

export const getArtifactId = (name: string) => `oeco-${name}`;

export const getPackageName = (name: string) =>
  `${groupId}.${underscore(name)}`;
