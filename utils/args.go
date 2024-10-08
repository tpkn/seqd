package utils

import (
	"errors"

	"seqd/models"
)

// ParseArgs returns parsed Args arguments
func ParseArgs(a []string) (models.Args, error) {
	var args = models.Args{}

	switch len(a) - 1 {
	case 1:
		if a[1] == "--help" {
			args.Help = true
		} else if a[1] == "--version" {
			args.Version = true
		} else {
			return args, errors.New("wrong argument")
		}
	case 3:
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
	default:
		return args, errors.New("wrong amount of arguments")
	}

	return args, nil
}
