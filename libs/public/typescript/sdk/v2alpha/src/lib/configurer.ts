// Represents the resolved configuration for the binding
export let resolvedConfiguration: Configuration | undefined;

// Configurable interface for bindings that can be configured
export interface Configurable {
  resolveConfiguration(): void;
  getDefaultConfiguration(): unknown;
  validateConfiguration(): void | Error;
}

// App class represents the application configuration
export class App {
  name?: string; // Corresponds to `yaml:"name,omitempty" env:"SERVICE_NAME"`
  version?: string; // Corresponds to `yaml:"version,omitempty" env:"VERSION_NUMBER"`
  environmentName?: string; // Corresponds to `yaml:"environmentName,omitempty" env:"ENV_NAME"`
  environmentType?: string; // Corresponds to `yaml:"environmentType,omitempty" env:"ENV_TYPE"`
  trace?: boolean; // Corresponds to `yaml:"trace,omitempty" env:"TRACE"`
  debug?: boolean; // Corresponds to `yaml:"debug,omitempty" env:"DEBUG"`

  constructor(
    name?: string,
    version?: string,
    environmentName?: string,
    environmentType?: string,
    trace?: boolean,
    debug?: boolean
  ) {
    this.name = name;
    this.version = version;
    this.environmentName = environmentName;
    this.environmentType = environmentType;
    this.trace = trace;
    this.debug = debug;
  }
}

// Configuration class holds the application's configuration
export class Configuration {
  app: App;

  constructor(app: App) {
    this.app = app;
  }
}

// Example usage:
// const appConfig = new App("MyApp", "1.0", "dev", "development", true, true);
// const configuration = new Configuration(appConfig);
