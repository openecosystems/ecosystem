import { ExecutorContext, logger, workspaceRoot } from '@nx/devkit';
import * as crypto from 'crypto';
import { join } from 'path';
import { ReleaseType, SemVer } from 'semver';
import { BumpExecutorSchema } from './schema';
import { getExecOutput, interpolate } from '@nx-tools/core';
import {
  conventionalCommitTypes,
  NonPreReleaseType,
  NonReleaseCommitType,
  ReleaseCommitType,
  updateJsonFileSystem,
  updateXmlFileSystem,
} from '@platform/nx-devkit';
import { readdir } from 'fs/promises';

export default async function runExecutor(
  options: BumpExecutorSchema,
  context: ExecutorContext
) {
  const {
    name,
    root: projectRoot,
    projectType,
    tags,
  } = context.projectsConfigurations.projects[context.projectName];
  const {
    commitMessage,
    branchName,
    packageJson,
    baseSha,
    headSha,
    checkGitDiff,
  } = options;
  const packagePath = join(context.root, packageJson);

  const isDotnet = tags?.includes('framework:dotnet') ?? false;
  const isLibrary = projectType === 'library';
  const isTest = name.endsWith('test');

  if (checkGitDiff && !(isLibrary && isDotnet && !isTest)) {
    const files = await getGitDiff(projectRoot, baseSha);

    if (files.length === 0) {
      logger.warn(`No files changed, skipping bump on ${context.projectName}.`);
      return { success: true };
    }
  }

  const mainVersion = await getVersionFromMain(packageJson, baseSha);

  const newVersion = generateVersion(
    mainVersion,
    commitMessage,
    branchName,
    headSha
  );

  logger.info(
    `[${context.projectName}] FROM ${mainVersion} (main) TO ${newVersion} (${branchName})`
  );

  const results = await Promise.allSettled([
    updateJsonFileSystem(packagePath, async (p) => {
      return { ...p, version: newVersion };
    }),
    findCsprojFile(projectRoot).then((csProjFile) => {
      if (!csProjFile) {
        return;
      }
      return updateXmlFileSystem(join(projectRoot, csProjFile), async ($) => {
        const versionElement = $('Version');
        if (versionElement.length === 0) {
          const firstPropertyGroup = $('PropertyGroup').first();
          firstPropertyGroup.append(`  <Version>${newVersion}</Version>\n  `);
        } else {
          versionElement.text(newVersion);
        }
        return $;
      });
    }),
  ]);

  const hasError = results.some((r) => r.status === 'rejected');
  if (hasError) {
    throw new Error('Unable to update version.');
  }

  return { success: true };
}

export async function getGitDiff(
  basePath: string,
  baseSha = 'main'
): Promise<Array<string>> {
  const { stdout } = await getExecOutput(
    'git',
    ['diff', '--name-only', baseSha, '--', basePath].map((arg) =>
      interpolate(arg)
    ),
    { cwd: workspaceRoot }
  );
  return stdout.split('\n').filter((f) => f.length > 0);
}

export async function getVersionFromMain(
  packagePath: string,
  baseSha = 'main'
): Promise<string> {
  const { stdout, exitCode } = await getExecOutput(
    'git',
    ['show', `${baseSha}:${packagePath}`].map((arg) => interpolate(arg)),
    { ignoreReturnCode: true, cwd: workspaceRoot }
  );

  const { version = '0.0.0' } = JSON.parse(
    exitCode === 0 ? stdout?.trim() : null ?? '{}'
  );
  return /-\d{6}$/.test(version) ? `${version}x` : version;
}

export function generateVersion(
  mainVersion: string,
  commitMessage: string,
  branchName: string,
  headSha?: string
): string {
  const newVersion = new SemVer(mainVersion);
  const isMain = branchName === 'main';
  if (isMain) {
    throw new Error('Unable to run version change on main branch.');
  }

  const analyzedReleaseType = analyzeCommits(
    conventionalCommitTypes,
    commitMessage
  );

  const isReleaseCandidate = commitMessage.startsWith('rc:');
  if (isReleaseCandidate) {
    return newVersion.raw.replace(/-.+$/, '');
  }

  const hash = generateHash(branchName);

  if (!analyzedReleaseType) {
    return headSha
      ? `${mainVersion.replace(/-.+$/, '')}-${hash}.${headSha.substring(0, 7)}`
      : mainVersion;
  }

  const releaseType: ReleaseType = `pre${analyzedReleaseType}`;

  newVersion.inc(releaseType, hash);

  const formattedVersion = newVersion.format();
  const version = formattedVersion.slice(0, formattedVersion.lastIndexOf('.'));
  return headSha ? `${version}.${headSha.substring(0, 7)}` : version;
}

export function generateHash(branchName: string): string {
  const sha1 = crypto.createHash('sha1').update(branchName).digest('hex');
  const match = sha1.search(/[a-zA-Z]/);
  if (match === -1 || match > 6) {
    return `${sha1.substring(0, 6)}x`;
  }
  return sha1.substring(0, 7);
}

export function analyzeCommits(
  releaseRules: Array<ReleaseCommitType | NonReleaseCommitType>,
  commit: string
): NonPreReleaseType | null {
  const match = releaseRules.find(({ type }) => commit.startsWith(type));
  return match?.release || null;
}

export async function findCsprojFile(directoryPath: string): Promise<string> {
  const files = await readdir(directoryPath);
  return files.find((f) => f.endsWith('.csproj'));
}
