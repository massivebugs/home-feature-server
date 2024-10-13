package cashbunny

import (
	"reflect"
	"testing"

	"github.com/Rhymond/go-money"
)

func TestNewCurrencySums(t *testing.T) {
	want := CurrencySums(map[string]*money.Money{
		money.JPY: money.New(1500, money.JPY),
	})

	type args struct {
		sums []*money.Money
	}
	tests := []struct {
		name string
		args args
		want CurrencySums
	}{
		{
			name: "success",
			args: args{
				sums: []*money.Money{money.New(1500, money.JPY)},
			},
			want: want,
		},
		{
			name: "success when arg sums is nil",
			args: args{
				sums: nil,
			},
			want: CurrencySums{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrencySums(tt.args.sums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCurrencySums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencySums_Add(t *testing.T) {
	type args struct {
		amount *money.Money
	}
	tests := []struct {
		name    string
		m       CurrencySums
		args    args
		wantErr bool
		want    CurrencySums
	}{
		{
			name: "success add to existing currency",
			m:    NewCurrencySums([]*money.Money{money.New(1000, money.JPY)}),
			args: args{
				amount: money.New(500, money.JPY),
			},
			want: NewCurrencySums([]*money.Money{money.New(1500, money.JPY)}),
		},
		{
			name: "success add to new currency",
			m:    NewCurrencySums([]*money.Money{money.New(1000, money.JPY)}),
			args: args{
				amount: money.New(500, money.CAD),
			},
			want: NewCurrencySums([]*money.Money{money.New(1000, money.JPY), money.New(500, money.CAD)}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.add(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CurrencySums.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}
