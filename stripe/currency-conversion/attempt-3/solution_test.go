package attempt3

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		s    string
		to   string
		from string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				s:    "USD:CAD:DHL:5,USD:GBP:FEDX:10",
				to:   "CAD",
				from: "USD",
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				s:    "USD:CAD:DHL:5,USD:GBP:FEDX:10",
				to:   "PKR",
				from: "USD",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.args.s, tt.args.to, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtMostOneHop(t *testing.T) {
	type args struct {
		s    string
		to   string
		from string
	}
	tests := []struct {
		name    string
		args    args
		want    *Res
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				s:    "USD:CAD:DHL:5,USD:GBP:FEDX:10",
				to:   "CAD",
				from: "GBP",
			},
			want: &Res{Cost: 15, Delivery: []string{"FEDX", "DHL"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtMostOneHop(tt.args.s, tt.args.to, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("AtMostOneHop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AtMostOneHop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinCost(t *testing.T) {
	type args struct {
		s    string
		to   string
		from string
	}
	tests := []struct {
		name    string
		args    args
		want    *Res
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				s:    "USD:CAD:DHL:5,USD:GBP:FEDX:10,CAD:GBP:TCS:10",
				to:   "CAD",
				from: "GBP",
			},
			want: &Res{Cost: 10, Delivery: []string{"TCS"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinCost(tt.args.s, tt.args.to, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinCost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
