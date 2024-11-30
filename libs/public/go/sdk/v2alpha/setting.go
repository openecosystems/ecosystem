package sdkv2alphalib

import (
	"dario.cat/mergo"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

var (
	viperInstance *viper.Viper
	viperOnce     sync.Once
	// Config Global Flags that can override configuration at runtime
	Config    *specv2pb.SpecSettings
	Overrides *RuntimeConfigurationOverrides
)

type SpecSettingsProvider interface {
	GetSettings() *specv2pb.SpecSettings
	WatchSettings() error
}

type DotConfigSettingsProvider struct {
	settings *specv2pb.SpecSettings
}

func NewDotConfigSettingsProvider() (*DotConfigSettingsProvider, error) {

	_, configurer, err := getFileSystemAndConfigurer("")
	if err != nil {
		return nil, err
	}

	c := specv2pb.SpecSettings{}

	err3 := configurer.Unmarshal(&c)
	if err3 != nil {
		fmt.Println("Could not process configuration file")
		return nil, err3
	}

	return &DotConfigSettingsProvider{
		settings: &c,
	}, nil
}

func (p *DotConfigSettingsProvider) GetSettings() *specv2pb.SpecSettings {
	return p.settings
}

func (p *DotConfigSettingsProvider) WatchSettings() error {
	return watchSettings(p.settings, Filesystem.ContextDirectory)
}

type SpecYamlSettingsProvider struct {
	settings *specv2pb.SpecSettings
}

func NewSpecYamlSettingsProvider() (*SpecYamlSettingsProvider, error) {

	viperOnce.Do(
		func() {
			err := godotenv.Load()
			if err != nil {
				//fmt.Println(err.Error())
			}

			viperInstance = viper.New()
			viperInstance.SetConfigName("spec")
			viperInstance.SetConfigType("yaml")
			viperInstance.AddConfigPath("/etc/spec")
			viperInstance.AddConfigPath("$HOME/.spec")
			viperInstance.AddConfigPath(".")
			//viperInstance.AutomaticEnv()
		},
	)

	var c specv2pb.SpecSettings
	Resolve(&c, specv2pb.SpecSettings{})

	return &SpecYamlSettingsProvider{
		settings: &c,
	}, nil
}

func (p *SpecYamlSettingsProvider) GetSettings() *specv2pb.SpecSettings {
	return p.settings
}

func (p *SpecYamlSettingsProvider) WatchSettings() error {
	return watchSettings(p.settings, ".", "/etc/spec")
}

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

type CLISettingsProvider struct {
	settings *specv2pb.SpecSettings
	Flags    *RuntimeConfigurationOverrides
}

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

	err3 := configurer.Unmarshal(&c)
	if err3 != nil {
		fmt.Println("Could not process configuration file")
		return nil, err3
	}

	Config = &c
	Overrides = flags

	return &CLISettingsProvider{
		settings: &c,
	}, nil
}

func (p *CLISettingsProvider) GetSettings() *specv2pb.SpecSettings {
	return p.settings
}

func (p *CLISettingsProvider) WatchSettings() error {
	return watchSettings(p.settings, Filesystem.ContextDirectory)
}

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
			return nil, nil, errors.New("no context found")
		}

		configurer.SetFs(ufs)
		configurer.SetConfigName(ctx)
		configurer.SetConfigType(ConfigurationExtension)
		configurer.AddConfigPath(contextDir)

	}

	// Set Environment Variable Overrides
	configurer.AutomaticEnv()

	if err2 := configurer.ReadInConfig(); err2 != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err2, &configFileNotFoundError) {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found: " + err2.Error())
		}
	}

	return fs, configurer, nil
}

func watchSettings(settings *specv2pb.SpecSettings, directories ...string) error {

	//If dynamic settings enabled, turn on filesystem notification
	if settings.Platform != nil && !settings.Platform.DynamicConfigReload {

		//fmt.Println("Dynamic reload is disabled in your settings. Please enable if you want to Watch Settings: settings.platform.DynamicConfigReload")
		return nil

	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {

		}
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
				for _, e := range settings.Systems2 {
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

type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func setEnv(envPrefix, yaml string) {
	envVar := strcase.ToScreamingSnake(strcase.ToLowerCamel(strings.ReplaceAll(yaml, ".", "_")))
	if envPrefix != "" {
		envVar = envPrefix + "_" + envVar
	}
	// Uncomment to print all variables
	// fmt.Println(envVar + "=")
	val, present := os.LookupEnv(envVar)
	if present {
		fmt.Println(fmt.Sprintf("Setting %s from %s to %s", yaml, envVar, val))
		viperInstance.Set(yaml, val)
	}
}

func ImportEnvironmentVariables(iface interface{}, envPrefix string, prefix string) {
	//https://github.com/spf13/viper/issues/188#issuecomment-399884438
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
				fmt.Println(fmt.Sprintf("Setting %s from %s to %s", fieldName, envVar, val))
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

func ImportPackageJson(config *Configuration) {
	content, err := os.ReadFile("./package.json")
	if err != nil {
		return
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

	config.App.Name = data.Name
	config.App.Version = data.Version
}

func StringExpandEnv() mapstructure.DecodeHookFuncKind {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{}) (interface{}, error) {
		if f != reflect.String || t != reflect.String {
			return data, nil
		}

		return os.ExpandEnv(data.(string)), nil
	}
}

func Resolve(dst, src interface{}) {
	err := viperInstance.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("No spec configuration found.")
		}
	}

	//ImportEnvironmentVariables(config, "", "")

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
	if err := validate.Struct(dst); err != nil {
		fmt.Println("Missing required attributes", err)
	}

	if err := mergo.Merge(dst, src); err != nil {
		fmt.Println("Error merging settings configuration:", err)
	}
}

func IsLower(s string) bool {
	for _, charNumber := range s {
		if charNumber > 122 || charNumber < 97 {
			return false
		}
	}
	return true
}
