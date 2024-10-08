package main

const Help = `seqd (v%v) | tpkn.me

Just like Linux 'seq', but for date.

Usage:
  seqd [-Y | -M | -D | -h | -m | -s] <start_date> <end_date>

Options:
  <end_date>   Can also be a "eom" (end of month) or "eoy" (end of year)
  -Y           Step by years
  -M           Step by months
  -D           Step by days
  -h           Step by hours
  -m           Step by minutes
  -s           Step by seconds
  --help       Help
  --version    Version

Examples:
  -- Date with time
  while read day_time; do
     echo $day_time
  done <<< "$(seqd -h "2024-02-02 12:00:00" "2024-03-01 13:00:00")"

  -- Just date
  for day in $(seqd -h "2024-02-02" "2024-03-01"); do
     echo $day
  done
`
