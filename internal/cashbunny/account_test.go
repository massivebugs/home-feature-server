package cashbunny

import (
	"reflect"
	"testing"

	"github.com/Rhymond/go-money"
)

func TestAccount_SumTransactions(t *testing.T) {
	type args struct {
		account *Account
	}
	tests := []struct {
		name string
		args args
		want CurrencySums
	}{
		{
			name: "success",
			args: args{
				account: &Account{
					IncomingTransactions: []*Transaction{
						{Amount: money.New(500, money.JPY)},
						{Amount: money.New(500, money.JPY)},
						{Amount: money.New(500, money.CAD)},
						{Amount: money.New(500, money.CAD)},
					},
				},
			},
			want: NewCurrencySums([]*money.Money{
				money.New(1000, money.JPY),
				money.New(1000, money.CAD),
			}),
		},
		{
			name: "success transactions is nil",
			args: args{
				account: &Account{
					IncomingTransactions: nil,
				},
			},
			want: NewCurrencySums(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.account.SumTransactions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.SumTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
