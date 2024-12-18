package cashbunny

import (
	"reflect"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
)

func TestAccount_GetType(t *testing.T) {
	type fields struct {
		Category AccountCategory
	}
	tests := []struct {
		name   string
		fields fields
		want   AccountType
	}{
		{
			name: "success - assets category is debit",
			fields: fields{
				Category: AccountCategoryAssets,
			},
			want: AccountTypeDebit,
		},
		{
			name: "success - expenses category is debit",
			fields: fields{
				Category: AccountCategoryExpenses,
			},
			want: AccountTypeDebit,
		},
		{
			name: "success - liabilities category is credit",
			fields: fields{
				Category: AccountCategoryLiabilities,
			},
			want: AccountTypeCredit,
		},
		{
			name: "success - revenues category is credit",
			fields: fields{
				Category: AccountCategoryRevenues,
			},
			want: AccountTypeCredit,
		},
		{
			name: "zero-value - anything else returns empty string",
			fields: fields{
				Category: "foo",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Category: tt.fields.Category,
			}
			if got := a.GetType(); got != tt.want {
				t.Errorf("Account.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_SumTransactions(t *testing.T) {
	from := time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 11, 5, 0, 0, 0, 0, time.UTC)

	type args struct {
		account *Account
		from    *time.Time
		to      *time.Time
	}
	tests := []struct {
		name string
		args args
		want CurrencySums
	}{
		{
			name: "success - credit account",
			args: args{
				account: &Account{
					Category: AccountCategoryAssets,
					IncomingTransactions: []*Transaction{
						{Amount: money.New(500, money.JPY)},
						{Amount: money.New(500, money.JPY)},
						{Amount: money.New(500, money.CAD)},
						{Amount: money.New(500, money.CAD)},
					},
					OutgoingTransactions: []*Transaction{
						{Amount: money.New(100, money.JPY)},
						{Amount: money.New(100, money.JPY)},
						{Amount: money.New(100, money.CAD)},
						{Amount: money.New(100, money.CAD)},
					},
				},
			},
			want: NewCurrencySums([]*money.Money{
				money.New(800, money.JPY),
				money.New(800, money.CAD),
			}),
		},
		{
			name: "success - debit account",
			args: args{
				account: &Account{
					Category: AccountCategoryLiabilities,
					IncomingTransactions: []*Transaction{
						{Amount: money.New(100, money.JPY)},
						{Amount: money.New(100, money.JPY)},
						{Amount: money.New(100, money.CAD)},
						{Amount: money.New(100, money.CAD)},
					},
					OutgoingTransactions: []*Transaction{
						{Amount: money.New(500, money.JPY)},
						{Amount: money.New(500, money.JPY)},
						{Amount: money.New(500, money.CAD)},
						{Amount: money.New(500, money.CAD)},
					},
				},
			},
			want: NewCurrencySums([]*money.Money{
				money.New(800, money.JPY),
				money.New(800, money.CAD),
			}),
		},
		{
			name: "success - includes transactions after(inclusive) a specific date",
			args: args{
				account: &Account{
					Category: AccountCategoryAssets,
					IncomingTransactions: []*Transaction{
						{
							Amount:       money.New(500, money.JPY),
							TransactedAt: time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(500, money.JPY),
							TransactedAt: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(500, money.CAD),
							TransactedAt: time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC),
						},
					},
					OutgoingTransactions: []*Transaction{
						{
							Amount:       money.New(200, money.JPY),
							TransactedAt: time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(200, money.JPY),
							TransactedAt: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(200, money.CAD),
							TransactedAt: time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				from: &from,
			},
			want: NewCurrencySums([]*money.Money{
				money.New(300, money.JPY),
				money.New(300, money.CAD),
			}),
		},
		{
			name: "success - includes transactions before(inclusive) a specific date",
			args: args{
				account: &Account{
					Category: AccountCategoryAssets,
					IncomingTransactions: []*Transaction{
						{
							Amount:       money.New(500, money.JPY),
							TransactedAt: time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(500, money.CAD),
							TransactedAt: time.Date(2024, 11, 5, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(500, money.JPY),
							TransactedAt: time.Date(2024, 11, 6, 0, 0, 0, 0, time.UTC),
						},
					},
					OutgoingTransactions: []*Transaction{
						{
							Amount:       money.New(200, money.JPY),
							TransactedAt: time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(200, money.CAD),
							TransactedAt: time.Date(2024, 11, 5, 0, 0, 0, 0, time.UTC),
						},
						{
							Amount:       money.New(200, money.JPY),
							TransactedAt: time.Date(2024, 11, 6, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				to: &to,
			},
			want: NewCurrencySums([]*money.Money{
				money.New(300, money.JPY),
				money.New(300, money.CAD),
			}),
		},
		{
			name: "success - transactions is nil",
			args: args{
				account: &Account{
					Category:             AccountCategoryAssets,
					IncomingTransactions: nil,
				},
			},
			want: NewCurrencySums(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.account.sumTransactions(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.SumTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_AddIncomingTransaction(t *testing.T) {
	type args struct {
		tr *Transaction
	}
	tests := []struct {
		name string
		args args
		want []*Transaction
	}{
		{
			name: "success - add to empty",
			args: args{
				tr: &Transaction{
					Description: "foo",
				},
			},
			want: []*Transaction{
				{
					Description: "foo",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{}
			a.addIncomingTransaction(tt.args.tr)

			got := a.IncomingTransactions
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_AddOutgoingTransaction(t *testing.T) {
	type args struct {
		tr *Transaction
	}
	tests := []struct {
		name string
		args args
		want []*Transaction
	}{
		{
			name: "success - add to empty",
			args: args{
				tr: &Transaction{
					Description: "foo",
				},
			},
			want: []*Transaction{
				{
					Description: "foo",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{}
			a.addOutgoingTransaction(tt.args.tr)

			got := a.OutgoingTransactions
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
