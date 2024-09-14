package cashbunny

import (
	"reflect"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
)

func TestLedger_GetProfitLoss(t *testing.T) {
	type args struct {
		from *time.Time
		to   *time.Time
	}

	type fields struct {
		accounts []*Account
	}

	tests := []struct {
		name         string
		args         args
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
							{Amount: money.New(100, money.JPY)},
							{Amount: money.New(100, money.CAD)},
						},
						OutgoingTransactions: []*Transaction{
							{Amount: money.New(500, money.JPY)},
							{Amount: money.New(500, money.CAD)},
						},
					},
					{
						Category: AccountCategoryExpenses,
						IncomingTransactions: []*Transaction{
							{Amount: money.New(500, money.JPY)},
							{Amount: money.New(500, money.CAD)},
						},
						OutgoingTransactions: []*Transaction{
							{Amount: money.New(100, money.JPY)},
							{Amount: money.New(100, money.CAD)},
						},
					},
				},
			},
			wantRevenues: NewCurrencySums([]*money.Money{
				money.New(400, money.JPY),
				money.New(400, money.CAD),
			}),
			wantExpenses: NewCurrencySums([]*money.Money{
				money.New(400, money.JPY),
				money.New(400, money.CAD),
			}),
			wantSums: NewCurrencySums([]*money.Money{
				money.New(0, money.JPY),
				money.New(0, money.CAD),
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
			gotRevenues, gotExpenses, gotSums := l.GetProfitLoss(tt.args.from, tt.args.to)
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
