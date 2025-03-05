package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// Parser is a structure that holds a reference to a FileSystem for interacting with the underlying filesystem.
type Parser struct {
	Filesystem *sdkv2alphalib.FileSystem
}

// getDefaultConfig returns a Config object initialized with the default configuration settings for the application.
func (parser Parser) getDefaultConfig() Config {
	return Config{
		Sections: []SectionConfig{
			{
				Title:       "Connectors",
				Description: "Get connected",
				Type:        ConnectorSection,
				Pages: []PageConfig{
					{
						Title:       "Connector Details",
						Description: "",
						Type:        ConnectorDetailsPage,
					},
				},
			},
		},
		Defaults: Defaults{
			Section: ConnectorSection,
			Page:    ConnectorDetailsPage,
			Sidebar: SidebarConfig{
				Open:  false,
				Width: 50,
			},
			RefetchIntervalMinutes: 30,
		},
		KeyBindings: KeyBindings{
			Universal:    []KeyBinding{},
			Enclave:      []KeyBinding{},
			Context:      []KeyBinding{},
			Organization: []KeyBinding{},
			Connector:    []KeyBinding{},
			API:          []KeyBinding{},
			Ecosystem:    []KeyBinding{},
		},
		Theme: &ThemeConfig{
			Tui: TuiThemeConfig{
				Table: TableUIThemeConfig{
					ShowSeparator: true,
					Compact:       false,
				},
			},
		},
		Pager:       Pager{},
		ConfirmQuit: true,
	}
}

// getDefaultConfigYamlContents serializes the default configuration into a YAML string format and returns it.
func (parser Parser) getDefaultConfigYamlContents() string {
	defaultConfig := parser.getDefaultConfig()
	y, _ := yaml.Marshal(defaultConfig)

	return string(y)
}

// getDefaultConfigFileOrCreateIfMissing checks if the default config file exists, creates it with default content if missing.
func (parser Parser) getDefaultConfigFileOrCreateIfMissing() (string, error) {
	fs := *parser.Filesystem

	exists, err1 := fs.Exists(TuiConfigurationFile)
	if err1 != nil {
		return "", configError{parser: parser, configDir: sdkv2alphalib.ConfigurationDirectory, err: err1}
	}

	if !exists {
		err2 := fs.CreateFile(TuiConfigurationFile)
		if err2 != nil {
			return "", configError{parser: parser, configDir: sdkv2alphalib.ConfigurationDirectory, err: err2}
		}

		defaultConfig := parser.getDefaultConfig()
		y, _ := yaml.Marshal(defaultConfig)

		err3 := fs.WriteFile(TuiConfigurationFile, y, os.ModePerm)
		if err3 != nil {
			return "", configError{parser: parser, configDir: sdkv2alphalib.ConfigurationDirectory, err: err3}
		}
	}

	return TuiConfigurationFile, nil
}

// readConfigFile reads and parses the configuration file from the specified path.
// It returns the parsed Config object and an error if the operation fails.
func (parser Parser) readConfigFile(path string) (Config, error) {
	config := parser.getDefaultConfig()

	data, err := parser.Filesystem.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	err = validate.Struct(config)
	return config, err
}

// initParser initializes and returns a new Parser object with configured filesystem and validation settings.
func initParser() Parser {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("yaml"), ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return Parser{
		Filesystem: sdkv2alphalib.NewFileSystem(),
	}
}

// ParseConfig reads and parses the application's configuration file or generates a default configuration if missing.
// Returns a Config struct and an error if parsing fails or configuration is invalid.
func ParseConfig() (Config, error) {
	parser := initParser()

	var config Config
	var err error
	var configFilePath string

	configFilePath, err = parser.getDefaultConfigFileOrCreateIfMissing()
	if err != nil {
		return config, parsingError{err: err}
	}

	config, err = parser.readConfigFile(configFilePath)
	if err != nil {
		return config, parsingError{err: err}
	}

	return config, nil
}
