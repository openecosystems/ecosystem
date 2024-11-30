package config

type SectionConfig struct {
	Title       string
	Description string
	Type        SectionType
	Pages       []PageConfig
}

type PageConfig struct {
	Title       string
	Description string
	Type        PageType
	Sidebar     SidebarConfig
}

type SidebarConfig struct {
	Open  bool
	Width int
}
