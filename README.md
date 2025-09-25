<p align="center">
   <img width="150" src="icon.png" alt="" align="center">
</p>
<p align="center">
   Just like <code>seq</code>, but for date.
</p>

## Usage

```
seqd [-Y | -M | -D | -h | -m | -s] <start_date> <end_date>
```

## Options

```
<end_date>   Can also be a "eom" (end of month) or "eoy" (end of year)
-Y           Step by years   (-Yr for reversed order)
-M           Step by months  (-Mr for reversed order)
-D           Step by days    (-Dr for reversed order)
-h           Step by hours   (-hr for reversed order)
-m           Step by minutes (-mr for reversed order)
-s           Step by seconds (-sr for reversed order)
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
