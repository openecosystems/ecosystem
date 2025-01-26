package config

import (
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

// TuiConfigurationFileName defines the default name of the configuration file used for TUI-related settings.
const (
	TuiConfigurationFileName = "tui.yaml"
)

// validate is a pointer to a validator.Validate instance used for input validation tasks.
// TuiConfigurationFile represents the full file path of the TUI configuration file within the configuration directory.
var (
	validate             *validator.Validate
	TuiConfigurationFile = filepath.Join(sdkv2alphalib.ConfigurationDirectory, TuiConfigurationFileName)
)

// SectionType represents the type of a section, used to categorize and organize content or configurations.
type SectionType string

// EnclaveSection represents the section type for an enclave.
// ContextSection represents the section type for a context.
// OrganizationSection represents the section type for an organization.
// PackageSection represents the section type for a package.
// ConnectorSection represents the section type for a connector.
// ApiSection represents the section type for an API.
// EcosystemSection represents the section type for an ecosystem.
const (
	EnclaveSection      SectionType = "enclave"
	ContextSection      SectionType = "context"
	OrganizationSection SectionType = "organization"
	PackageSection      SectionType = "package"
	ConnectorSection    SectionType = "connector"
	APISection          SectionType = "api"
	EcosystemSection    SectionType = "ecosystem"
)

// PageType represents the type of a page within the configuration or application context.
type PageType string

// EmptyPage represents a page type for an empty page.
// ConnectorDetailsPage represents a page type for connector details.
// ConnectorLogsPage represents a page type for connector logs.
// ConnectorRequestsPage represents a page type for connector requests.
// ConnectorPacketsPage represents a page type for connector packets.
// APIExplorerListPage represents a page type for the API explorer list.
const (
	EmptyPage             PageType = "empty"
	ConnectorDetailsPage  PageType = "connector_details"
	ConnectorLogsPage     PageType = "connector_logs"
	ConnectorRequestsPage PageType = "connector_requests"
	ConnectorPacketsPage  PageType = "connector_packets"
	APIExplorerListPage   PageType = "api_explorer_list"
)

// Defaults defines the default configuration settings for sections, pages, and sidebar, including refresh intervals and date format.
type Defaults struct {
	Section SectionType   `yaml:"section"`
	Page    PageType      `yaml:"page"`
	Sidebar SidebarConfig `yaml:"sidebar"`

	RefetchIntervalMinutes int    `yaml:"refetchIntervalMinutes,omitempty"`
	DateFormat             string `yaml:"dateFormat,omitempty"`
}

// Pager represents the configuration for paging functionality with a specific diff style.
type Pager struct {
	Diff string `yaml:"diff"`
}

// Config represents the main configuration structure for the application.
type Config struct {
	Sections    []SectionConfig `yaml:"sections"`
	Defaults    Defaults        `yaml:"defaults"`
	KeyBindings KeyBindings     `yaml:"keyBindings"`
	Theme       *ThemeConfig    `yaml:"theme,omitempty" validate:"omitempty"`
	Pager       Pager           `yaml:"pager"`
	ConfirmQuit bool            `yaml:"confirmQuit"`
}
