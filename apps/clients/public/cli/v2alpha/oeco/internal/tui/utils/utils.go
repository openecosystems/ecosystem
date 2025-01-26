package utils

import (
	"math"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// ApproxDaysInYear represents the approximate number of days in a year.
// ApproxDaysInMonth represents the approximate number of days in a month.
// DaysInWeek represents the number of days in a week.
const (
	ApproxDaysInYear  = 365
	ApproxDaysInMonth = 28
	DaysInWeek        = 7
)

// Max returns the greater of two integer values a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of the two integer values a and b.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TimeElapsed returns a human-readable string representing the elapsed time from the given time to the current time.
func TimeElapsed(then time.Time) string {
	var parts []string
	var text string

	now := time.Now()
	diff := now.Sub(then)
	day := math.Floor(diff.Hours() / 24)
	year := math.Floor(day / ApproxDaysInYear)
	month := math.Floor(day / ApproxDaysInMonth)
	week := math.Floor(day / DaysInWeek)
	hour := math.Floor(math.Abs(diff.Hours()))
	minute := math.Floor(math.Abs(diff.Minutes()))
	second := math.Floor(math.Abs(diff.Seconds()))

	if year > 0 {
		parts = append(parts, strconv.Itoa(int(year))+"y")
	}

	if month > 0 {
		parts = append(parts, strconv.Itoa(int(month))+"mo")
	}

	if week > 0 {
		parts = append(parts, strconv.Itoa(int(week))+"w")
	}

	if day > 0 {
		parts = append(parts, strconv.Itoa(int(day))+"d")
	}

	if hour > 0 {
		parts = append(parts, strconv.Itoa(int(hour))+"h")
	}

	if minute > 0 {
		parts = append(parts, strconv.Itoa(int(minute))+"m")
	}

	if second > 0 {
		parts = append(parts, strconv.Itoa(int(second))+"s")
	}

	if len(parts) == 0 {
		return "now"
	}

	return parts[0] + text
}

// BoolPtr returns a pointer to the given boolean value.
func BoolPtr(b bool) *bool { return &b }

// StringPtr takes a string and returns a pointer to that string.
func StringPtr(s string) *string { return &s }

// UintPtr takes a uint value and returns a pointer to that value.
func UintPtr(u uint) *uint { return &u }

// IntPtr takes an integer value and returns a pointer to that integer.
func IntPtr(u int) *int { return &u }

// ShortNumber formats an integer as a shorter string representation, using "k" for thousands and "m" for millions.
func ShortNumber(n int) string {
	if n < 1000 {
		return strconv.Itoa(n)
	}

	if n < 1000000 {
		return strconv.Itoa(n/1000) + "k"
	}

	return strconv.Itoa(n/1000000) + "m"
}

// ShowError is a function that logs a fatal error message with detailed context such as timestamp and caller information.
var ShowError = func(err error) {
	styles := log.DefaultStyles()
	styles.Key = lipgloss.NewStyle().
		Foreground(lipgloss.Color("1")).
		Bold(true)
	styles.Separator = lipgloss.NewStyle()

	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	logger.SetTimeFormat(time.RFC3339)
	logger.SetReportTimestamp(true)
	logger.SetReportCaller(true)

	logger.
		Fatal(
			"failed:",
			"err",
			err,
		)
}
