package tasks

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/data"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
	"fmt"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// SectionIdentifier represents a distinct section within a program, identified by an integer Id and a string Type.
type SectionIdentifier struct {
	Id   int
	Type string
}

// UpdatePRMsg represents a message used to update the state or properties of a pull request in a system.
type UpdatePRMsg struct {
	PrNumber         int
	IsClosed         *bool
	NewComment       *data.Comment
	ReadyForReview   *bool
	IsMerged         *bool
	AddedAssignees   *data.Assignees
	RemovedAssignees *data.Assignees
}

// UpdateBranchMsg represents a message used for updating branch-related data, including branch name, creation status, and a pull request.
type UpdateBranchMsg struct {
	Name      string
	IsCreated *bool
	NewPr     *data.PullRequestData
}

// buildTaskId generates a unique task identifier by combining the given prefix and pull request number.
func buildTaskId(prefix string, prNumber int) string {
	return fmt.Sprintf("%s_%d", prefix, prNumber)
}

// GenericTask represents a task to be executed with associated metadata and a customizable message construction function.
type GenericTask struct {
	Id           string
	Args         []string
	Section      SectionIdentifier
	StartText    string
	FinishedText string
	Msg          func(c *exec.Cmd, err error) tea.Msg
}

// fireTask executes a task within a given program context, managing its lifecycle and returning a command for processing.
// It initializes the task state, executes associated shell commands, and handles the completion response.
func fireTask(ctx *context.ProgramContext, task GenericTask) tea.Cmd {
	start := context.Task{
		ID:           task.Id,
		StartText:    task.StartText,
		FinishedText: task.FinishedText,
		State:        context.TaskStart,
		Error:        nil,
	}

	startCmd := ctx.StartTask(start)
	return tea.Batch(startCmd, func() tea.Msg {
		log.Debug("Running task", "cmd", "gh "+strings.Join(task.Args, " "))
		c := exec.Command("gh", task.Args...)

		err := c.Run()
		return constants.TaskFinishedMsg{
			TaskID:      task.Id,
			SectionID:   task.Section.Id,
			SectionType: task.Section.Type,
			Err:         err,
			Msg:         task.Msg(c, err),
		}
	})
}

// OpenBrowser initiates a task to open a URL in the default web browser and returns a command to handle its lifecycle.
func OpenBrowser(ctx *context.ProgramContext) tea.Cmd {
	taskID := fmt.Sprintf("open_browser_%d", time.Now().Unix())
	task := context.Task{
		ID:           taskID,
		StartText:    "Opening in browser",
		FinishedText: "Opened in browser",
		State:        context.TaskStart,
		Error:        nil,
	}
	startCmd := ctx.StartTask(task)
	openCmd := func() tea.Msg {
		// err := browser.OpenURL("https://www.google.com")
		// return constants.TaskFinishedMsg{TaskID: taskID, Err: err}

		return constants.TaskFinishedMsg{TaskID: taskID, Err: nil}
	}
	return tea.Batch(startCmd, openCmd)
}

// OpenBranchPR initiates a task to open a pull request for a specified branch and section using the GitHub CLI.
func OpenBranchPR(ctx *context.ProgramContext, section SectionIdentifier, branch string) tea.Cmd {
	return fireTask(ctx, GenericTask{
		Id: fmt.Sprintf("branch_open_%s", branch),
		Args: []string{
			"pr",
			"view",
			"--web",
			branch,
			"-R",
		},
		Section:      section,
		StartText:    fmt.Sprintf("Opening PR for branch %s", branch),
		FinishedText: fmt.Sprintf("PR for branch %s has been opened", branch),
		Msg: func(c *exec.Cmd, err error) tea.Msg {
			return UpdatePRMsg{}
		},
	})
}

// ReopenPR reopens a pull request with the given PR number and repository, updating the state and triggering necessary tasks.
func ReopenPR(ctx *context.ProgramContext, section SectionIdentifier, pr data.RowData) tea.Cmd {
	prNumber := pr.GetNumber()
	return fireTask(ctx, GenericTask{
		Id: buildTaskId("pr_reopen", prNumber),
		Args: []string{
			"pr",
			"reopen",
			fmt.Sprint(prNumber),
			"-R",
			pr.GetRepoNameWithOwner(),
		},
		Section:      section,
		StartText:    fmt.Sprintf("Reopening PR #%d", prNumber),
		FinishedText: fmt.Sprintf("PR #%d has been reopened", prNumber),
		Msg: func(c *exec.Cmd, err error) tea.Msg {
			return UpdatePRMsg{
				PrNumber: prNumber,
				IsClosed: utils.BoolPtr(false),
			}
		},
	})
}

// ClosePR initiates the process of closing a pull request and returns a command to manage its lifecycle.
func ClosePR(ctx *context.ProgramContext, section SectionIdentifier, pr data.RowData) tea.Cmd {
	prNumber := pr.GetNumber()
	return fireTask(ctx, GenericTask{
		Id: buildTaskId("pr_close", prNumber),
		Args: []string{
			"pr",
			"close",
			fmt.Sprint(prNumber),
			"-R",
			pr.GetRepoNameWithOwner(),
		},
		Section:      section,
		StartText:    fmt.Sprintf("Closing PR #%d", prNumber),
		FinishedText: fmt.Sprintf("PR #%d has been closed", prNumber),
		Msg: func(c *exec.Cmd, err error) tea.Msg {
			return UpdatePRMsg{
				PrNumber: prNumber,
				IsClosed: utils.BoolPtr(true),
			}
		},
	})
}

// PRReady marks a pull request as ready for review and triggers an asynchronous task with related messages.
func PRReady(ctx *context.ProgramContext, section SectionIdentifier, pr data.RowData) tea.Cmd {
	prNumber := pr.GetNumber()
	return fireTask(ctx, GenericTask{
		Id: buildTaskId("pr_ready", prNumber),
		Args: []string{
			"pr",
			"ready",
			fmt.Sprint(prNumber),
			"-R",
			pr.GetRepoNameWithOwner(),
		},
		Section:      section,
		StartText:    fmt.Sprintf("Marking PR #%d as ready for review", prNumber),
		FinishedText: fmt.Sprintf("PR #%d has been marked as ready for review", prNumber),
		Msg: func(c *exec.Cmd, err error) tea.Msg {
			return UpdatePRMsg{
				PrNumber:       prNumber,
				ReadyForReview: utils.BoolPtr(true),
			}
		},
	})
}

// MergePR attempts to merge a pull request for a given repository using the "gh" CLI and tracks its progress as a task.
func MergePR(ctx *context.ProgramContext, section SectionIdentifier, pr data.RowData) tea.Cmd {
	prNumber := pr.GetNumber()
	c := exec.Command(
		"gh",
		"pr",
		"merge",
		fmt.Sprint(prNumber),
		"-R",
		pr.GetRepoNameWithOwner(),
	)

	taskId := fmt.Sprintf("merge_%d", prNumber)
	task := context.Task{
		ID:           taskId,
		StartText:    fmt.Sprintf("Merging PR #%d", prNumber),
		FinishedText: fmt.Sprintf("PR #%d has been merged", prNumber),
		State:        context.TaskStart,
		Error:        nil,
	}
	startCmd := ctx.StartTask(task)

	return tea.Batch(startCmd, tea.ExecProcess(c, func(err error) tea.Msg {
		isMerged := false
		if err == nil && c.ProcessState.ExitCode() == 0 {
			isMerged = true
		}

		return constants.TaskFinishedMsg{
			SectionID:   section.Id,
			SectionType: section.Type,
			TaskID:      taskId,
			Err:         err,
			Msg: UpdatePRMsg{
				PrNumber: prNumber,
				IsMerged: &isMerged,
			},
		}
	}))
}

// CreatePR creates a new pull request using the provided context, section details, branch name, and title.
// It executes the "gh pr create" command and manages the task lifecycle, returning a tea.Cmd for updates.
func CreatePR(ctx *context.ProgramContext, section SectionIdentifier, branchName string, title string) tea.Cmd {
	c := exec.Command(
		"gh",
		"pr",
		"create",
		"--title",
		title,
		"-R",
	)

	taskId := fmt.Sprintf("create_pr_%s", title)
	task := context.Task{
		ID:           taskId,
		StartText:    fmt.Sprintf(`Creating PR "%s"`, title),
		FinishedText: fmt.Sprintf(`PR "%s" has been created`, title),
		State:        context.TaskStart,
		Error:        nil,
	}
	startCmd := ctx.StartTask(task)

	return tea.Batch(startCmd, tea.ExecProcess(c, func(err error) tea.Msg {
		isCreated := false
		if err == nil && c.ProcessState.ExitCode() == 0 {
			isCreated = true
		}

		return constants.TaskFinishedMsg{
			SectionID:   section.Id,
			SectionType: section.Type,
			TaskID:      taskId,
			Err:         nil,
			Msg:         UpdateBranchMsg{Name: branchName, IsCreated: &isCreated},
		}
	}))
}

// UpdatePR triggers an operation to update a pull request branch, displaying progress messages and firing a task command.
func UpdatePR(ctx *context.ProgramContext, section SectionIdentifier, pr data.RowData) tea.Cmd {
	prNumber := pr.GetNumber()
	return fireTask(ctx, GenericTask{
		Id: buildTaskId("pr_update", prNumber),
		Args: []string{
			"pr",
			"update-branch",
			fmt.Sprint(prNumber),
			"-R",
			pr.GetRepoNameWithOwner(),
		},
		Section:      section,
		StartText:    fmt.Sprintf("Updating PR #%d", prNumber),
		FinishedText: fmt.Sprintf("PR #%d has been updated", prNumber),
		Msg: func(c *exec.Cmd, err error) tea.Msg {
			return UpdatePRMsg{
				PrNumber: prNumber,
				IsClosed: utils.BoolPtr(true),
			}
		},
	})
}
