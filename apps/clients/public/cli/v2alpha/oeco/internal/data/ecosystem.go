package data

import (
	"fmt"
	"time"
)

// EcosystemData represents the data structure containing information about a specific issue or item in the ecosystem.
type EcosystemData struct {
	Number int
	Title  string
	Body   string
	State  string
	Author struct {
		Login string
	}
	UpdatedAt  time.Time
	URL        string
	Repository Repository
	Assignees  Assignees      `graphql:"assignees(first: 3)"`
	Comments   IssueComments  `graphql:"comments(first: 15)"`
	Reactions  IssueReactions `graphql:"reactions(first: 1)"`
	Labels     IssueLabels    `graphql:"labels(first: 3)"`
}

// IssueComments represents a collection of issue comments and their total count.
type IssueComments struct {
	Nodes      []IssueComment
	TotalCount int
}

// IssueComment represents a comment on an issue with its author, body content, and last updated timestamp.
type IssueComment struct {
	Author struct {
		Login string
	}
	Body      string
	UpdatedAt time.Time
}

// IssueReactions represents the reactions associated with an issue, including the total count of reactions.
type IssueReactions struct {
	TotalCount int
}

// Label represents a label with a color and name, typically used for categorization or identification.
type Label struct {
	Color string
	Name  string
}

// IssueLabels represents a collection of labels associated with an issue.
// Nodes is a slice of Label type that contains individual label details.
type IssueLabels struct {
	Nodes []Label
}

// GetTitle returns the title of the EcosystemData.
func (data EcosystemData) GetTitle() string {
	return data.Title
}

// GetRepoNameWithOwner returns the repository name along with the owner's name from the EcosystemData structure.
func (data EcosystemData) GetRepoNameWithOwner() string {
	return data.Repository.NameWithOwner
}

// GetNumber returns the number associated with the EcosystemData instance.
func (data EcosystemData) GetNumber() int {
	return data.Number
}

// GetURL returns the Url field from the EcosystemData structure.
func (data EcosystemData) GetURL() string {
	return data.URL
}

// GetUpdatedAt returns the time when the EcosystemData was last updated.
func (data EcosystemData) GetUpdatedAt() time.Time {
	return data.UpdatedAt
}

// makeIssuesQuery formats a GitHub issues search query string by appending "is:issue" and "sort:updated" to the input query.
//
//nolint:unused
func makeIssuesQuery(query string) string {
	return fmt.Sprintf("is:issue %s sort:updated", query)
}

// FetchEcosystems retrieves a list of ecosystem issues based on the provided query, limit, and pagination information.
func FetchEcosystems(_ string, _ int, _ *PageInfo) (IssuesResponse, error) {
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
		URL:       "https://openecosystems.com",
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

// IssuesResponse represents a response containing a list of issues, the total count of issues, and pagination information.
type IssuesResponse struct {
	Issues     []EcosystemData
	TotalCount int
	PageInfo   PageInfo
}
