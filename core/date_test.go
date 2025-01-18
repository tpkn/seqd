package core

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseDate(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		want_format string
		wantErr     bool
	}{
		{name: "Wrong date", args: args{d: "2023"}, want: "", want_format: "", wantErr: true},
		{name: "Just date", args: args{d: "20/12 12:30:00PM"}, want: "", want_format: "", wantErr: true},
		{name: "Just date", args: args{d: "2023-12-01"}, want: "2023-12-01", want_format: time.DateOnly, wantErr: false},
		{name: "Date and time", args: args{d: "2023-12-01 22:04:44"}, want: "2023-12-01 22:04:44", want_format: time.DateTime, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got_format, err := parseDate(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Format(got_format), tt.want) && !tt.wantErr {
				t.Errorf("parseDate() got = %v, want %v", got.Format(got_format), tt.want)
			}
			if got_format != tt.want_format && !tt.wantErr {
				t.Errorf("parseDate() got_format = %v, want %v", got_format, tt.want_format)
			}
		})
	}
}

func Test_GetDateRangeBounds(t *testing.T) {
	type args struct {
		start_date string
		end_date   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		{name: "No end date", args: args{start_date: "2024-01-01", end_date: ""}, want: "", want1: "", want2: "", wantErr: true},
		{name: "Wrong date format", args: args{start_date: "2024-01-01", end_date: "2024"}, want: "", want1: "", want2: "", wantErr: true},
		{name: "Start date < end date", args: args{start_date: "2024-01-01", end_date: "2023-02-02"}, want: "", want1: "", want2: "", wantErr: true},
		{name: "Formats are not the same", args: args{start_date: "2024-01-01", end_date: "2024-02-02 23:59:59"}, want: "", want1: "", want2: "", wantErr: true},
		{name: "YYYY-MM-DD", args: args{start_date: "2024-01-01", end_date: "2024-02-02"}, want: "2024-01-01", want1: "2024-02-02", want2: time.DateOnly, wantErr: false},
		{name: "YYYY-MM-DD hh:mm:ss", args: args{start_date: "2024-01-01 05:13:59", end_date: "2024-02-02 23:59:59"}, want: "2024-01-01 05:13:59", want1: "2024-02-02 23:59:59", want2: time.DateTime, wantErr: false},
		{name: "End of Month", args: args{start_date: "2024-02-01 05:13:59", end_date: "eom"}, want: "2024-02-01 05:13:59", want1: "2024-02-29 05:13:59", want2: time.DateTime, wantErr: false},
		{name: "End of Year", args: args{start_date: "2024-01-01 05:13:59", end_date: "eoy"}, want: "2024-01-01 05:13:59", want1: "2024-12-31 05:13:59", want2: time.DateTime, wantErr: false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDateRangeBounds(tt.args.start_date, tt.args.end_date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDateRangeBounds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.StartDate.Format(got.Format), tt.want) && !tt.wantErr {
				t.Errorf("GetDateRangeBounds() got = %v, want %v", got.StartDate.Format(got.Format), tt.want)
				return
			}
			if !reflect.DeepEqual(got.EndDate.Format(got.Format), tt.want1) && !tt.wantErr {
				t.Errorf("GetDateRangeBounds() got.EndDate = %v, want %v", got.EndDate.Format(got.Format), tt.want1)
				return
			}
			if got.Format != tt.want2 && !tt.wantErr {
				t.Errorf("GetDateRangeBounds() got.Format = %v, want %v", got.Format, tt.want2)
				return
			}
			
			// fmt.Println("OK:", got.StartDate.Format(got.Format), "->", got.EndDate.Format(got.Format))
		})
	}
}

func Benchmark_parseDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseDate("2023-12-01 22:04:44")
	}
}

func Benchmark_GetDateRangeBounds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDateRangeBounds("2023-12-01 22:04:44", "2024-12-01 22:04:44")
	}
}
