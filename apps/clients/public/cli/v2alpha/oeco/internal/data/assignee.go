package data

// Assignees represents a collection of assignees, each represented by the Assignee type.
type Assignees struct {
	Nodes []Assignee
}

// Assignee represents a user or entity assigned to a specific task or issue within a project or system.
type Assignee struct {
	Login string
}
