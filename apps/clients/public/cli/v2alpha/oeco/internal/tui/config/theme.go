package config

type HexColor string

type ColorThemeText struct {
	Primary   HexColor `yaml:"primary"   validate:"omitempty,hexcolor"`
	Secondary HexColor `yaml:"secondary" validate:"omitempty,hexcolor"`
	Inverted  HexColor `yaml:"inverted"  validate:"omitempty,hexcolor"`
	Faint     HexColor `yaml:"faint"     validate:"omitempty,hexcolor"`
	Warning   HexColor `yaml:"warning"   validate:"omitempty,hexcolor"`
	Success   HexColor `yaml:"success"   validate:"omitempty,hexcolor"`
	Error     HexColor `yaml:"error"     validate:"omitempty,hexcolor"`
}

type ColorThemeBorder struct {
	Primary   HexColor `yaml:"primary"   validate:"omitempty,hexcolor"`
	Secondary HexColor `yaml:"secondary" validate:"omitempty,hexcolor"`
	Faint     HexColor `yaml:"faint"     validate:"omitempty,hexcolor"`
}

type ColorThemeBackground struct {
	Selected HexColor `yaml:"selected" validate:"omitempty,hexcolor"`
}

type ColorTheme struct {
	Text       ColorThemeText       `yaml:"text"       validate:"required"`
	Background ColorThemeBackground `yaml:"background" validate:"required"`
	Border     ColorThemeBorder     `yaml:"border"     validate:"required"`
}

type ColorThemeConfig struct {
	Inline ColorTheme `yaml:",inline"`
}

type TableUIThemeConfig struct {
	ShowSeparator bool `yaml:"showSeparator" default:"true"`
	Compact       bool `yaml:"compact" default:"false"`
}

type TuiThemeConfig struct {
	Table TableUIThemeConfig `yaml:"table"`
}

type ThemeConfig struct {
	Tui    TuiThemeConfig    `yaml:"tui,omitempty"     validate:"omitempty"`
	Colors *ColorThemeConfig `yaml:"colors,omitempty" validate:"omitempty"`
}
