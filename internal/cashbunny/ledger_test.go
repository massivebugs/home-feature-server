package cashbunny

import (
	"reflect"
	"testing"

	"github.com/Rhymond/go-money"
)

func TestLedger_GetProfitLoss(t *testing.T) {
	type fields struct {
		accounts []*Account
	}
	tests := []struct {
		name         string
		fields       fields
		wantRevenues CurrencySums
		wantExpenses CurrencySums
		wantSums     CurrencySums
	}{
		{
			name: "success",
			fields: fields{
				accounts: []*Account{
					{
						Category: AccountCategoryRevenues,
						IncomingTransactions: []*Transaction{
							{Amount: money.New(500, money.JPY)},
							{Amount: money.New(500, money.JPY)},
							{Amount: money.New(500, money.CAD)},
							{Amount: money.New(500, money.CAD)},
						},
					},
					{
						Category: AccountCategoryExpenses,
						IncomingTransactions: []*Transaction{
							{Amount: money.New(700, money.JPY)},
							{Amount: money.New(700, money.CAD)},
						},
					},
				},
			},
			wantRevenues: NewCurrencySums([]*money.Money{
				money.New(1000, money.JPY),
				money.New(1000, money.CAD),
			}),
			wantExpenses: NewCurrencySums([]*money.Money{
				money.New(700, money.JPY),
				money.New(700, money.CAD),
			}),
			wantSums: NewCurrencySums([]*money.Money{
				money.New(300, money.JPY),
				money.New(300, money.CAD),
			}),
		},
		{
			name: "success when no accounts",
			fields: fields{
				accounts: []*Account{},
			},
			wantRevenues: NewCurrencySums(nil),
			wantExpenses: NewCurrencySums(nil),
			wantSums:     NewCurrencySums(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Ledger{
				accounts: tt.fields.accounts,
			}
			gotRevenues, gotExpenses, gotSums := l.GetProfitLoss()
			if !reflect.DeepEqual(gotRevenues, tt.wantRevenues) {
				t.Errorf("Ledger.GetProfitLoss() gotRevenues = %v, want %v", gotRevenues, tt.wantRevenues)
			}
			if !reflect.DeepEqual(gotExpenses, tt.wantExpenses) {
				t.Errorf("Ledger.GetProfitLoss() gotExpenses = %v, want %v", gotExpenses, tt.wantExpenses)
			}
			if !reflect.DeepEqual(gotSums, tt.wantSums) {
				t.Errorf("Ledger.GetProfitLoss() gotSums = %v, want %v", gotSums, tt.wantSums)
			}
		})
	}
}
