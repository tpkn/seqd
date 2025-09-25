package main

const Help = `                                     /▒▒
                                    | ▒▒
  /▒▒▒▒▒▒▒  /▒▒▒▒▒▒   /▒▒▒▒▒▒   /▒▒▒▒▒▒▒
 /▒▒_____/ /▒▒__  ▒▒ /▒▒__  ▒▒ /▒▒__  ▒▒
|  ▒▒▒▒▒▒ | ▒▒▒▒▒▒▒▒| ▒▒  \ ▒▒| ▒▒  | ▒▒
 \____  ▒▒| ▒▒_____/| ▒▒  | ▒▒| ▒▒  | ▒▒
 /▒▒▒▒▒▒▒/|  ▒▒▒▒▒▒▒|  ▒▒▒▒▒▒▒|  ▒▒▒▒▒▒▒ v%v
|_______/  \_______/ \____  ▒▒ \_______/ https://tpkn.me
                          | ▒▒
                          | ▒▒
                          |__/

Just like 'seq', but for date.

Usage:
  seqd [-Y | -M | -D | -h | -m | -s] <start_date> <end_date>

Options:
  <end_date>   Can also be a "eom" (end of month) or "eoy" (end of year)
  -Y           Step by years   (-Yr for reversed order)
  -M           Step by months  (-Mr for reversed order)
  -D           Step by days    (-Dr for reversed order)
  -h           Step by hours   (-hr for reversed order)
  -m           Step by minutes (-mr for reversed order)
  -s           Step by seconds (-sr for reversed order)
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
  done <<< "$(seqd -Dr "2024-02-02" "2024-03-01")"

  -- By hours
  while read date_time; do
     day=$(date -d "$date_time" '+%%F')
     hour=$(date -d "$date_time" '+%%H')
     echo "$day -> $hour"
  done <<< "$(seqd -h "2024-02-02 12:00:00" "2024-03-01 13:00:00")"
`
