package config

import (
	"libs/public/go/sdk/v2alpha"
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

const (
	TuiConfigurationFileName = "tui.yaml"
)

var (
	validate             *validator.Validate
	TuiConfigurationFile = filepath.Join(sdkv2alphalib.ConfigurationDirectory, TuiConfigurationFileName)
)

type SectionType string

const (
	EnclaveSection      SectionType = "enclave"
	ContextSection      SectionType = "context"
	OrganizationSection SectionType = "organization"
	PackageSection      SectionType = "package"
	ConnectorSection    SectionType = "connector"
	ApiSection          SectionType = "api"
	EcosystemSection    SectionType = "ecosystem"
)

type PageType string

const (
	EmptyPage             PageType = "empty"
	ConnectorDetailsPage  PageType = "connector_details"
	ConnectorLogsPage     PageType = "connector_logs"
	ConnectorRequestsPage PageType = "connector_requests"
	ConnectorPacketsPage  PageType = "connector_packets"
	ApiExplorerListPage   PageType = "api_explorer_list"
)

type Defaults struct {
	Section SectionType   `yaml:"section"`
	Page    PageType      `yaml:"page"`
	Sidebar SidebarConfig `yaml:"sidebar"`

	RefetchIntervalMinutes int    `yaml:"refetchIntervalMinutes,omitempty"`
	DateFormat             string `yaml:"dateFormat,omitempty"`
}

type Pager struct {
	Diff string `yaml:"diff"`
}

type Config struct {
	Sections    []SectionConfig `yaml:"sections"`
	Defaults    Defaults        `yaml:"defaults"`
	KeyBindings KeyBindings     `yaml:"keyBindings"`
	Theme       *ThemeConfig    `yaml:"theme,omitempty" validate:"omitempty"`
	Pager       Pager           `yaml:"pager"`
	ConfirmQuit bool            `yaml:"confirmQuit"`
}
