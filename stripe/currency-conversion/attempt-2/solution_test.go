package currencyconversion

import (
	"reflect"
	"testing"
)

func TestParseCurrencies(t *testing.T) {
	type args struct {
		curr string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]Rate
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				curr: "USD:CAD:DHL:5,USD:GBP:FEDX:10",
			},
			want: map[string][]Rate{
				"USD": {{To: "CAD", Delivery: "DHL", Cost: 5}, {To: "GBP", Delivery: "FEDX", Cost: 10}},
				"CAD": {{To: "USD", Delivery: "DHL", Cost: 5}},
				"GBP": {{To: "USD", Delivery: "FEDX", Cost: 10}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseCurrencies(tt.args.curr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCurrencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertCurrency(t *testing.T) {
	ParseCurrencies("USD:CAD:DHL:5,USD:GBP:FEDX:10")
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "USD to CAD",
			args: args{
				from: "USD",
				to:   "CAD",
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "GBP to USD",
			args: args{
				from: "GBP",
				to:   "USD",
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Invalid currency",
			args: args{
				from: "PKR",
				to:   "USD",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertCurrency(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtMostOneHop(t *testing.T) {
	ParseCurrencies("USD:CAD:DHL:5,USD:GBP:FEDX:10")
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Res
		wantErr bool
	}{
		{
			name: "CAD to GBP",
			args: args{
				from: "CAD",
				to:   "GBP",
			},
			want: &Res{Delivery: []string{"DHL", "FEDX"}, Cost: 15},
		},
		{
			name: "USD to GBP",
			args: args{
				from: "USD",
				to:   "GBP",
			},
			want: &Res{Delivery: []string{"FEDX"}, Cost: 10},
		},
		{
			name: "Invalid conversion",
			args: args{
				from: "PKR",
				to:   "GBP",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtMostOneHop(tt.args.from, tt.args.to)
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

func TestMinAtMostOneHop(t *testing.T) {
	ParseCurrencies("USD:CAD:DHL:5,USD:GBP:FEDX:10,CAD:GBP:FEDX:2")
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Res
		wantErr bool
	}{
		{
			name: "CAD to GBP",
			args: args{
				from: "CAD",
				to:   "GBP",
			},
			want: &Res{Delivery: []string{"FEDX"}, Cost: 2},
		},
		{
			name: "USD to GBP",
			args: args{
				from: "USD",
				to:   "GBP",
			},
			want: &Res{Delivery: []string{"DHL", "FEDX"}, Cost: 7},
		},
		{
			name: "Invalid conversion",
			args: args{
				from: "PKR",
				to:   "GBP",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinAtMostOneHop(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinAtMostOneHop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinAtMostOneHop() = %v, want %v", got, tt.want)
			}
		})
	}
}
