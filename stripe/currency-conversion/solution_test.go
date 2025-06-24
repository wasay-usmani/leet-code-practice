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
				"USD": {{To: "CAD", Delivery: "DHL", Cost: 5},
					{To: "GBP", Delivery: "FEDX", Cost: 10}},
				"GBP": {{To: "USD", Delivery: "FEDX", Cost: 10}},
				"CAD": {{To: "USD", Delivery: "DHL", Cost: 5}},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCurrencies(tt.args.curr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCurrencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertCurrency(t *testing.T) {
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
			name: "Test 1",
			args: args{
				to:   "CAD",
				from: "USD",
			},
			want:    5,
			wantErr: false,
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
			name: "Test 1",
			args: args{
				to:   "CAD",
				from: "GBP",
			},
			want: &Res{
				Cost:     15,
				Delivery: []string{"DHL", "FEDX"},
			},
			wantErr: false,
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
			name: "Test 1",
			args: args{
				to:   "GBP",
				from: "CAD",
			},
			want: &Res{
				Cost:     10,
				Delivery: []string{"FEDX"},
			},
			wantErr: false,
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
