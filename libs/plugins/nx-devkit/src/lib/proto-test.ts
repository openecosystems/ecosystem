import { ApiVisibility, apiVisibilityToSafeProtoDir, Version } from './proto';
import { kebabCase, snakeCase } from 'lodash';
import { join, relative } from 'path';
import {
  addProjectConfiguration,
  ProjectConfiguration,
  workspaceRoot,
} from '@nx/devkit';
import { readFile } from 'fs/promises';
import { directories, Directory } from './nx';

export const getProtoRelativePath = (system, version: Version, servicePath) => {
  const [service, ...serviceParentDirs] = servicePath.split('/').reverse();
  return `${system}/${version}/${serviceParentDirs?.join('/')}${
    serviceParentDirs?.length ? '/' : ''
  }${snakeCase(service)}.proto`;
};

export const getProtoPath = (
  visibility: ApiVisibility = 'public',
  org,
  system,
  version: Version,
  servicePath
) => {
  const relativePath = getProtoRelativePath(system, version, servicePath);
  return `proto/${apiVisibilityToSafeProtoDir[visibility]}/${org}/${relativePath}`;
};

export const getProto = (org, system, version) => {
  return `syntax = "proto3";
package ${org}.${system}.${version};

import "platform/options/v2/annotations.proto";

service TestService {
  option (platform.options.v2.service) = {grpc_port: 7100, http_port: 52100};
}`;
};

export const getTestProto = (
  system,
  version: Version,
  servicePath,
  org = 'platform',
  visibility: ApiVisibility = 'public'
) => {
  const protoData = getProto(org, system, version);
  const protoPath = getProtoPath(visibility, org, system, version, servicePath);
  const protoRelativePath = getProtoRelativePath(system, version, servicePath);
  return { protoData, protoPath, protoRelativePath };
};

export const addSchema = async (tree, dir: string) => {
  const schemaAbsolutePath = join(dir, 'schema.json');
  const schemaBuffer = await readFile(schemaAbsolutePath);
  const schemaPath = relative(workspaceRoot, schemaAbsolutePath);
  tree.write(schemaPath, schemaBuffer);
};

export interface ProjectOptions {
  system?: string;
  version?: Version;
  servicePath?: string;
  serviceNumber?: number;
  visibility?: ApiVisibility;
  directory?: Directory;
}

const defaults: ProjectOptions = {
  system: 'test',
  version: 'v2alpha',
  servicePath: 'admin/sponge-bob',
  visibility: 'public',
  directory: 'microservice',
  serviceNumber: null,
};

export const getTestProject = (options: ProjectOptions = {}) => {
  const {
    system,
    version,
    servicePath,
    serviceNumber,
    directory: subdirectoryType,
    visibility,
  } = { ...defaults, ...options };
  const appLibsSubdirectory = directories[subdirectoryType];
  const projectSuffix = serviceNumber ? `-${serviceNumber}` : '';
  const project = `${appLibsSubdirectory}-${visibility}-${system}-${version}-${kebabCase(
    servicePath
  )}${projectSuffix}`;
  const directory = `${appLibsSubdirectory}/${visibility}/${system}/${version}`;
  const root = `apps/${directory}/${servicePath}`;
  const [service, ...serviceParentDirs] = servicePath.split('/').reverse();
  return {
    system,
    version,
    servicePath,
    project,
    root,
    directory,
    visibility,
    service,
    serviceParentDirs,
  };
};

export type ProtoOptions = ProjectOptions & { org?: string };

export const addProto = (tree, options: ProtoOptions = {}) => {
  const { system, version, servicePath, org, visibility } = {
    ...defaults,
    ...options,
  };
  const { protoData, protoPath } = getTestProto(
    system,
    version,
    servicePath,
    org,
    visibility
  );
  tree.write(protoPath, protoData);
};

export const addDockerCompose = (tree, options: ProjectOptions = null) => {
  const dockerComposeContents = `
version: '3.9'
services:
    `;
  tree.write('docker-compose.yaml', dockerComposeContents);
};

export const addDockerComposeProxy = (tree, options: ProjectOptions = null) => {
  const dockerComposeContents = `
version: '3.9'
services:
    `;
  tree.write('infrastructure/local/docker/proxies.yaml', dockerComposeContents);
};

export const addTestProject = (tree, options: ProjectOptions = null) => {
  const { project, root } = getTestProject(options);
  addProjectConfiguration(tree, project, {
    root,
    projectType: 'application',
    targets: { build: { executor: '@nx/webpack:webpack' } },
  } as ProjectConfiguration);
  addProto(tree, options);
};

export const addTestLibProject = (tree, project: string) => {
  addProjectConfiguration(tree, kebabCase(project), {
    root: `libs/${project}`,
    sourceRoot: `libs/${project}/src`,
    projectType: 'library',
    targets: {},
  } as ProjectConfiguration);
};
