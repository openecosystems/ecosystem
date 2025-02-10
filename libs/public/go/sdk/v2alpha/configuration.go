//nolint:revive
package sdkv2alphalib

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	"dario.cat/mergo"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

// viperInstance is a singleton instance of the viper.Viper configuration library.
// viperOnce ensures viperInstance is initialized only once in a thread-safe manner.
// Config holds global settings for the application, managed via SpecSettings.
// Overrides contains runtime-specific configuration that can override the default settings.
var (
	viperInstance *viper.Viper
	viperOnce     sync.Once
	// Config Global Flags that can override configuration at runtime
	Config    *specv2pb.SpecSettings
	Overrides *RuntimeConfigurationOverrides
)

// SpecConfigurationProvider defines an interface for obtaining and monitoring specification settings.
// GetConfigurations retrieves the current specification settings.
// WatchConfigurations initiates a watching mechanism to monitor settings updates.
type SpecConfigurationProvider interface {
	CreateConfiguration() (interface{}, error)
	GetConfiguration() interface{}
	WatchConfigurations() error
	ResolveConfiguration()
	GetDefaultConfiguration() interface{}
	ValidateConfiguration() error
}

// ConfigurationContext represents a configuration and context manager for managing platform-specific and file system settings.
type ConfigurationContext struct {
	filesystem *FileSystem
	configurer *viper.Viper
	cfg        *configurationProviderOption
}

// ConfigurationProvider manages and provides application server configuration settings.
type ConfigurationProvider struct {
	filesystem *FileSystem
	configurer *viper.Viper
	cfg        *configurationProviderOption
}

// NewConfigurationProvider initializes a new ConfigurationProvider with default SpecSettings configuration.
func NewConfigurationProvider(opts ...ConfigurationProviderOption) (*ConfigurationProvider, error) {
	sctx, err := getSettingsContext(opts...)
	if err != nil {
		return nil, err
	}

	sctx.configurer.SetConfigName(getConfigFileName(sctx.cfg.ConfigPathPrefix, sctx.cfg.PlatformContext))
	sctx.configurer.SetConfigType(ConfigurationExtension)
	sctx.configurer.AddConfigPath(ConfigurationDirectory)
	sctx.configurer.AddConfigPath(".")
	if sctx.cfg.ConfigPath != "" {
		sctx.configurer.AddConfigPath(sctx.cfg.ConfigPath)
	}

	// var c specv2pb.SpecSettings
	// ResolveConfiguration(sctx.configurer, &c, sctx.cfg.DefaultSettings)

	// Config = &c

	return &ConfigurationProvider{
		filesystem: sctx.filesystem,
		configurer: sctx.configurer,
		cfg:        sctx.cfg,
	}, nil
}

func (s ConfigurationProvider) CreateConfiguration() (interface{}, error) {
	_r := *s.cfg.ConfigurationResolver
	return _r.CreateConfiguration()
}

func (s ConfigurationProvider) GetConfiguration() interface{} {
	_r := *s.cfg.ConfigurationResolver

	structType := reflect.TypeOf(_r)

	fmt.Println("Struct Type:", structType.Name())
	fmt.Println("Full Type:", structType.String())

	return _r.GetConfiguration()
}

func (s ConfigurationProvider) WatchConfigurations() error {
	_r := *s.cfg.ConfigurationResolver
	return _r.WatchConfigurations()
}

func (s ConfigurationProvider) ResolveConfiguration() {
	_r := *s.cfg.ConfigurationResolver
	_r.ResolveConfiguration()
}

func (s ConfigurationProvider) GetDefaultConfiguration() interface{} {
	_r := *s.cfg.ConfigurationResolver
	return _r.GetDefaultConfiguration()
}

func (s ConfigurationProvider) ValidateConfiguration() error {
	_r := *s.cfg.ConfigurationResolver
	return _r.ValidateConfiguration()
}

// DotConfigSettingsProvider is responsible for managing and providing configuration settings parsed
// from a .config file. It encapsulates the SpecSettings structure to hold configuration data.
type DotConfigSettingsProvider struct {
	settings *specv2pb.SpecSettings
}

// NewDotConfigSettingsProvider initializes a new instance of DotConfigSettingsProvider by loading and unmarshalling configuration.
func NewDotConfigSettingsProvider() (*DotConfigSettingsProvider, error) {
	_, configurer, err := getFileSystemAndConfigurer("")
	if err != nil {
		return nil, err
	}

	c := specv2pb.SpecSettings{}

	err = configurer.Unmarshal(&c)
	if err != nil {
		fmt.Println("Could not process configuration file")
		return nil, err
	}

	Config = &c

	return &DotConfigSettingsProvider{
		settings: &c,
	}, nil
}

// CreateConfiguration initializes and returns a new SpecSettings instance along with any potential errors encountered.
func (p *DotConfigSettingsProvider) CreateConfiguration() (*specv2pb.SpecSettings, error) {
	return createContextSettings("test", "test")
}

// GetConfiguration retrieves the current SpecSettings instance from the DotConfigSettingsProvider.
func (p *DotConfigSettingsProvider) GetConfiguration() *specv2pb.SpecSettings {
	return p.settings
}

// WatchConfigurations sets up file system watchers to monitor changes in settings files and triggers appropriate updates.
func (p *DotConfigSettingsProvider) WatchConfigurations() error {
	return watchSettings(p.settings, Filesystem.ContextDirectory)
}

// SpecYamlSettingsProvider provides access to spec settings configured in a YAML file.
// It encapsulates and manages a SpecSettings instance for configuration retrieval and watching updates.
type SpecYamlSettingsProvider struct {
	settings *specv2pb.SpecSettings
}

// NewSpecYamlSettingsProvider creates a new instance of SpecYamlSettingsProvider by loading and resolving YAML configuration.
func NewSpecYamlSettingsProvider() (*SpecYamlSettingsProvider, error) {
	viperOnce.Do(
		func() {
			_ = godotenv.Load()

			viperInstance = viper.New()
			viperInstance.SetConfigName("spec")
			viperInstance.SetConfigType("yaml")
			viperInstance.AddConfigPath("/etc/spec")
			viperInstance.AddConfigPath("$HOME/.spec")
			viperInstance.AddConfigPath(".")
			// viperInstance.AutomaticEnv()
		},
	)

	var c specv2pb.SpecSettings
	Resolve(&c, specv2pb.SpecSettings{})

	Config = &c

	return &SpecYamlSettingsProvider{
		settings: &c,
	}, nil
}

// CreateSettings initializes and returns a new SpecSettings instance along with any potential errors encountered.
func (p *SpecYamlSettingsProvider) CreateConfiguration() (*specv2pb.SpecSettings, error) {
	return createContextSettings("test", "test")
}

// GetSettings retrieves the SpecSettings instance associated with the SpecYamlSettingsProvider.
func (p *SpecYamlSettingsProvider) GetConfiguration() *specv2pb.SpecSettings {
	return p.settings
}

// WatchSettings monitors file changes in specified directories to dynamically reload SpecSettings if enabled.
func (p *SpecYamlSettingsProvider) WatchConfigurations() error {
	return watchSettings(p.settings, ".", "/etc/spec")
}

// RuntimeConfigurationOverrides holds runtime configuration flags and settings that override default behavior.
type RuntimeConfigurationOverrides struct {
	Context      *string
	Logging      *bool
	Verbose      *bool
	VerboseLog   *bool
	LogFile      *string
	Quiet        *bool
	FieldMask    string
	ValidateOnly bool
}

// CLISettingsProvider manages and provides specification settings for a CLI runtime environment.
// It includes parsed settings and runtime overrides for configuration flexibility.
type CLISettingsProvider struct {
	settings *specv2pb.SpecSettings
	Flags    *RuntimeConfigurationOverrides
}

// NewCLISettingsProvider creates a new CLISettingsProvider by loading configuration and applying runtime overrides.
func NewCLISettingsProvider(flags *RuntimeConfigurationOverrides) (*CLISettingsProvider, error) {
	platformContext := ""
	if flags != nil && flags.Context != nil {
		platformContext = *flags.Context
	}

	_, configurer, err := getFileSystemAndConfigurer(platformContext)
	if err != nil {
		return nil, err
	}

	c := specv2pb.SpecSettings{}

	err = configurer.Unmarshal(&c)
	if err != nil {
		fmt.Println("Could not process configuration file")
		return nil, err
	}

	Config = &c
	Overrides = flags

	return &CLISettingsProvider{
		settings: &c,
	}, nil
}

// CreateSettings initializes and returns a new SpecSettings instance along with any potential errors encountered.
func (p *CLISettingsProvider) CreateConfiguration() (*specv2pb.SpecSettings, error) {
	return createContextSettings("test", "test")
}

// GetSettings retrieves the current SpecSettings instance managed by the CLISettingsProvider.
func (p *CLISettingsProvider) GetConfiguration() *specv2pb.SpecSettings {
	return p.settings
}

// WatchSettings monitors filesystem changes to dynamically reload settings if enabled in the configuration.
func (p *CLISettingsProvider) WatchConfigurations() error {
	return watchSettings(p.settings, Filesystem.ContextDirectory)
}

// ResolveConfiguration reads configuration, unmarshals it into the destination, validates required fields, and merges with the source structure.
func ResolveConfiguration(configurer *viper.Viper, dst, src interface{}) {
	err := configurer.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("No spec configuration found.")
		}
	}

	err = configurer.Unmarshal(dst, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.TextUnmarshallerHookFunc(),
			StringExpandEnv(),
		),
	))
	if err != nil {
		panic(err)
	}

	validate := validator.New()
	if err = validate.Struct(dst); err != nil {
		fmt.Println("Missing required attributes", err)
	}

	if err = mergo.Merge(dst, src); err != nil {
		fmt.Println("Error merging settings configuration:", err)
	}

	ImportEnvironmentVariables(dst, "", "")
}

// getSettingsContext creates and initializes a FileSystem and viper.Viper configurer, handling context overrides and config file setup.
//
//nolint:unparam
func getSettingsContext(opts ...ConfigurationProviderOption) (*ConfigurationContext, error) {
	fs := NewFileSystem()
	configurer := viper.New()

	var cfg configurationProviderOption
	for _, opt := range opts {
		opt.apply(&cfg)
	}

	var ctx string
	if cfg.PlatformContext != "" {
		fmt.Println("Overriding context to: " + cfg.PlatformContext)

		_, err := fs.Exists(filepath.Join(ContextDirectory, cfg.PlatformContext))
		if err != nil {
			fmt.Println("Error: context does not exists: " + err.Error())
			return nil, errors.New("context does not exists")
		}

		ctx = cfg.PlatformContext
	} else {
		file, err := fs.ReadFile(DefaultContextFile)
		if err != nil {
			return nil, errors.New("could not read config file: " + err.Error())
		}

		ctx = strings.TrimSpace(string(file))

		if ctx == "" {
			// Set the oeco workspace in the "default" file
			err = fs.WriteFile(DefaultContextFile, []byte(OecoContextFileName), os.ModePerm)
			if err != nil {
				return nil, errors.New("internal error: Cannot create default context")
			}

			_, err = createContextSettings(OecoContextFileName, DefaultCIDR)
			if err != nil {
				return nil, err
			}
		}
	}

	WithPlatformContext(ctx).apply(&cfg)

	return &ConfigurationContext{
		filesystem: fs,
		configurer: configurer,
		cfg:        &cfg,
	}, nil
}

// getFileSystemAndConfigurer creates and initializes a FileSystem and viper.Viper configurer, handling context overrides and config file setup.
//
//nolint:unparam
func getFileSystemAndConfigurer(platformContext string) (*FileSystem, *viper.Viper, error) {
	fs := NewFileSystem()
	ufs := fs.UnderlyingFileSystem
	contextDir := fs.ContextDirectory

	configurer := viper.New()

	// Set Flag Overrides
	if platformContext != "" {
		fmt.Println("Overriding context to: " + platformContext)

		exists, err := fs.Exists(filepath.Join(ContextDirectory, platformContext))
		if err != nil {
			fmt.Println("Error: context does not exists: " + err.Error())
			return nil, nil, errors.New("context does not exists")
		}

		if exists {
			// Use config file from the flag
			configurer.SetConfigFile(platformContext)
		}
	} else {
		file, err := fs.ReadFile(DefaultContextFile)
		if err != nil {
			return nil, nil, errors.New("could not read config file: " + err.Error())
		}

		ctx := strings.TrimSpace(string(file))

		if ctx == "" {
			// Set the oeco workspace in the "default" file
			err := fs.WriteFile(DefaultContextFile, []byte(OecoContextFileName), os.ModePerm)
			if err != nil {
				return nil, nil, errors.New("internal error: Cannot create default context")
			}

			fixthis := `
---
name: oeco
description: Open Economic System Context
platform:
    endpoint: http://localhost:6577
    insecure: true
    mesh:
      enabled: true
      endpoint: http://192.168.100.5:6477
      insecure: true
    dynamicconfigreload: false
    punchy:
      punch: true
      respond: true
      delay: 1s
      respond_delay: 5s
    statichostmap:
      - '192.168.100.1':
          map:
            "45.63.49.173:4242"
    lighthouse:
      interval: 60
      hosts:
        - '192.168.100.1'
context:
    headers:
      - key: "x-spec-ecosystem-slug"
        values:
          - "oeco"
systems2:
  - name: configuration
    version: v2alpha
  - name: iam
    version: v2alpha
`

			err = fs.WriteFile(OecoContextFile+"."+ConfigurationExtension, []byte(fixthis), os.ModePerm)
			if err != nil {
				return nil, nil, errors.New("internal error: Cannot create default context")
			}

			ctx = OecoContextFileName
		}

		configurer.SetFs(ufs)
		configurer.SetConfigName(ctx)
		configurer.SetConfigType(ConfigurationExtension)
		configurer.AddConfigPath(contextDir)
	}

	// Set Environment Variable Overrides
	configurer.AutomaticEnv()

	if err := configurer.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found: " + err.Error())
		}
	}

	return fs, configurer, nil
}

// c
func createContextSettings(ecosystemName string, cidr string) (*specv2pb.SpecSettings, error) {
	// TODO: Sanitize ecosystemName

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Println("IP:", ip.String())
	fmt.Println("Subnet Mask:", ipnet.Mask)

	fs := NewFileSystem()
	contextFile := filepath.Join(ContextDirectory, ecosystemName)

	exists, err := fs.Exists(contextFile)
	if err != nil {
		return nil, errors.New("could not check if context file exists: " + err.Error())
	}

	if exists {
		return nil, errors.New("context file already exists: " + contextFile)
	}

	slugHeader := typev2pb.Header{
		Key:    "x-spec-ecosystem-slug",
		Values: []string{ecosystemName},
	}

	configurationSystem := specv2pb.SpecSystem{
		Name:    "configuration",
		Version: "v2alpha",
	}

	iamSystem := specv2pb.SpecSystem{
		Name:    "iam",
		Version: "v2alpha",
	}

	settings := specv2pb.SpecSettings{
		Name: ecosystemName,
		App: &specv2pb.App{
			Name:            ecosystemName,
			Version:         "v2.0.0",
			Description:     strings.ToUpper(ecosystemName[:1]) + ecosystemName[1:] + " Ecosystem Context",
			EnvironmentName: "local-1",
			EnvironmentType: "local",
			Debug:           false,
			Verbose:         false,
		},
		Platform: &specv2pb.Platform{
			Endpoint:            "localhost:6577",
			Insecure:            true,
			DnsEndpoints:        []string{"localhost:4242"},
			DynamicConfigReload: false,
			Mesh: &specv2pb.Mesh{
				Enabled:     true,
				Endpoint:    "",
				Insecure:    true,
				DnsEndpoint: "",
			},
		},
		Context: &specv2pb.Context{
			Headers: []*typev2pb.Header{&slugHeader},
		},
		Systems: []*specv2pb.SpecSystem{&configurationSystem, &iamSystem},
	}

	settingBytes, err := proto.Marshal(&settings)
	if err != nil {
		return nil, err
	}

	err = fs.WriteFile(OecoContextFile+"."+ConfigurationExtension, settingBytes, os.ModePerm)
	if err != nil {
		return nil, errors.New("internal error: Cannot write ecosystem settings file")
	}

	return nil, nil
}

// watchSettings enables watching for configuration file changes in specified directories if dynamic reload is allowed.
// settings: The specification settings containing configuration details and dynamic reload flag.
// directories: One or more directories to monitor for configuration file changes.
// Returns an error if the file watcher setup fails or if issues occur during monitoring events.
func watchSettings(settings *specv2pb.SpecSettings, directories ...string) error {
	// If dynamic settings enabled, turn on filesystem notification
	if settings.Platform != nil && !settings.Platform.DynamicConfigReload {
		// fmt.Println("Dynamic reload is disabled in your settings. Please enable if you want to Watch Configurations: settings.platform.DynamicConfigReload")
		return nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer func(watcher *fsnotify.Watcher) {
		_ = watcher.Close()
	}(watcher)

	// Add the directory to be watched

	for _, dir := range directories {
		err = watcher.Add(dir)
		if err != nil {
			return err
		}
	}

	fmt.Println("Watching for context file changes in: ", strings.Join(directories, " "))

	// Watch for events
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 && filepath.Ext(event.Name) == ".yaml" {
				fmt.Println("Detected change in:", event.Name)

				// Update settings
				for _, e := range settings.Systems {
					fmt.Println(e.Name + ": " + e.Version)
				}
				//err := reloadProtoFile(event.Name)
				//if err != nil {
				//	fmt.Println("Error reloading proto file:", err)
				//}
				//
				//protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
				//	fmt.Println("File:", fd.Path())
				//	return true // continue iteration
				//})
			}
		case err := <-watcher.Errors:
			fmt.Println("Watcher error:", err)
		}
	}
}

// PackageJson represents the structure of a package.json file containing name and version information.
type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// setEnv maps and sets an environment variable value to a corresponding YAML configuration field using Viper.
// It converts the YAML field into an environment variable format, optionally prepending a prefix if provided.
// If the environment variable is present, its value is used to update the corresponding Viper configuration field.
func setEnv(envPrefix, yaml string) {
	envVar := strcase.ToScreamingSnake(strcase.ToLowerCamel(strings.ReplaceAll(yaml, ".", "_")))
	if envPrefix != "" {
		envVar = envPrefix + "_" + envVar
	}
	// Uncomment to print all variables
	// fmt.Println(envVar + "=")
	val, present := os.LookupEnv(envVar)
	if present {
		fmt.Printf("Setting %s from %s to %s\n", yaml, envVar, val)
		viperInstance.Set(yaml, val)
	}
}

// ImportEnvironmentVariables sets configuration values from environment variables based on struct fields and provided tags.
// It processes "mapstructure" and "env" tags, supports nested structs, and recursively applies prefix logic for key formatting.
func ImportEnvironmentVariables(iface interface{}, envPrefix string, prefix string) {
	// https://github.com/spf13/viper/issues/188#issuecomment-399884438
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)

		fieldName := t.Name
		tagValue, ok := t.Tag.Lookup("mapstructure")
		if ok {
			tagValue = strings.SplitN(tagValue, ",", 2)[0]
			if strings.Contains(tagValue, ",") {
				fieldName = tagValue
			}
		}
		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		// Check for env tag to override config field
		envVar, ok := t.Tag.Lookup("env")
		if ok {
			val, present := os.LookupEnv(envVar)
			if present {
				fmt.Printf("Setting %s from %s to %s\n", fieldName, envVar, val)
				viperInstance.Set(fieldName, val)
				continue
			}
		}

		switch v.Kind() {
		case reflect.Struct:
			instance := reflect.New(t.Type)
			if _, ok := instance.Interface().(encoding.TextUnmarshaler); ok {
				setEnv(envPrefix, fieldName)
			} else {
				if !IsLower(fieldName) {
					ImportEnvironmentVariables(v.Interface(), envPrefix, fieldName)
				}
			}
		default:
			setEnv(envPrefix, fieldName)
		}
	}
}

// ImportPackageJson reads and parses the `package.json` file, updating the Configuration with the app name and version.
// Panics if the file cannot be read or the `name` and `version` fields are missing.
func ImportPackageJson() (string, string, error) {
	content, err := os.ReadFile("./package.json")
	if err != nil {
		return "", "", err
	}

	var data PackageJson
	err = json.Unmarshal(content, &data)
	if err != nil {
		panic("Error during Unmarshal of package.json: " + err.Error())
	}

	if data.Name == "" {
		panic("name not found in package.json")
	}

	if data.Version == "" {
		panic("version not found in package.json")
	}

	return data.Name, data.Version, nil
}

// StringExpandEnv creates a DecodeHookFuncKind that replaces environment variable placeholders in strings with their values.
func StringExpandEnv() mapstructure.DecodeHookFuncKind {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{},
	) (interface{}, error) {
		if f != reflect.String || t != reflect.String {
			return data, nil
		}

		return os.ExpandEnv(data.(string)), nil
	}
}

// Resolve reads configuration, unmarshals it into the destination, validates required fields, and merges with the source structure.
func Resolve(dst, src interface{}) {
	err := viperInstance.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("No spec configuration found.")
		}
	}

	// ImportEnvironmentVariables(config, "", "")

	err = viperInstance.Unmarshal(dst, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.TextUnmarshallerHookFunc(),
			StringExpandEnv(),
		),
	))
	if err != nil {
		panic(err)
	}

	validate := validator.New()
	if err = validate.Struct(dst); err != nil {
		fmt.Println("Missing required attributes", err)
	}

	if err = mergo.Merge(dst, src); err != nil {
		fmt.Println("Error merging settings configuration:", err)
	}
}

func getConfigFileName(prefix, e string) string {
	if prefix == "" {
		return e + ".yaml"
	}
	return prefix + "." + e + ".mesh.yaml"
}

// IsLower checks if all characters in the provided string are lowercase alphabetic letters (a-z). Returns true if so, otherwise false.
func IsLower(s string) bool {
	for _, charNumber := range s {
		if charNumber > 122 || charNumber < 97 {
			return false
		}
	}
	return true
}

// configurationProviderOption is the configuration for a Server.
type configurationProviderOption struct {
	PlatformContext       string
	ConfigPath            string
	ConfigPathPrefix      string
	DefaultSettings       *specv2pb.SpecSettings
	WatchSettings         bool
	ConfigurationResolver *SpecConfigurationProvider
}

// ConfigurationProviderOption is an interface for applying configurations to a config object.
type ConfigurationProviderOption interface {
	apply(*configurationProviderOption)
}

type optionFunc func(*configurationProviderOption)

func (f optionFunc) apply(cfg *configurationProviderOption) { f(cfg) }

// WithOptions composes multiple Options into one.
func WithOptions(opts ...ConfigurationProviderOption) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

func WithPlatformContext(context string) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		cfg.PlatformContext = context
	})
}

// WithConfigPathPrefix sets the configuration path in the settings provider options to the given path.
func WithConfigPathPrefix(prefix string) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigPathPrefix = prefix
	})
}

// WithConfigPath an additional path to look for settings. Can either be relative to the binary or absolute
func WithConfigPath(path string) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigPath = path
	})
}

func WithConfigurationResolver(resolver SpecConfigurationProvider) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigurationResolver = &resolver
	})
}

func WithDefaultSpecSettings(settings *specv2pb.SpecSettings) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		cfg.DefaultSettings = settings
	})
}

func WithWatchSettings(watch bool) ConfigurationProviderOption {
	return optionFunc(func(cfg *configurationProviderOption) {
		cfg.WatchSettings = watch
	})
}
