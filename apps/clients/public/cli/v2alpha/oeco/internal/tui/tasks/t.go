package tasks

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// SectionIdentifier represents a distinct section within a program, identified by an integer ID and a string Type.
type SectionIdentifier struct {
	ID   int
	Type string
}

// GenericTask represents a task to be executed with associated metadata and a customizable message construction function.
type GenericTask struct {
	ID           string
	Args         []string
	Section      SectionIdentifier
	StartText    string
	FinishedText string
	Msg          func(c *exec.Cmd, err error) tea.Msg
}

// OpenBrowser initiates a task to open a URL in the default web browser and returns a command to handle its lifecycle.
func OpenBrowser() tea.Cmd {
	//taskID := fmt.Sprintf("open_browser_%d", time.Now().Unix())
	//task := Task{
	//	ID:           taskID,
	//	StartText:    "Opening in browser",
	//	FinishedText: "Opened in browser",
	//	State:        TaskStart,
	//	Error:        nil,
	//}
	openCmd := func() tea.Msg {
		// err := browser.OpenURL("https://www.google.com")
		// return constants.TaskFinishedMsg{TaskID: taskID, Err: err}

		return TaskFinishedMsg{}
	}
	return tea.Batch(openCmd)
}
