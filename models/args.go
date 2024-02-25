package models

type Args struct {
	StartDateTime    string
	EndDateTime      string
	IncreaseByYear   bool
	IncreaseByMonth  bool
	IncreaseByDay    bool
	IncreaseByHour   bool
	IncreaseByMinute bool
	IncreaseBySecond bool
	Help             bool
	Version          bool
}
