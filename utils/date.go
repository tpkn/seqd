package utils

import (
	"errors"
	"fmt"
	"time"
)

// GetDateRangeBounds returns date range start and end points
func GetDateRangeBounds(start_date, end_date string) (time.Time, time.Time, string, error) {
	d1, start_format, err := parseDate(start_date)
	if err != nil {
		return time.Time{}, time.Time{}, "", err
	}
	
	d2, end_format, err := parseDate(end_date)
	if err != nil {
		return time.Time{}, time.Time{}, "", err
	}
	
	if start_format != end_format {
		return time.Time{}, time.Time{}, "", errors.New(fmt.Sprintf("start date and end date has different format: '%v' != '%v'", start_format, end_format))
	}
	
	if d1.After(d2) {
		return time.Time{}, time.Time{}, "", errors.New(fmt.Sprintf("start date is greater than end date: '%v' > '%v'", d1.Format(start_format), d2.Format(start_format)))
	}
	
	return d1, d2, start_format, nil
}

// parseDate returns parsed date string and it's format
func parseDate(d string) (time.Time, string, error) {
	format := time.DateTime
	date, err := time.Parse(format, d)
	if err != nil {
		// Now just as date
		format = time.DateOnly
		if date, err = time.Parse(format, d); err != nil {
			return time.Time{}, "", errors.New(fmt.Sprintf("can't parse date '%v' beacause it has a wrong format: should be 'YYYY-MM-DD' or 'YYYY-MM-DD hh:mm:ss')", d))
		}
	}
	return date, format, nil
}
