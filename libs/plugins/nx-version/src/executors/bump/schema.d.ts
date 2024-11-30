export interface BumpExecutorSchema {
  branchName?: string;
  commitMessage: string;
  packageJson: string;
  baseSha?: string;
  headSha?: string;
  checkGitDiff?: boolean;
}
