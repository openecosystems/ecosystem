package config

// SectionConfig represents the configuration for a specific section, including its title, description, type, and pages.
type SectionConfig struct {
	Title       string
	Description string
	Type        SectionType
	Pages       []PageConfig
}

// PageConfig represents the configuration for a single page, including title, description, type, and sidebar settings.
type PageConfig struct {
	Title       string
	Description string
	Type        PageType
	Sidebar     SidebarConfig
}

// SidebarConfig defines the configuration options for a sidebar, including its visibility and width properties.
type SidebarConfig struct {
	Open  bool
	Width int
}
