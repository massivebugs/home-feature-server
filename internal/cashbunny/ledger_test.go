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
						category: AccountCategoryRevenues,
						incomingTransactions: []*Transaction{
							{amount: money.New(100, money.JPY)},
							{amount: money.New(100, money.CAD)},
						},
						outgoingTransactions: []*Transaction{
							{amount: money.New(500, money.JPY)},
							{amount: money.New(500, money.CAD)},
						},
					},
					{
						category: AccountCategoryExpenses,
						incomingTransactions: []*Transaction{
							{amount: money.New(500, money.JPY)},
							{amount: money.New(500, money.CAD)},
						},
						outgoingTransactions: []*Transaction{
							{amount: money.New(100, money.JPY)},
							{amount: money.New(100, money.CAD)},
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
			gotRevenues, gotExpenses, gotSums := l.getProfitLoss(tt.args.from, tt.args.to)
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
