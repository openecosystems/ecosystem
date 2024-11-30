import { generateHash, generateVersion } from './executor';
import {
  conventionalCommitTypes,
  ReleaseCommitType,
} from '@platform/nx-devkit';
import { SemVer } from 'semver';

const jiraBranchName = `DEVOPS-123/feature`;
const jiraBranchHash = `a9d1183`;

const headSha = '123456';

const releaseCandidatePrTitle = 'rc: DEVOPS-245 3.7 desc goes here';

const v1 = '1.0.0';
const releaseCommitTypes = conventionalCommitTypes
  .filter(({ release }) => release !== null)
  .filter(({ release }) => release !== false)
  .map(({ type, description, release }: ReleaseCommitType) => {
    const semVerV1 = new SemVer(v1);

    const commitMessage = `${type}: ${description}`;
    const newVersion = semVerV1.inc(release).raw.replace('-0', '');

    return { commitMessage, newVersion };
  });

const nonReleaseCommitTypes = conventionalCommitTypes
  .filter(({ release }) => release === false)
  .map(({ type, description }) => {
    const commitMessage = `${type}: ${description}`;
    return { commitMessage };
  });

describe('generate versions with valid jira branch', () => {
  it('should generate new prerelease versions based on the commit message', async () => {
    for (const { commitMessage, newVersion } of releaseCommitTypes) {
      const version = generateVersion(v1, commitMessage, jiraBranchName);
      expect(version).toBe(`${newVersion}-${jiraBranchHash}`);
    }
  });

  it('should retain main versions based on the commit message', async () => {
    for (const { commitMessage } of nonReleaseCommitTypes) {
      const version = generateVersion(v1, commitMessage, jiraBranchName);
      expect(version).toBe(v1);
    }
  });

  it('should retain main versions based on the commit message & add a hash if not present and add suffix when headSha is provided', async () => {
    for (const { commitMessage } of nonReleaseCommitTypes) {
      const existingHash = generateHash(jiraBranchName);
      const version = generateVersion(
        v1,
        commitMessage,
        jiraBranchName,
        headSha
      );
      expect(version).toBe(`${v1}-${existingHash}.${headSha}`);
    }
  });

  it('should retain main versions based on the commit message & add a suffix when headSha is provided', async () => {
    const mainVersion = `${v1}-a9d118`;
    for (const { commitMessage } of nonReleaseCommitTypes) {
      const existingHash = generateHash(jiraBranchName);
      const version = generateVersion(
        mainVersion,
        commitMessage,
        jiraBranchName,
        headSha
      );
      expect(version).toBe(`${v1}-${existingHash}.${headSha}`);
    }
  });

  it('should return a new version if the same hash exists on the main branch', async () => {
    const existingHash = generateHash('test');
    const version = generateVersion(
      `1.0.0-${existingHash}`,
      'break: removed duplicate',
      'test'
    );

    expect(version).toBe(`2.0.0-${existingHash}`);
  });

  it('should return a new version if the same hash exists on the main branch & add a suffix when headSha is provided', async () => {
    const existingHash = generateHash('test');
    const version = generateVersion(
      `1.0.0-${existingHash}`,
      'break: removed duplicate',
      'test',
      headSha
    );

    expect(version).toBe(`2.0.0-${existingHash}.${headSha}`);
  });

  it('should remove the hash when commit type is rc', async () => {
    expect(
      generateVersion('1.3.1-ad75ef', releaseCandidatePrTitle, jiraBranchName)
    ).toBe('1.3.1');
    expect(
      generateVersion('1.1.0-hash', releaseCandidatePrTitle, jiraBranchName)
    ).toBe('1.1.0');
  });
});

describe('generate versions on main branch', () => {
  it('should error if ran on main', () => {
    expect(() => generateVersion('1.0.0', 'feat: new feature', 'main')).toThrow(
      'Unable to run version change on main branch.'
    );
  });
});

describe('generate hash', () => {
  it('should generate a hash based on the branch name', () => {
    expect(generateHash('test-branch-name')).toBe('b515f41');
  });

  it('should replace the last digit with an x if the branch name does not contain a letter before position 6', () => {
    expect(generateHash('EQ6RWxRP02')).toBe('531079x');
  });
});
