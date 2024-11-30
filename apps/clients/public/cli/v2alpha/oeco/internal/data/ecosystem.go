package data

import (
	"fmt"
	"time"
)

type EcosystemData struct {
	Number int
	Title  string
	Body   string
	State  string
	Author struct {
		Login string
	}
	UpdatedAt  time.Time
	Url        string
	Repository Repository
	Assignees  Assignees      `graphql:"assignees(first: 3)"`
	Comments   IssueComments  `graphql:"comments(first: 15)"`
	Reactions  IssueReactions `graphql:"reactions(first: 1)"`
	Labels     IssueLabels    `graphql:"labels(first: 3)"`
}

type IssueComments struct {
	Nodes      []IssueComment
	TotalCount int
}

type IssueComment struct {
	Author struct {
		Login string
	}
	Body      string
	UpdatedAt time.Time
}

type IssueReactions struct {
	TotalCount int
}

type Label struct {
	Color string
	Name  string
}

type IssueLabels struct {
	Nodes []Label
}

func (data EcosystemData) GetTitle() string {
	return data.Title
}

func (data EcosystemData) GetRepoNameWithOwner() string {
	return data.Repository.NameWithOwner
}

func (data EcosystemData) GetNumber() int {
	return data.Number
}

func (data EcosystemData) GetUrl() string {
	return data.Url
}

func (data EcosystemData) GetUpdatedAt() time.Time {
	return data.UpdatedAt
}

func makeIssuesQuery(query string) string {
	return fmt.Sprintf("is:issue %s sort:updated", query)
}

func FetchEcosystems(query string, limit int, pageInfo *PageInfo) (IssuesResponse, error) {
	issues := make([]EcosystemData, 0, 1)
	issues = append(issues, EcosystemData{
		Number: 0,
		Title:  "Hello World",
		Body:   "Test Bod",
		State:  "Active",
		Author: struct{ Login string }{
			Login: "12345",
		},
		UpdatedAt: time.Now(),
		Url:       "https://openecosystems.com",
		Repository: Repository{
			Name:          "test",
			NameWithOwner: "world",
			IsArchived:    false,
		},
		Assignees: Assignees{
			Nodes: nil,
		},
		Comments: IssueComments{
			Nodes:      nil,
			TotalCount: 0,
		},
		Reactions: IssueReactions{
			TotalCount: 0,
		},
		Labels: IssueLabels{
			Nodes: nil,
		},
	})

	return IssuesResponse{
		Issues:     issues,
		TotalCount: 1,
		PageInfo: PageInfo{
			HasNextPage: false,
			StartCursor: "0",
			EndCursor:   "0",
		},
	}, nil
}

type IssuesResponse struct {
	Issues     []EcosystemData
	TotalCount int
	PageInfo   PageInfo
}
