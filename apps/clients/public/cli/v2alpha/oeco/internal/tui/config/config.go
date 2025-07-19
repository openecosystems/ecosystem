package config

import (
	"path/filepath"

	"github.com/go-playground/validator/v10"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

// TuiConfigurationFileName defines the default name of the configuration file used for TUI-related settings.
const (
	TuiConfigurationFileName = "tui.yaml"
)

// validate is a pointer to a validator.Validate instance used for input validation tasks.
// TuiConfigurationFile represents the full file path of the TUI configuration file within the configuration directory.
var (
	validate             *validator.Validate
	TuiConfigurationFile = filepath.Join(sdkv2betalib.ConfigurationDirectory, TuiConfigurationFileName)
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
// HomePage represents a page type for a home page.
// ConnectorDetailsPage represents a page type for connector details.
// ConnectorLogsPage represents a page type for connector logs.
// ConnectorRequestsPage represents a page type for connector requests.
// ConnectorPacketsPage represents a page type for connector packets.
// APIExplorerListPage represents a page type for the API explorer list.
// EcosystemDashboardPage represents a page type for the Ecosystem Dashboard.
const (
	EmptyPage              PageType = "empty"
	HomePage               PageType = "home"
	ConnectorDetailsPage   PageType = "connector_details"
	ConnectorLogsPage      PageType = "connector_logs"
	ConnectorRequestsPage  PageType = "connector_requests"
	ConnectorPacketsPage   PageType = "connector_packets"
	APIExplorerListPage    PageType = "api_explorer_list"
	EcosystemCreatePage    PageType = "ecosystem_create"
	EcosystemDashboardPage PageType = "ecosystem_dashboard"
)

// Defaults defines the default configuration settings for sections, pages, and sidebar, including refresh intervals and date format.
type Defaults struct {
	Section SectionType   `yaml:"section"`
	Page    PageType      `yaml:"page"`
	Sidebar SidebarConfig `yaml:"sidebar"`

	Layout LayoutConfig `yaml:"layout,omitempty"`

	StreamMaxRecordsToRetain int    `yaml:"streamMaxRecordsToRetain,omitempty"`
	RefetchIntervalMinutes   int    `yaml:"refetchIntervalMinutes,omitempty"`
	DateFormat               string `yaml:"dateFormat,omitempty"`
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

// LayoutConfig defines the configuration structure for layout settings, including packet container configurations.
type LayoutConfig struct {
	Packets PacketContainerLayoutConfig `yaml:"issues,omitempty"`
}

// ColumnConfig represents the configuration for a column, including width and visibility properties.
type ColumnConfig struct {
	Width  *int  `yaml:"width,omitempty"  validate:"omitempty,gt=0"`
	Hidden *bool `yaml:"hidden,omitempty"`
}
