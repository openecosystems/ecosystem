package data

import (
	"fmt"
	"time"
)

// PullRequestData represents detailed information about a pull request, including metadata, state, and related entities.
type PullRequestData struct {
	Number int
	Title  string
	Body   string
	Author struct {
		Login string
	}
	UpdatedAt      time.Time
	URL            string
	State          string
	Mergeable      string
	ReviewDecision string
	Additions      int
	Deletions      int
	HeadRefName    string
	BaseRefName    string
	HeadRepository struct {
		Name string
	}
	HeadRef struct {
		Name string
	}
	Repository       Repository
	Assignees        Assignees     `graphql:"assignees(first: 3)"`
	Comments         Comments      `graphql:"comments(last: 5, orderBy: { field: UPDATED_AT, direction: DESC })"`
	LatestReviews    Reviews       `graphql:"latestReviews(last: 3)"`
	ReviewThreads    ReviewThreads `graphql:"reviewThreads(last: 20)"`
	IsDraft          bool
	Commits          Commits          `graphql:"commits(last: 1)"`
	Labels           PRLabels         `graphql:"labels(first: 3)"`
	MergeStateStatus MergeStateStatus `graphql:"mergeStateStatus"`
}

// Commits represents a collection of git commits with nested data such as deployments and status check contexts.
type Commits struct {
	Nodes []struct {
		Commit struct {
			Deployments struct {
				Nodes []struct{}
			}
			StatusCheckRollup struct {
				Contexts struct {
					Nodes []struct{}
				}
			}
		}
	}
}

// Comment represents a user's comment with associated author details, content, and the last updated timestamp.
type Comment struct {
	Author struct {
		Login string
	}
	Body      string
	UpdatedAt time.Time
}

// ReviewComment represents a comment made in a code review, including details like author, content, and line information.
type ReviewComment struct {
	Author struct {
		Login string
	}
	Body      string
	UpdatedAt time.Time
	StartLine int
	Line      int
}

// ReviewComments represents a collection of review comments and the total count of comments in a review thread.
type ReviewComments struct {
	Nodes      []ReviewComment
	TotalCount int
}

// Comments represents a collection of comments with a list of nodes and their total count.
type Comments struct {
	Nodes      []Comment
	TotalCount int
}

// Review represents a review submitted by an author with associated details like body content, state, and update timestamp.
type Review struct {
	Author struct {
		Login string
	}
	Body      string
	State     string
	UpdatedAt time.Time
}

// Reviews represents a collection of Review nodes associated with a pull request or similar entity.
type Reviews struct {
	Nodes []Review
}

// ReviewThreads represents a collection of review threads, each containing metadata and associated comments.
type ReviewThreads struct {
	Nodes []struct {
		ID           string
		IsOutdated   bool
		OriginalLine int
		StartLine    int
		Line         int
		Path         string
		Comments     ReviewComments
	}
}

// PRLabel represents a pull request label with a specified color and name.
type PRLabel struct {
	Color string
	Name  string
}

// PRLabels represents a collection of labels associated with a pull request.
type PRLabels struct {
	Nodes []Label
}

// MergeStateStatus represents the state of a pull request in relation to its merge readiness or conflicts.
type MergeStateStatus string

// PageInfo provides pagination details, indicating if more items exist and the cursors for the current page.
type PageInfo struct {
	HasNextPage bool
	StartCursor string
	EndCursor   string
}

// GetTitle returns the title of the pull request as a string.
func (data PullRequestData) GetTitle() string {
	return data.Title
}

// GetRepoNameWithOwner returns the full repository name including the owner in the format "owner/repo".
func (data PullRequestData) GetRepoNameWithOwner() string {
	return data.Repository.NameWithOwner
}

// GetNumber returns the pull request number associated with the PullRequestData struct.
func (data PullRequestData) GetNumber() int {
	return data.Number
}

// GetURL returns the URL of the pull request stored in the PullRequestData instance.
func (data PullRequestData) GetURL() string {
	return data.URL
}

// GetUpdatedAt returns the timestamp of the last update for the pull request.
func (data PullRequestData) GetUpdatedAt() time.Time {
	return data.UpdatedAt
}

// makePullRequestsQuery generates a search query string for pull requests by appending sorting and filtering criteria.
//
//nolint:unused
func makePullRequestsQuery(query string) string {
	return fmt.Sprintf("is:pr %s sort:updated", query)
}

// PullRequestsResponse represents a response containing a list of pull requests, their total count, and pagination info.
type PullRequestsResponse struct {
	Prs        []PullRequestData
	TotalCount int
	PageInfo   PageInfo
}

// FetchPullRequests retrieves a list of pull requests with pagination details for a repository.
// It accepts a repository identifier, a limit for pagination, and a PageInfo pointer.
// Returns a PullRequestsResponse containing pull request data, or an error if the operation fails.
func FetchPullRequests(_ string, _ int, _ *PageInfo) (PullRequestsResponse, error) {
	prs := make([]PullRequestData, 0, 1)
	prs = append(prs, PullRequestData{
		Number:           0,
		Title:            "Hello",
		Body:             "WOrld",
		UpdatedAt:        time.Now(),
		URL:              "http://openecosystems.com",
		State:            "",
		Mergeable:        "",
		ReviewDecision:   "",
		Additions:        0,
		Deletions:        0,
		HeadRefName:      "",
		BaseRefName:      "",
		HeadRepository:   struct{ Name string }{},
		HeadRef:          struct{ Name string }{},
		Repository:       Repository{},
		Assignees:        Assignees{},
		Comments:         Comments{},
		LatestReviews:    Reviews{},
		ReviewThreads:    ReviewThreads{},
		IsDraft:          false,
		Commits:          Commits{},
		Labels:           PRLabels{},
		MergeStateStatus: "",
	})

	return PullRequestsResponse{
		Prs:        prs,
		TotalCount: 1,
		PageInfo: PageInfo{
			HasNextPage: false,
			StartCursor: "0",
			EndCursor:   "0",
		},
	}, nil
}

// FetchPullRequest retrieves data for a pull request from the provided URL and returns it as a PullRequestData object.
func FetchPullRequest(_ string) (PullRequestData, error) {
	return PullRequestData{
		Number:           0,
		Title:            "Hello",
		Body:             "WOrld",
		UpdatedAt:        time.Now(),
		URL:              "http://openecosystems.com",
		State:            "",
		Mergeable:        "",
		ReviewDecision:   "",
		Additions:        0,
		Deletions:        0,
		HeadRefName:      "",
		BaseRefName:      "",
		HeadRepository:   struct{ Name string }{},
		HeadRef:          struct{ Name string }{},
		Repository:       Repository{},
		Assignees:        Assignees{},
		Comments:         Comments{},
		LatestReviews:    Reviews{},
		ReviewThreads:    ReviewThreads{},
		IsDraft:          false,
		Commits:          Commits{},
		Labels:           PRLabels{},
		MergeStateStatus: "",
	}, nil
}
