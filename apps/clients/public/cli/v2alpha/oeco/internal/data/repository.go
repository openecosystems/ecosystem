package data

// Repository represents a Git repository with its name, owner information, and archival status.
type Repository struct {
	Name          string
	NameWithOwner string
	IsArchived    bool
}
