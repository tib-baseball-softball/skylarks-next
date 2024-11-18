package utility

import "time"

// YearStartAndEnd returns the start and end of the given year as formatted strings
func YearStartAndEnd(year int) (string, string) {
	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	startOfYearStr := startOfYear.Format("2006-01-02 15:04:05.000Z")

	endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999000000, time.UTC)
	endOfYearStr := endOfYear.Format("2006-01-02 15:04:05.000Z")

	return startOfYearStr, endOfYearStr
}
