package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"seqd/utils"
)

var env = "development"
var version = "0.0.0"
var help = fmt.Sprintf(Help, version)

func main() {
	log.SetFlags(0)
	log.SetPrefix("seqd: ")
	runtime.GOMAXPROCS(1)

	// Short CLI args parser
	args, err := utils.ParseArgs(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	if args.Help {
		fmt.Println(help)
		os.Exit(0)
	}

	if args.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	// Parse input dates first
	start_date, end_date, format, err := utils.GetDateRangeBounds(args.StartDateTime, args.EndDateTime)
	if err != nil {
		log.Fatalln(err)
	}

	if start_date.Equal(end_date) {
		fmt.Println(start_date.Format(format))
		os.Exit(0)
	}

	if (args.IncreaseByHour || args.IncreaseByMinute || args.IncreaseBySecond) && format == time.DateOnly {
		log.Fatalln("flags '-h', '-m' and '-s' can only be used with 'YYYY-MM-DD hh:mm:ss' format of input date")
	}

	// Print start date
	fmt.Println(start_date.Format(format))

	for !start_date.Equal(end_date) {
		if args.IncreaseByYear {
			start_date = start_date.AddDate(1, 0, 0)
		} else if args.IncreaseByMonth {
			start_date = start_date.AddDate(0, 1, 0)
		} else if args.IncreaseByDay {
			start_date = start_date.AddDate(0, 0, 1)
		} else if args.IncreaseByHour {
			start_date = start_date.Add(time.Hour)
		} else if args.IncreaseByMinute {
			start_date = start_date.Add(time.Minute)
		} else if args.IncreaseBySecond {
			start_date = start_date.Add(time.Second)
		}

		// Failsafe
		if start_date.After(end_date) {
			break
		}

		fmt.Println(start_date.Format(format))
	}
}
