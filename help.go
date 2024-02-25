package main

const Help = `seqd (v%v) | tpkn.me
Just like Linux 'seq', but for date.

Usage:
  seqd [-Y | -M | -D | -h | -m | -s] "<start_date>" "<end_date>"

Example of usage in bash:
  while read date_time; do
     echo $date_time
  done <<< "$(seqd -h "2024-02-02 12:00:00" "2024-02-02 13:00:00")"
  
  # or
  
  IFS=$'\n'
  for date_time in $(seqd -h "2024-02-02 12:00:00" "2024-03-01 23:00:00"); do
     echo $date_time
  done

Options:
  -Y           Step by years
  -M           Step by months
  -D           Step by days
  -h           Step by hours
  -m           Step by minutes
  -s           Step by seconds
  --help       Help
  --version    Version
`
