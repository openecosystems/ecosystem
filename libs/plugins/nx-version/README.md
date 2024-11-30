# plugins-nx-version

This library was generated with [Nx](https://nx.dev).

## Building

Run `nx build plugins-nx-version` to build the library.

## Running unit tests

Run `nx test plugins-nx-version` to execute the unit tests via [Jest](https://jestjs.io).

## Executors

### Bump

The bump executor increments the version number for the `package.json` file associated with the project. It is not able to be run on the main branch and will typically be run against pull requests.

The version on the main branch of the repository is the base version and the `commitMessage` parameter informs how the base version should be incremented. The prefix of the `commitMessage` should designate a release or non-release change and will apply the appropriate increment based on this. For `rc` commit messages the prerelease suffix will be removed from the version during publishing.
