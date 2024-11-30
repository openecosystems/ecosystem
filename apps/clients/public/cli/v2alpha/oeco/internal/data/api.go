package data

import (
	"fmt"
	"time"
)

type PullRequestData struct {
	Number int
	Title  string
	Body   string
	Author struct {
		Login string
	}
	UpdatedAt      time.Time
	Url            string
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

type Commits struct {
	Nodes []struct {
		Commit struct {
			Deployments struct {
				Nodes []struct {
				}
			}
			StatusCheckRollup struct {
				Contexts struct {
					Nodes []struct {
					}
				}
			}
		}
	}
}

type Comment struct {
	Author struct {
		Login string
	}
	Body      string
	UpdatedAt time.Time
}

type ReviewComment struct {
	Author struct {
		Login string
	}
	Body      string
	UpdatedAt time.Time
	StartLine int
	Line      int
}

type ReviewComments struct {
	Nodes      []ReviewComment
	TotalCount int
}

type Comments struct {
	Nodes      []Comment
	TotalCount int
}

type Review struct {
	Author struct {
		Login string
	}
	Body      string
	State     string
	UpdatedAt time.Time
}

type Reviews struct {
	Nodes []Review
}

type ReviewThreads struct {
	Nodes []struct {
		Id           string
		IsOutdated   bool
		OriginalLine int
		StartLine    int
		Line         int
		Path         string
		Comments     ReviewComments
	}
}

type PRLabel struct {
	Color string
	Name  string
}

type PRLabels struct {
	Nodes []Label
}

type MergeStateStatus string

type PageInfo struct {
	HasNextPage bool
	StartCursor string
	EndCursor   string
}

func (data PullRequestData) GetTitle() string {
	return data.Title
}

func (data PullRequestData) GetRepoNameWithOwner() string {
	return data.Repository.NameWithOwner
}

func (data PullRequestData) GetNumber() int {
	return data.Number
}

func (data PullRequestData) GetUrl() string {
	return data.Url
}

func (data PullRequestData) GetUpdatedAt() time.Time {
	return data.UpdatedAt
}

func makePullRequestsQuery(query string) string {
	return fmt.Sprintf("is:pr %s sort:updated", query)
}

type PullRequestsResponse struct {
	Prs        []PullRequestData
	TotalCount int
	PageInfo   PageInfo
}

func FetchPullRequests(query string, limit int, pageInfo *PageInfo) (PullRequestsResponse, error) {
	prs := make([]PullRequestData, 0, 1)
	prs = append(prs, PullRequestData{
		Number:           0,
		Title:            "Hello",
		Body:             "WOrld",
		UpdatedAt:        time.Now(),
		Url:              "http://openecosystems.com",
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

func FetchPullRequest(prUrl string) (PullRequestData, error) {

	return PullRequestData{
		Number:           0,
		Title:            "Hello",
		Body:             "WOrld",
		UpdatedAt:        time.Now(),
		Url:              "http://openecosystems.com",
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
