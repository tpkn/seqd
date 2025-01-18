package main

const Help = `seqd (v%v) | tpkn.me

Just like Linux 'seq', but for date.

Usage:
  seqd [-Y | -M | -D | -h | -m | -s] <start_date> <end_date> [-r]

Options:
  <end_date>   Can also be a "eom" (end of month) or "eoy" (end of year)
  -Y           Step by years
  -M           Step by months
  -D           Step by days
  -h           Step by hours
  -m           Step by minutes
  -s           Step by seconds
  -r           Reversed order
  --help       Help
  --version    Version

Examples:
  -- By days
  while read day; do
     echo $day
  done <<< "$(seqd -D "2024-02-02" "2024-03-01")"

  -- By days (reversed order)
  while read day; do
     echo $day
  done <<< "$(seqd -D "2024-02-02" "2024-03-01" -r)"

  -- By hours
  while read date_time; do
     day=$(cut -d ' ' -f 1 <<< "date_time")
     hour=$(cut -d ' ' -f 2 <<< "date_time" | cut -d ':' -f 1)
     echo "$day -> $hour"
  done <<< "$(seqd -h "2024-02-02 12:00:00" "2024-03-01 13:00:00")"
`
