package core

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
	Format    string
}

// PrintDateRange walk over date range and print each date
func PrintDateRange(output *os.File, args *Args) error {
	// Parse input dates first
	date_range, err := GetDateRangeBounds(args.StartDateTime, args.EndDateTime)
	if err != nil {
		return err
	}
	
	// If start date and end dates are equal, nothing needs to be done
	if date_range.StartDate.Equal(date_range.EndDate) {
		_, err = output.WriteString(date_range.StartDate.Format(date_range.Format) + "\n")
		if err != nil {
			return err
		}
		return nil
	}
	
	if (args.IncreaseByHour || args.IncreaseByMinute || args.IncreaseBySecond) && date_range.Format == time.DateOnly {
		return errors.New("flags '-h', '-m' and '-s' can only be used with 'YYYY-MM-DD hh:mm:ss' format of input date")
	}
	
	if !args.ReversedOrder {
		_, err = output.WriteString(date_range.StartDate.Format(date_range.Format) + "\n")
		if err != nil {
			return err
		}
		
		for !date_range.StartDate.Equal(date_range.EndDate) {
			if args.IncreaseByYear {
				date_range.StartDate = date_range.StartDate.AddDate(1, 0, 0)
			} else if args.IncreaseByMonth {
				date_range.StartDate = date_range.StartDate.AddDate(0, 1, 0)
			} else if args.IncreaseByDay {
				date_range.StartDate = date_range.StartDate.AddDate(0, 0, 1)
			} else if args.IncreaseByHour {
				date_range.StartDate = date_range.StartDate.Add(time.Hour)
			} else if args.IncreaseByMinute {
				date_range.StartDate = date_range.StartDate.Add(time.Minute)
			} else if args.IncreaseBySecond {
				date_range.StartDate = date_range.StartDate.Add(time.Second)
			}
			if date_range.StartDate.After(date_range.EndDate) {
				break
			}
			
			_, err = output.WriteString(date_range.StartDate.Format(date_range.Format) + "\n")
			if err != nil {
				return err
			}
		}
	} else {
		_, err = output.WriteString(date_range.EndDate.Format(date_range.Format) + "\n")
		if err != nil {
			return err
		}
		
		for !date_range.EndDate.Equal(date_range.StartDate) {
			if args.IncreaseByYear {
				date_range.EndDate = date_range.EndDate.AddDate(-1, 0, 0)
			} else if args.IncreaseByMonth {
				date_range.EndDate = date_range.EndDate.AddDate(0, -1, 0)
			} else if args.IncreaseByDay {
				date_range.EndDate = date_range.EndDate.AddDate(0, 0, -1)
			} else if args.IncreaseByHour {
				date_range.EndDate = date_range.EndDate.Add(-time.Hour)
			} else if args.IncreaseByMinute {
				date_range.EndDate = date_range.EndDate.Add(-time.Minute)
			} else if args.IncreaseBySecond {
				date_range.EndDate = date_range.EndDate.Add(-time.Second)
			}
			if date_range.EndDate.Before(date_range.StartDate) {
				break
			}
			
			_, err = output.WriteString(date_range.EndDate.Format(date_range.Format) + "\n")
			if err != nil {
				return err
			}
		}
	}
	
	return nil
}

// GetDateRangeBounds returns date range start and end points
func GetDateRangeBounds(start_date, end_date string) (DateRange, error) {
	result := DateRange{}
	
	d1, start_format, err := parseDate(start_date)
	if err != nil {
		return result, err
	}
	
	// End of month date
	if end_date == "eom" {
		end_date = endOfMonth(d1).Format(start_format)
	}
	
	// End of year date
	if end_date == "eoy" {
		end_date = endOfYear(d1).Format(start_format)
	}
	
	d2, end_format, err := parseDate(end_date)
	if err != nil {
		return result, err
	}
	
	if start_format != end_format {
		return result, errors.New(fmt.Sprintf("start date and end date has different format (%v != %v)", start_format, end_format))
	}
	
	if d1.After(d2) {
		return result, errors.New(fmt.Sprintf("start date is greater than end date (%v > %v)", d1.Format(start_format), d2.Format(start_format)))
	}
	
	result.StartDate = d1
	result.EndDate = d2
	result.Format = start_format
	
	return result, nil
}

// parseDate returns parsed date string and it's format
func parseDate(d string) (time.Time, string, error) {
	format := time.DateTime
	date, err := time.Parse(format, d)
	if err != nil {
		// Now just as date
		format = time.DateOnly
		if date, err = time.Parse(format, d); err != nil {
			return time.Time{}, "", errors.New(fmt.Sprintf("can't parse date '%v' beacause it has a wrong format (should be 'YYYY-MM-DD' or 'YYYY-MM-DD hh:mm:ss')", d))
		}
	}
	return date, format, nil
}

// endOfMonth returns last day of the month ('2023-03-06' => '2023-03-31')
func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

// endOfYear returns last day of the year ('2023-03-06' => '2023-12-31')
func endOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), 12, 31, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
}
