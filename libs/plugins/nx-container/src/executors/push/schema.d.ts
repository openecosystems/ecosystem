export interface VersionSchema {
  path: string;
  key: string;
}

export interface PushExecutorSchema {
  image: string;
  version: string | VersionSchema;
  registries: Array<string>;
}
