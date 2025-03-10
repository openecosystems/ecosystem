package ecosystemdashboardcontent

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) details() tea.Cmd {
	return tea.Batch()

	//issue := m.GetCurrRow()
	//issueNumber := issue.GetNumber()
	//taskId := fmt.Sprintf("issue_close_%d", issueNumber)
	//task := tasks.Task{
	//	ID:           taskId,
	//	StartText:    fmt.Sprintf("Closing issue #%d", issueNumber),
	//	FinishedText: fmt.Sprintf("Issue #%d has been closed", issueNumber),
	//	State:        tasks.TaskStart,
	//	Error:        nil,
	//}
	//startCmd := m.Ctx.StartTask(task)
	//return tea.Batch(startCmd, func() tea.Msg {
	//	c := exec.Command(
	//		"gh",
	//		"issue",
	//		"close",
	//		fmt.Sprint(m.GetCurrRow().GetNumber()),
	//		"-R",
	//		m.GetCurrRow().GetRepoNameWithOwner(),
	//	)
	//
	//	err := c.Run()
	//	return tasks.TaskFinishedMsg{
	//		SectionID:   m.ID,
	//		SectionType: PageType,
	//		TaskID:      taskId,
	//		Err:         err,
	//		Msg: UpdateIssueMsg{
	//			IssueNumber: issueNumber,
	//			IsClosed:    utils.BoolPtr(true),
	//		},
	//	}
	//})
}
