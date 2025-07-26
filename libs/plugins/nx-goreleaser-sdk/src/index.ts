import { ProjectConfiguration } from 'nx/src/config/workspace-json-project-json';
import { dirname, join, relative, sep } from 'path';
import { existsSync } from 'fs';
import * as fs from 'fs';
import * as path from 'path';

export const createNodes = [
    '**/go.mod',
    (file: string): { projects: Record<string, ProjectConfiguration> } => {
        const projectRoot = dirname(file);
        const distPath = join('dist', projectRoot);
        const goreleaserPath = join(projectRoot, '.goreleaser.sdk.yaml');
        const packageJsonPath = join(projectRoot, 'package.json');

        if (!existsSync(packageJsonPath)) {
            return { projects: {} };
        }

        if (!existsSync(goreleaserPath)) {
            return { projects: {} };
        }

        try {
            const content = fs.readFileSync(packageJsonPath, 'utf-8');
            const json = JSON.parse(content);
            const projectName = json.name;
            const targets: ProjectConfiguration['targets'] = {
                build: {
                    executor: 'nx:run-commands',
                    outputs: [`{workspaceRoot}/${distPath}/artifacts.json`],
                    options: {
                        commands: [`goreleaser build --clean --snapshot --config ${projectRoot}/.goreleaser.sdk.yaml`],
                        parallel: false,
                        forwardAllArgs: false,
                        cwd: '',
                    },
                },
                test: {
                    executor: 'nx:run-commands',
                    options: {
                        command: 'go test -v ./... -cover -race',
                        cwd: projectRoot,
                    },
                },
                clean: {
                    executor: 'nx:run-commands',
                    options: {
                        command: `rm -rf ${distPath}`,
                        cwd: '',
                    },
                },
                lint: {
                    executor: 'nx:run-commands',
                    options: {
                        commands: ['golangci-lint run ./... --timeout=5m'],
                        parallel: false,
                        cwd: projectRoot,
                    },
                },
                format: {
                    executor: 'nx:run-commands',
                    options: {
                        commands: ['go mod tidy', 'gofumpt -l -w .', 'golangci-lint run ./... --timeout=5m --fix'],
                        parallel: false,
                        cwd: projectRoot,
                    },
                },
                'nx-release-publish': {
                    executor: 'nx:run-commands',
                    dependsOn: ['build'],
                    outputs: [`{workspaceRoot}/dist/${projectRoot}`],
                    options: {
                        commands: [`goreleaser release --config ${projectRoot}/.goreleaser.sdk.yaml --clean`],
                        parallel: false,
                        forwardAllArgs: false,
                    },
                    cache: false,
                },
                snapshot: {
                    executor: 'nx:run-commands',
                    dependsOn: ['build'],
                    outputs: [`{workspaceRoot}/dist/${projectRoot}`],
                    options: {
                        commands: [
                            `goreleaser release --config ${projectRoot}/.goreleaser.sdk.yaml --snapshot --clean --skip=sign,sbom --verbose`,
                        ],
                        parallel: false,
                        forwardAllArgs: false,
                    },
                    cache: false,
                },
                nightly: {
                    executor: 'nx:run-commands',
                    dependsOn: ['build'],
                    outputs: [`{workspaceRoot}/dist/${projectRoot}`],
                    options: {
                        commands: [
                            `goreleaser release --config ${projectRoot}/.goreleaser.sdk.yaml --nightly --clean --verbose`,
                        ],
                        parallel: false,
                        forwardAllArgs: false,
                    },
                    cache: false,
                },
            };

            const project: ProjectConfiguration = {
                name: projectName,
                root: projectRoot,
                projectType: 'library',
                targets,
                tags: ['lang:go', 'language:golang', 'type:sdk'],
            };

            return {
                projects: {
                    [projectName]: project,
                },
            };
        } catch (err) {
            console.warn(`Failed to read or parse package.json at ${packageJsonPath}:`, err);
            return { projects: {} };
        }
    },
];
