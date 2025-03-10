package config

// PacketContainerConfig defines the configuration for a packet container with title, filters, optional limit, and layout.
type PacketContainerConfig struct {
	Title   string
	Filters string
	Limit   *int                        `yaml:"limit,omitempty"`
	Layout  PacketContainerLayoutConfig `yaml:"layout,omitempty"`
}

// PacketContainerLayoutConfig layout configuration
type PacketContainerLayoutConfig struct {
	UpdatedAt ColumnConfig `yaml:"updatedAt,omitempty"`
	CreatedAt ColumnConfig `yaml:"createdAt,omitempty"`
	State     ColumnConfig `yaml:"state,omitempty"`
	Repo      ColumnConfig `yaml:"repo,omitempty"`
	Title     ColumnConfig `yaml:"title,omitempty"`
	Creator   ColumnConfig `yaml:"creator,omitempty"`
	Assignees ColumnConfig `yaml:"assignees,omitempty"`
	Comments  ColumnConfig `yaml:"comments,omitempty"`
	Reactions ColumnConfig `yaml:"reactions,omitempty"`
}
