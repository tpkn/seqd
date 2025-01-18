package core

import (
	"reflect"
	"testing"
)

func Test_ParseArgs(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name    string
		args    args
		want    Args
		want1   bool
		want2   bool
		wantErr bool
	}{
		{name: "No args", args: args{a: []string{"seqd"}}, want: Args{}, wantErr: true},
		{name: "Wrong single argument", args: args{a: []string{"seqd", "--kek"}}, want: Args{}, wantErr: true},
		{name: "Too many arguments", args: args{a: []string{"seqd", "--kek", "--kek", "--kek", "--kek", "--kek"}}, want: Args{}, wantErr: true},
		{name: "Help", args: args{a: []string{"seqd", "--help"}}, want: Args{Help: true}, wantErr: false},
		{name: "Version", args: args{a: []string{"seqd", "--version"}}, want: Args{Version: true}, wantErr: false},
		{name: "Full set of args", args: args{a: []string{"seqd", "-Y", "2023-12-31", "2024-01-01"}}, want: Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByYear: true, ReversedOrder: false}, wantErr: false},
		{name: "Full set of args (reversed)", args: args{a: []string{"seqd", "-Y", "2023-12-31", "2024-01-01", "-r"}}, want: Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByYear: true, ReversedOrder: true}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
