package constants

import tea "github.com/charmbracelet/bubbletea"

// TaskFinishedMsg represents a message indicating the completion of a task, including task details and an optional error.
type TaskFinishedMsg struct {
	TaskId      string
	SectionId   int
	SectionType string
	Err         error
	Msg         tea.Msg
}

// ClearTaskMsg represents a message used to clear or remove a specific task identified by its TaskId.
type ClearTaskMsg struct {
	TaskId string
}
