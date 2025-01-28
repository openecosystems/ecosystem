package data

import (
	"time"
)

// RowData represents an interface for accessing fields related to a row of data in a structured manner.
type RowData interface {
	GetRepoNameWithOwner() string
	GetTitle() string
	GetNumber() int
	GetUrl() string
	GetUpdatedAt() time.Time
}

// IsStatusWaiting checks if the provided status indicates a waiting state, such as PENDING, QUEUED, IN_PROGRESS, or WAITING.
func IsStatusWaiting(status string) bool {
	return status == "PENDING" ||
		status == "QUEUED" ||
		status == "IN_PROGRESS" ||
		status == "WAITING"
}

// IsConclusionAFailure checks if the given conclusion indicates a failure state.
func IsConclusionAFailure(conclusion string) bool {
	return conclusion == "FAILURE" || conclusion == "TIMED_OUT" || conclusion == "STARTUP_FAILURE"
}
