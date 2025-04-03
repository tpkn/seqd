package seqd

import (
	"errors"
	"time"
)

type OutputWriter interface {
	WriteString(string) (n int, err error)
}

// GenerateDateRange generates date range an send each date to specified ouptput
func GenerateDateRange(output OutputWriter, args *Args) error {
	// Parse input dates first
	dateRange, err := getDateRangeBounds(args.StartDateTime, args.EndDateTime)
	if err != nil {
		return err
	}

	if (args.IncreaseByHour || args.IncreaseByMinute || args.IncreaseBySecond) && dateRange.Format == time.DateOnly {
		return errors.New("flags '-h', '-m' and '-s' can only be used with 'YYYY-MM-DD hh:mm:ss' format of input date")
	}

	// If start date and end dates are equal, nothing needs to be done
	if dateRange.Start.Equal(dateRange.End) {
		_, err = output.WriteString(dateRange.Start.Format(dateRange.Format))
		if err != nil {
			return err
		}
		return nil
	}

	if !args.ReversedOrder {
		_, err = output.WriteString(dateRange.Start.Format(dateRange.Format))
		if err != nil {
			return err
		}

		for !dateRange.Start.Equal(dateRange.End) {
			if args.IncreaseByYear {
				dateRange.Start = dateRange.Start.AddDate(1, 0, 0)
			} else if args.IncreaseByMonth {
				dateRange.Start = dateRange.Start.AddDate(0, 1, 0)
			} else if args.IncreaseByDay {
				dateRange.Start = dateRange.Start.AddDate(0, 0, 1)
			} else if args.IncreaseByHour {
				dateRange.Start = dateRange.Start.Add(time.Hour)
			} else if args.IncreaseByMinute {
				dateRange.Start = dateRange.Start.Add(time.Minute)
			} else if args.IncreaseBySecond {
				dateRange.Start = dateRange.Start.Add(time.Second)
			}
			if dateRange.Start.After(dateRange.End) {
				break
			}

			_, err = output.WriteString(dateRange.Start.Format(dateRange.Format))
			if err != nil {
				return err
			}
		}
	} else {
		_, err = output.WriteString(dateRange.End.Format(dateRange.Format))
		if err != nil {
			return err
		}

		for !dateRange.End.Equal(dateRange.Start) {
			if args.IncreaseByYear {
				dateRange.End = dateRange.End.AddDate(-1, 0, 0)
			} else if args.IncreaseByMonth {
				dateRange.End = dateRange.End.AddDate(0, -1, 0)
			} else if args.IncreaseByDay {
				dateRange.End = dateRange.End.AddDate(0, 0, -1)
			} else if args.IncreaseByHour {
				dateRange.End = dateRange.End.Add(-time.Hour)
			} else if args.IncreaseByMinute {
				dateRange.End = dateRange.End.Add(-time.Minute)
			} else if args.IncreaseBySecond {
				dateRange.End = dateRange.End.Add(-time.Second)
			}
			if dateRange.End.Before(dateRange.Start) {
				break
			}

			_, err = output.WriteString(dateRange.End.Format(dateRange.Format))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
