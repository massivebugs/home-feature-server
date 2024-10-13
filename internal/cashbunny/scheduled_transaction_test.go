package cashbunny

import (
	"reflect"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/teambition/rrule-go"
)

func TestScheduledTransaction_ToTransactions(t *testing.T) {
	rule, err := rrule.NewRRule(
		rrule.ROption{
			Freq:     rrule.DAILY,
			Dtstart:  time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
			Interval: 1,
			Until:    time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	)
	if err != nil {
		t.Fatal("Failed on rrule.NewRRule()")
	}

	type fields struct {
		SrcAccountID       uint32
		DestAccountID      uint32
		Description        string
		Amount             *money.Money
		RecurrenceRule     *RecurrenceRule
		SourceAccount      *Account
		DestinationAccount *Account
	}
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Transaction
	}{
		{
			name: "success - all transactions between from and to are generated",
			fields: fields{
				SrcAccountID:  1,
				DestAccountID: 1,
				Description:   "test scheduled transaction",
				Amount:        money.NewFromFloat(100, "JPY"),
				RecurrenceRule: &RecurrenceRule{
					rule: rule,
				},
				SourceAccount: &Account{
					name: "test source account",
				},
				DestinationAccount: &Account{
					name: "test destination account",
				},
			},
			args: args{
				from: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC),
			},
			want: []*Transaction{
				{
					srcAccountID:  1,
					destAccountID: 1,
					description:   "test scheduled transaction",
					amount:        money.NewFromFloat(100, "JPY"),
					transactedAt:  time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
					sourceAccount: &Account{
						name: "test source account",
					},
					destinationAccount: &Account{
						name: "test destination account",
					},
					scheduledTransaction: &ScheduledTransaction{
						srcAccountID:  1,
						destAccountID: 1,
						description:   "test scheduled transaction",
						amount:        money.NewFromFloat(100, "JPY"),
						recurrenceRule: &RecurrenceRule{
							rule: rule,
						},
						sourceAccount: &Account{
							name: "test source account",
						},
						destinationAccount: &Account{
							name: "test destination account",
						},
					},
				},
				{
					srcAccountID:  1,
					destAccountID: 1,
					description:   "test scheduled transaction",
					amount:        money.NewFromFloat(100, "JPY"),
					transactedAt:  time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC),
					sourceAccount: &Account{
						name: "test source account",
					},
					destinationAccount: &Account{
						name: "test destination account",
					},
					scheduledTransaction: &ScheduledTransaction{
						srcAccountID:  1,
						destAccountID: 1,
						description:   "test scheduled transaction",
						amount:        money.NewFromFloat(100, "JPY"),
						recurrenceRule: &RecurrenceRule{
							rule: rule,
						},
						sourceAccount: &Account{
							name: "test source account",
						},
						destinationAccount: &Account{
							name: "test destination account",
						},
					},
				},
				{
					srcAccountID:  1,
					destAccountID: 1,
					description:   "test scheduled transaction",
					amount:        money.NewFromFloat(100, "JPY"),
					transactedAt:  time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC),
					sourceAccount: &Account{
						name: "test source account",
					},
					destinationAccount: &Account{
						name: "test destination account",
					},
					scheduledTransaction: &ScheduledTransaction{
						srcAccountID:  1,
						destAccountID: 1,
						description:   "test scheduled transaction",
						amount:        money.NewFromFloat(100, "JPY"),
						recurrenceRule: &RecurrenceRule{
							rule: rule,
						},
						sourceAccount: &Account{
							name: "test source account",
						},
						destinationAccount: &Account{
							name: "test destination account",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &ScheduledTransaction{
				srcAccountID:       tt.fields.SrcAccountID,
				destAccountID:      tt.fields.DestAccountID,
				description:        tt.fields.Description,
				amount:             tt.fields.Amount,
				recurrenceRule:     tt.fields.RecurrenceRule,
				sourceAccount:      tt.fields.SourceAccount,
				destinationAccount: tt.fields.DestinationAccount,
			}
			if got := st.toTransactions(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScheduledTransaction.ToTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
