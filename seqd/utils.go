package seqd

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateRange struct {
	Start  time.Time
	End    time.Time
	Format string
}

// parseDate returns parsed date and it's format
func parseDate(d string) (time.Time, string, error) {
	format := time.DateTime
	date, err := time.Parse(format, d)
	if err != nil {
		// Now just as date
		format = time.DateOnly
		if date, err = time.Parse(format, d); err != nil {
			return time.Time{}, "", fmt.Errorf("can't parse date '%v' beacause it has a wrong format (should be 'YYYY-MM-DD' or 'YYYY-MM-DD hh:mm:ss')", d)
		}
	}
	return date, format, nil
}

// getDateRangeBounds returns date range start and end points
func getDateRangeBounds(start_date, end_date string) (DateRange, error) {
	result := DateRange{}

	if strings.Contains(start_date, "eom") {
		return result, errors.New("'eom' macros can only be used in end date")
	}
	if strings.Contains(start_date, "eoy") {
		return result, errors.New("'eoy' macros can only be used in end date")
	}

	d1, start_format, err := parseDate(start_date)
	if err != nil {
		return result, err
	}

	// End of month date
	if end_date == "eom" {
		end_date = getEndOfMonthDate(d1).Format(start_format)
	}

	// End of year date
	if end_date == "eoy" {
		end_date = getAndOfYearDate(d1).Format(start_format)
	}

	d2, end_format, err := parseDate(end_date)
	if err != nil {
		return result, err
	}

	if start_format != end_format {
		return result, fmt.Errorf("start date and end date has different format (%v != %v)", start_format, end_format)
	}

	if d1.After(d2) {
		return result, fmt.Errorf("start date is greater than end date (%v > %v)", d1.Format(start_format), d2.Format(start_format))
	}

	result.Start = d1
	result.End = d2
	result.Format = start_format

	return result, nil
}

// getEndOfMonthDate returns last day of the month ('2023-03-06' => '2023-03-31')
func getEndOfMonthDate(date time.Time) time.Time {
	date = date.AddDate(0, 1, -date.Day())
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, date.Nanosecond(), date.Location())
}

// getAndOfYearDate returns last day of the year ('2023-03-06' => '2023-12-31')
func getAndOfYearDate(date time.Time) time.Time {
	return time.Date(date.Year(), 12, 31, 23, 59, 59, date.Nanosecond(), date.Location())
}
