package utils

import (
	"reflect"
	"testing"
	
	"seqd/models"
)

func Test_ParseArgs(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Args
		want1   bool
		want2   bool
		wantErr bool
	}{
		{name: "No args", args: args{a: []string{"seqd"}}, want: models.Args{}, want1: false, want2: false, wantErr: true},
		{name: "Wrong single argument", args: args{a: []string{"seqd", "--kek"}}, want: models.Args{}, want1: false, want2: false, wantErr: true},
		{name: "Too many arguments", args: args{a: []string{"seqd", "--kek", "--kek", "--kek", "--kek"}}, want: models.Args{}, want1: false, want2: false, wantErr: true},
		{name: "Help", args: args{a: []string{"seqd", "--help"}}, want: models.Args{}, want1: true, want2: false, wantErr: false},
		{name: "Version", args: args{a: []string{"seqd", "--version"}}, want: models.Args{}, want1: false, want2: true, wantErr: false},
		{name: "Full set of args", args: args{a: []string{"seqd", "-Y", "2023-12-31", "2024-01-01"}}, want: models.Args{StartDateTime: "2023-12-31", EndDateTime: "2024-01-01", IncreaseByYear: true}, want1: false, want2: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ParseArgs(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 && !tt.wantErr {
				t.Errorf("ParseArgs() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 && !tt.wantErr {
				t.Errorf("ParseArgs() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
