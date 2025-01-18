package core

import (
	"errors"
)

type Args struct {
	StartDateTime    string
	EndDateTime      string
	IncreaseByYear   bool
	IncreaseByMonth  bool
	IncreaseByDay    bool
	IncreaseByHour   bool
	IncreaseByMinute bool
	IncreaseBySecond bool
	ReversedOrder    bool
	Help             bool
	Version          bool
}

// ParseArgs returns parsed Args arguments
func ParseArgs(a []string) (Args, error) {
	var args = Args{}
	
	switch len(a) - 1 {
	case 1:
		if a[1] == "--help" {
			args.Help = true
		} else if a[1] == "--version" {
			args.Version = true
		} else {
			return args, errors.New("wrong argument")
		}
	case 3, 4:
		args.StartDateTime = a[1]
		switch a[1] {
		case "-Y":
			args.IncreaseByYear = true
		case "-M":
			args.IncreaseByMonth = true
		case "-D":
			args.IncreaseByDay = true
		case "-h":
			args.IncreaseByHour = true
		case "-m":
			args.IncreaseByMinute = true
		case "-s":
			args.IncreaseBySecond = true
		default:
			return args, errors.New("wrong arguments")
		}
		
		args.StartDateTime = a[2]
		args.EndDateTime = a[3]
		
		if len(a) == 5 {
			switch a[4] {
			case "-r":
				args.ReversedOrder = true
			default:
				return args, errors.New("wrong argument")
			}
		}
	default:
		return args, errors.New("wrong amount of arguments")
	}
	
	return args, nil
}
