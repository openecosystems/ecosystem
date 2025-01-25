package commands

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/data"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	"bytes"
	"fmt"
	"maps"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// IssueCommandTemplateInput defines the input parameters required to execute an issue command template in a repository context.
type IssueCommandTemplateInput struct {
	RepoName    string
	RepoPath    string
	IssueNumber int
	HeadRefName string
}

// Model represents the core data structure for handling commands and logic in the application.
type Model struct{}

// runCustomCommand parses a command template with provided context data and executes the resulting command.
//
//nolint:unused
func (m *Model) runCustomCommand(commandTemplate string, contextData *map[string]any) tea.Cmd {
	// A generic map is a pretty easy & flexible way to populate a template if there's no pressing need
	// for structured data, existing structs, etc. Especially if holes in the data are expected.
	// Common data shared across contexts could be set here.
	input := map[string]any{}

	// Merge data specific to the context the command is being run in onto any common data, overwriting duplicate keys.
	if contextData != nil {
		maps.Copy(input, *contextData)
	}

	cmd, err := template.New("keybinding_command").Parse(commandTemplate)
	if err != nil {
		log.Fatal("Failed parse keybinding template", "error", err)
	}

	// Set the command to error out if required input (e.g. RepoPath) is missing
	cmd = cmd.Option("missingkey=error")

	var buff bytes.Buffer
	err = cmd.Execute(&buff, input)
	if err != nil {
		return func() tea.Msg {
			return constants.ErrMsg{Err: fmt.Errorf("failed to parsetemplate %s", commandTemplate)}
		}
	}
	return m.executeCustomCommand(buff.String())
}

// runCustomPRCommand generates and runs a custom command for a pull request using provided template and PR data.
//
//nolint:unused
func (m *Model) runCustomPRCommand(commandTemplate string, prData *data.PullRequestData) tea.Cmd {
	return m.runCustomCommand(commandTemplate,
		&map[string]any{
			"RepoName":    prData.GetRepoNameWithOwner(),
			"PrNumber":    prData.Number,
			"HeadRefName": prData.HeadRefName,
			"BaseRefName": prData.BaseRefName,
		})
}

// execProcessFinishedMsg represents a message signaling the completion of an execution process.
//
//nolint:unused
type execProcessFinishedMsg struct{}

// executeCustomCommand executes a custom shell command and returns a tea.Cmd to handle the process execution.
//
//nolint:unused
func (m *Model) executeCustomCommand(_ string) tea.Cmd {
	//log.Debug("executing custom command", "cmd", cmd)
	//shell := os.Getenv("SHELL")
	//if shell == "" {
	//	shell = "sh"
	//}
	//c := exec.Command(shell, "-c", cmd)
	//return tea.ExecProcess(c, func(err error) tea.Msg {
	//	if err != nil {
	//		mdRenderer := markdown.GetMarkdownRenderer(m.ctx.ScreenWidth)
	//		md, mdErr := mdRenderer.Render(fmt.Sprintf("While running: `%s`", cmd))
	//		if mdErr != nil {
	//			return constants.ErrMsg{Err: mdErr}
	//		}
	//		return constants.ErrMsg{Err: errors.New(
	//			lipgloss.JoinVertical(lipgloss.Left,
	//				fmt.Sprintf("Whoops, got an error: %s", err),
	//				md,
	//			),
	//		)}
	//	}
	//	return execProcessFinishedMsg{}
	//})

	return nil
}

//
//func (m *Model) notify(text string) tea.Cmd {
//	id := fmt.Sprint(time.Now().Unix())
//	startCmd := m.ctx.StartTask(
//		context.Task{
//			Id:           id,
//			StartText:    text,
//			FinishedText: text,
//			State:        context.TaskStart,
//		})
//
//	finishCmd := func() tea.Msg {
//		return constants.TaskFinishedMsg{
//			TaskId: id,
//		}
//	}
//
//	return tea.Sequence(startCmd, finishCmd)
//}
//
//func (m *Model) notifyErr(text string) tea.Cmd {
//	id := fmt.Sprint(time.Now().Unix())
//	startCmd := m.ctx.StartTask(
//		context.Task{
//			Id:           id,
//			StartText:    text,
//			FinishedText: text,
//			State:        context.TaskStart,
//		})
//
//	finishCmd := func() tea.Msg {
//		return constants.TaskFinishedMsg{
//			TaskId: id,
//			Err:    errors.New(text),
//		}
//	}
//
//	return tea.Sequence(startCmd, finishCmd)
//}
