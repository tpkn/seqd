<p align="center">
   <img width="250" src="icon.png" alt="" align="center">
</p>
<p align="center">
   Just like <code>seq</code>, but for date.
</p>

## Usage

```
seqd [ -options ] <start_date> <end_date>
```

## Options

```
<end_date>   Can also be a "eom" (end of month) or "eoy" (end of year)
-Y           Step by years   (reversed order: -Yr)
-M           Step by months  (reversed order: -Mr)
-D           Step by days    (reversed order: -Dr)
-h           Step by hours   (reversed order: -hr)
-m           Step by minutes (reversed order: -mr)
-s           Step by seconds (reversed order: -sr)
--help       Help
--version    Version
```

## Examples

By days

```shell
while read day; do
   echo $day
done <<< "$(seqd -D "2024-02-02" "2024-03-01")"
```

By days (reversed order)

```shell
while read day; do
   echo $day
done <<< "$(seqd -Dr "2024-02-02" "2024-03-01")"
```

By hours

```shell
while read date_time; do
   day=$(date -d "$date_time" '+%F')
   hour=$(date -d "$date_time" '+%H')
   echo "$day -> $hour"
done <<< "$(seqd -h "2024-02-02 12:00:00" "2024-03-01 13:00:00")"
```
