import 'dotenv/config';
import { readJsonFile, workspaceRoot } from '@nx/devkit';
import * as core from '@nx-tools/core';
import { existsSync } from 'node:fs';
import { rm } from 'node:fs/promises';
import { join } from 'node:path';
import * as context from '@nx-tools/nx-container/src/executors/build/context';
import { PushExecutorSchema, VersionSchema } from './schema';

export default async function runExecutor(
  options: PushExecutorSchema
): Promise<{ success: true }> {
  const tmpDir = context.tmpDir();

  try {
    const version =
      typeof options?.version === 'string'
        ? options.version
        : await readVersion(options.version);

    for (const registry of options.registries) {
      const args = [
        'tag',
        options.image,
        `${registry}/${options.image}:${version}`,
      ];
      const tagCmd = { command: 'docker', args };
      const { stderr, exitCode } = await core.getExecOutput(
        tagCmd.command,
        tagCmd.args.map((arg) => core.interpolate(arg)),
        { ignoreReturnCode: true }
      );

      if (stderr.length > 0 && exitCode != 0) {
        const message = `tag failed with: ${
          stderr.match(/(.*)\s*$/)?.[0]?.trim() ?? 'unknown error'
        }`;
        throw new Error(message);
      }
    }

    for (const registry of options.registries) {
      const args = ['image', 'push', `${registry}/${options.image}:${version}`];
      const pushCmd = { command: 'docker', args };
      const { stderr, exitCode } = await core.getExecOutput(
        pushCmd.command,
        pushCmd.args.map((arg) => core.interpolate(arg)),
        { ignoreReturnCode: true }
      );

      if (stderr.length > 0 && exitCode != 0) {
        const message = `push failed with: ${
          stderr.match(/(.*)\s*$/)?.[0]?.trim() ?? 'unknown error'
        }`;
        throw new Error(message);
      }
    }
  } finally {
    await cleanup(tmpDir);
  }

  return { success: true };
}

async function readVersion({ path, key }: VersionSchema): Promise<string> {
  const data = await readJsonFile(join(workspaceRoot, path));
  if (!(key in data)) {
    throw new Error(`Version key: ${key} can not be read from ${path}`);
  }
  return data[key];
}

async function cleanup(tmpDir: string): Promise<void> {
  if (tmpDir.length > 0 && existsSync(tmpDir)) {
    await rm(tmpDir, { recursive: true });
  }
}
