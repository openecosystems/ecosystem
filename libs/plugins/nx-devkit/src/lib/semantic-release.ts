export type ReleaseTypeKeys = typeof releaseTypeKeys[number];
export type NonReleaseTypeKeys = typeof nonReleaseTypeKeys[number];
export type NonPreReleaseType = 'major' | 'minor' | 'patch' | null;

export type ReleaseCommitType = {
  type: ReleaseTypeKeys;
  description: string;
  title: string;
  release: NonPreReleaseType;
};

export type NonReleaseCommitType = {
  type: NonReleaseTypeKeys;
  description: string;
  title: string;
  release: false;
};

export const conventionalCommitTypes: Array<
  ReleaseCommitType | NonReleaseCommitType
> = [
  {
    type: 'feat',
    description: 'A new feature',
    title: 'Features',
    release: 'minor',
  },
  {
    type: 'fix',
    description: 'A bug fix',
    title: 'Bug Fixes',
    release: 'patch',
  },
  {
    type: 'docs',
    description: 'Documentation only changes',
    title: 'Documentation',
    release: false,
  },
  {
    type: 'style',
    description:
      'Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)',
    title: 'Styles',
    release: false,
  },
  {
    type: 'refactor',
    description: 'A code change that neither fixes a bug nor adds a feature',
    title: 'Code Refactoring',
    release: false,
  },
  {
    type: 'perf',
    description: 'A code change that improves performance',
    title: 'Performance Improvements',
    release: 'patch',
  },
  {
    type: 'test',
    description: 'Adding missing tests or correcting existing tests',
    title: 'Tests',
    release: false,
  },
  {
    type: 'build',
    description:
      'Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)',
    title: 'Builds',
    release: false,
  },
  {
    type: 'ci',
    description:
      'Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)',
    title: 'Continuous Integrations',
    release: false,
  },
  {
    type: 'chore',
    description: "Other changes that don't modify src or test files",
    title: 'Chores',
    release: false,
  },
  {
    type: 'revert',
    description: 'Reverts a previous commit',
    title: 'Reverts',
    release: false,
  },
  {
    type: 'break',
    description:
      'Breaking change that will require a major version number increment',
    title: 'Braking',
    release: 'major',
  },
  {
    type: 'rc',
    description: 'Convert prerelease versions to stable versions.',
    title: 'Release Candidate',
    release: null,
  },
];

const releaseTypeKeys = ['feat', 'fix', 'perf', 'break', 'rc'] as const;
const nonReleaseTypeKeys = [
  'docs',
  'style',
  'refactor',
  'test',
  'build',
  'ci',
  'chore',
  'revert',
] as const;
