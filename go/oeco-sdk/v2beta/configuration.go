//nolint:revive
package sdkv2betalib

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

	"dario.cat/mergo"
	"github.com/apex/log"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"
)

const (

	// EdgePrefixConfiguration defines the prefix used for edge configuration settings within the system.
	EdgePrefixConfiguration = "edge"

	// ApiPrefixConfiguration defines the prefix used for API-related configurations.
	ApiPrefixConfiguration = "api"
)

// once ensures viperInstance is initialized only once in a thread-safe manner.
// Config holds global settings for the application, managed via SpecSettings.
// Overrides contains runtime-specific configuration that can override the default settings.
var (
	Overrides *RuntimeConfigurationOverrides

	once sync.Once
)

// BaseSpecConfigurationProvider Non-generic interface (parent)
type BaseSpecConfigurationProvider interface {
	ValidateConfiguration() error
	WatchConfigurations(directories ...string) error
	ResolveConfiguration(opts ...ConfigurationProviderOption) (*Configurer, error)
	GetConfigurationBytes() ([]byte, error)
}

// SpecConfigurationProvider defines an interface for obtaining and monitoring specification settings.
// GetConfigurations retrieves the current specification settings.
// WatchConfigurations initiates a watching mechanism to monitor settings updates.
type SpecConfigurationProvider[P any] interface {
	BaseSpecConfigurationProvider // Embed non-generic base interface

	CreateConfiguration() (*P, error)
	GetConfiguration() *P
	GetDefaultConfiguration() *P
}

// Configurer manages and provides application server configuration settings.
type Configurer struct {
	Filesystem *FileSystem
	Configurer *viper.Viper
	Cfg        *configurationProviderOption
	Options    []ConfigurationProviderOption
}

// NewConfigurer initializes a new ConfigurationProvider with default SpecSettings configuration.
func NewConfigurer(opts ...ConfigurationProviderOption) (*Configurer, error) {
	once.Do(
		func() {
			_ = godotenv.Load()
			// viperInstance.AutomaticEnv()
		},
	)

	sctx, err := initializeConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	if sctx.Cfg.ConfigName != "" {
		sctx.Configurer.SetConfigName(sctx.Cfg.ConfigName)
	} else {
		sctx.Configurer.SetConfigName(getConfigFileName(sctx.Cfg.ConfigPathPrefix, sctx.Cfg.PlatformContext))
	}

	if sctx.Cfg.ConfigExtension != "" {
		sctx.Configurer.SetConfigType(sctx.Cfg.ConfigExtension)
	} else {
		sctx.Configurer.SetConfigType(ConfigurationExtension)
	}

	if sctx.Cfg.ConfigPath != "" {
		sctx.Configurer.AddConfigPath(sctx.Cfg.ConfigPath)
	} else {
		sctx.Configurer.AddConfigPath(ConfigurationDirectory)
	}
	sctx.Configurer.AddConfigPath(".")

	return &Configurer{
		Filesystem: sctx.Filesystem,
		Configurer: sctx.Configurer,
		Cfg:        sctx.Cfg,
		Options:    opts,
	}, nil
}

func WatchConfigurations(directories ...string) error {
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
				//err = _r.WatchConfigurationsHandler(event)
				//if err != nil {
				//	return err
				//}

				// Update settings
				//for _, e := range settings.Systems {
				//	fmt.Println(e.Name + ": " + e.Version)
				//}
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

// initializeConfigurer creates and initializes a FileSystem and viper.Viper configurer, handling context overrides and config file setup.
//
//nolint:unparam
func initializeConfigurer(opts ...ConfigurationProviderOption) (*Configurer, error) {
	var _cfg configurationProviderOption
	for _, opt := range opts {
		opt.apply(&_cfg)
	}

	configurer := _cfg.Configurer
	if configurer == nil {
		configurer = &Configurer{
			Filesystem: NewFileSystem(),
			Configurer: viper.New(),
			Cfg:        &_cfg,
			Options:    opts,
		}
	}

	cfg := configurer.Cfg

	filesystem := configurer.Filesystem

	var ctx string
	if cfg.PlatformContext != "" {
		ctx = cfg.PlatformContext
	}

	if cfg.RuntimeConfigurationOverrides.Context != "" {
		ctx = cfg.RuntimeConfigurationOverrides.Context
	}

	if ctx != "" {
		_, err := filesystem.Exists(filepath.Join(ContextDirectory, ctx))
		if err != nil {
			fmt.Println("Error: context does not exists: " + err.Error())
			return nil, errors.New("context does not exists")
		}
	} else {
		file, err := filesystem.ReadFile(DefaultContextFile)
		if err != nil {
			return nil, errors.New("could not read config file: " + err.Error())
		}

		ctx = strings.TrimSpace(string(file))

		if ctx == "" {
			// Set the oeco workspace in the "default" file
			err = filesystem.WriteFile(DefaultContextFile, []byte(OecoContextFileName), os.ModePerm)
			if err != nil {
				return nil, errors.New("internal error: Cannot create default context")
			}

			_, err = createDefaultContextSettings(OecoContextFileName, DefaultCIDR)
			if err != nil {
				return nil, err
			}
		}
	}

	WithConfigurationPlatformContext(ctx).apply(cfg)

	return &Configurer{
		Filesystem: filesystem,
		Configurer: configurer.Configurer,
		Cfg:        cfg,
		Options:    opts,
	}, nil
}

func createDefaultContextSettings(ecosystemName string, cidr string) (*specv2pb.SpecSettings, error) {
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

	return &settings, nil
}

// PackageJson represents the structure of a package.json file containing name and version information.
type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// setEnv maps and sets an environment variable value to a corresponding YAML configuration field using Viper.
// It converts the YAML field into an environment variable format, optionally prepending a prefix if provided.
// If the environment variable is present, its value is used to update the corresponding Viper configuration field.
func setEnv(configurer *viper.Viper, envPrefix, yaml string) {
	envVar := strcase.ToScreamingSnake(strcase.ToLowerCamel(strings.ReplaceAll(yaml, ".", "_")))
	if envPrefix != "" {
		envVar = envPrefix + "_" + envVar
	}
	// Uncomment to print all variables
	// fmt.Println(envVar + "=")
	val, present := os.LookupEnv(envVar)
	if present {
		// fmt.Printf("Setting %s from %s to %s\n", yaml, envVar, val)
		configurer.Set(yaml, val)
	}
}

// importEnvironmentVariables sets configuration values from environment variables based on struct fields and provided tags.
// It processes "mapstructure" and "env" tags, supports nested structs, and recursively applies prefix logic for key formatting.
func importEnvironmentVariables(configurer *viper.Viper, iface interface{}, envPrefix string, prefix string) {
	// https://github.com/spf13/viper/issues/188#issuecomment-399884438
	ifv := reflect.ValueOf(iface)

	// Check if it's a pointer and dereference it
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}

	// Ensure it's a struct before calling NumField
	if ifv.Kind() != reflect.Struct {
		log.Debugf("Expected struct, got %v", ifv.Kind())
		return
	}

	ift := reflect.TypeOf(iface)
	// Ensure it's a struct before calling NumField

	// Check if it's a pointer and dereference it
	if ift.Kind() == reflect.Ptr {
		ift = ift.Elem()
	}

	if ift.Kind() != reflect.Struct {
		log.Debugf("Expected struct, got %v", ift.Kind())
		return
	}

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		// Skip unexported fields
		if !t.IsExported() {
			continue
		}

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
				configurer.Set(fieldName, val)
				continue
			}
		}

		switch v.Kind() {
		case reflect.Struct:
			instance := reflect.New(t.Type)
			if _, ok := instance.Interface().(encoding.TextUnmarshaler); ok {
				setEnv(configurer, envPrefix, fieldName)
			} else {
				if !IsLower(fieldName) {
					importEnvironmentVariables(configurer, v.Interface(), envPrefix, fieldName)
				}
			}
		default:
			setEnv(configurer, envPrefix, fieldName)
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
func Resolve(_configurer *Configurer, destinationConfiguration, defaultConfiguration interface{}) {
	if _configurer == nil {
		panic("ConfigurationProvider is nil")
	}

	configurer := _configurer.Configurer

	err := configurer.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("No spec configuration found. Falling back to defaults.")
		} else {
			panic(fmt.Errorf("error reading config: %w", err))
		}
	}

	if defaultConfiguration != nil {
		err = mergo.Merge(destinationConfiguration, defaultConfiguration, mergo.WithOverride)
		if err != nil {
			fmt.Println("Error merging settings configuration:", err)
		}
	} else {
		fmt.Println("No default configuration provided")
	}

	if err = mergo.Merge(destinationConfiguration, defaultConfiguration); err != nil {
		fmt.Println("Error merging settings configuration:", err)
	}

	importEnvironmentVariables(configurer, destinationConfiguration, "", "")

	err = configurer.Unmarshal(destinationConfiguration, viper.DecodeHook(
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
	if err = validate.Struct(destinationConfiguration); err != nil {
		fmt.Println("Missing required attributes", err)
	}
}

// Merge merge two configurations
func Merge(dst, src interface{}) {
	if err := mergo.Merge(dst, src); err != nil {
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

// RuntimeConfigurationOverrides holds runtime configuration flags and settings that override default behavior.
type RuntimeConfigurationOverrides struct {
	Context      string
	Debug        bool
	Verbose      bool
	VerboseLog   bool
	LogFile      string
	Quiet        bool
	FieldMask    string
	ValidateOnly bool

	Overridden bool
}

// configurationProviderOption is the configuration for a Server.
type configurationProviderOption struct {
	PlatformContext               string
	ConfigPath                    string
	ConfigPathPrefix              string
	ConfigName                    string
	ConfigExtension               string
	WatchSettings                 bool
	RuntimeConfigurationOverrides RuntimeConfigurationOverrides
	Configurer                    *Configurer
}

// ConfigurationProviderOption is an interface for applying configurations to a config object.
type ConfigurationProviderOption interface {
	apply(*configurationProviderOption)
}

type configurationOptionFunc func(*configurationProviderOption)

func (f configurationOptionFunc) apply(cfg *configurationProviderOption) { f(cfg) }

// WithOptions composes multiple Options into one.
func WithOptions(opts ...ConfigurationProviderOption) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

func WithConfigurationPlatformContext(context string) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.PlatformContext = context
	})
}

// WithConfigPathPrefix sets the configuration path in the settings provider options to the given path.
func WithConfigPathPrefix(prefix string) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigPathPrefix = prefix
	})
}

// WithConfigPath an additional path to look for settings. Can either be relative to the binary or absolute
func WithConfigPath(path string) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigPath = path
	})
}

func WithConfigName(name string) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigName = name
	})
}

// WithConfigExtension sets the configuration file extension and applies it to the configuration provider options.
// Do not include a dot/period prefix. For example, yaml is the extension, not .yaml
func WithConfigExtension(extension string) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.ConfigExtension = extension
	})
}

func WithWatchSettings(watch bool) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.WatchSettings = watch
	})
}

func WithRuntimeOverrides(overrides RuntimeConfigurationOverrides) ConfigurationProviderOption {
	overrides.Overridden = true
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.RuntimeConfigurationOverrides = overrides
	})
}

// WithConfigurer an additional path to look for settings. Can either be relative to the binary or absolute
func WithConfigurer(configurer *Configurer) ConfigurationProviderOption {
	return configurationOptionFunc(func(cfg *configurationProviderOption) {
		cfg.Configurer = configurer
	})
}
