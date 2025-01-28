package config

// HexColor represents a string value expected to be a valid hexadecimal color code, e.g., #RRGGBB or #RRGGBBAA.
type HexColor string

// ColorThemeText defines a set of text colors for a color theme with optional hex color validation.
// Each field represents a specific type of color used across various UI components.
type ColorThemeText struct {
	Primary   HexColor `yaml:"primary"   validate:"omitempty,hexcolor"`
	Secondary HexColor `yaml:"secondary" validate:"omitempty,hexcolor"`
	Inverted  HexColor `yaml:"inverted"  validate:"omitempty,hexcolor"`
	Faint     HexColor `yaml:"faint"     validate:"omitempty,hexcolor"`
	Warning   HexColor `yaml:"warning"   validate:"omitempty,hexcolor"`
	Success   HexColor `yaml:"success"   validate:"omitempty,hexcolor"`
	Error     HexColor `yaml:"error"     validate:"omitempty,hexcolor"`
}

// ColorThemeBorder represents the color scheme for component borders in the theme configuration.
// Primary defines the main border color in hex format.
// Secondary defines the secondary border color in hex format.
// Faint defines a lighter or less prominent border color in hex format.
type ColorThemeBorder struct {
	Primary   HexColor `yaml:"primary"   validate:"omitempty,hexcolor"`
	Secondary HexColor `yaml:"secondary" validate:"omitempty,hexcolor"`
	Faint     HexColor `yaml:"faint"     validate:"omitempty,hexcolor"`
}

// ColorThemeBackground represents the background colors for a color theme.
// Selected defines the selected background color in hexadecimal format, validated as an optional hex color value.
type ColorThemeBackground struct {
	Selected HexColor `yaml:"selected" validate:"omitempty,hexcolor"`
}

// ColorTheme represents a collection of theming options, including text, background, and border color configurations.
type ColorTheme struct {
	Text       ColorThemeText       `yaml:"text"       validate:"required"`
	Background ColorThemeBackground `yaml:"background" validate:"required"`
	Border     ColorThemeBorder     `yaml:"border"     validate:"required"`
}

// ColorThemeConfig represents the configuration for color themes, embedding a ColorTheme structure as an inline field.
type ColorThemeConfig struct {
	Inline ColorTheme `yaml:",inline"`
}

// TableUIThemeConfig defines the configuration options for the Table UI theme, including separator visibility and compact mode.
type TableUIThemeConfig struct {
	ShowSeparator bool `yaml:"showSeparator" default:"true"`
	Compact       bool `yaml:"compact" default:"false"`
}

// TuiThemeConfig defines the theme configuration for the text-based user interface.
type TuiThemeConfig struct {
	Table TableUIThemeConfig `yaml:"table"`
}

// ThemeConfig represents configuration for theming, including TUI-related themes and color-based themes.
// Tui defines the theming configuration for textual user interfaces.
// Colors provides optional theme customization for colors and their attributes.
type ThemeConfig struct {
	Tui    TuiThemeConfig    `yaml:"tui,omitempty"     validate:"omitempty"`
	Colors *ColorThemeConfig `yaml:"colors,omitempty" validate:"omitempty"`
}
