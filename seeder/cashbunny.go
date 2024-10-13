package seeder

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
	"github.com/teambition/rrule-go"
)

func (s *Seeder) createCashbunnyDataForPublicUser(ctx context.Context, tx db.DB, userID uint32, now time.Time) error {
	// Create Accounts
	jpyCheckingAccountID,
		cadCheckingAccountID,
		jpyExpenseAccountID,
		cadExpenseAccountID,
		cadBigExpenseAccountID,
		jpyRevenueAccountID,
		cadRevenueAccountID,
		cadLiabilitiesAccountID,
		err := s.createCashbunnyPublicUserAccounts(ctx, tx, userID)

	if err != nil {
		return err
	}

	// Create Category for monthly income
	cadIncomeCategoryResult, err := s.querier.CreateTransactionCategory(ctx, tx, queries.CreateTransactionCategoryParams{
		UserID: userID,
		Name:   "Monthly income",
	})
	if err != nil {
		return err
	}
	cadIncomeCategoryID, err := cadIncomeCategoryResult.LastInsertId()
	if err != nil {
		return err
	}

	// Create Category for loan payment
	cadLoanPaymentCategoryResult, err := s.querier.CreateTransactionCategory(ctx, tx, queries.CreateTransactionCategoryParams{
		UserID: userID,
		Name:   "Loan payment",
	})
	if err != nil {
		return err
	}
	cadLoanPaymentCategoryID, err := cadLoanPaymentCategoryResult.LastInsertId()
	if err != nil {
		return err
	}

	// Create Category for recurring expense (Cost of Living)
	cadExpenseCategoryResult, err := s.querier.CreateTransactionCategory(ctx, tx, queries.CreateTransactionCategoryParams{
		UserID: userID,
		Name:   "Costs of living",
	})
	if err != nil {
		return err
	}
	cadExpenseCategoryID, err := cadExpenseCategoryResult.LastInsertId()
	if err != nil {
		return err
	}

	// Create Category Allocation for expense (Costs of living)
	_, err = s.querier.CreateTransactionCategoryAllocation(ctx, tx, queries.CreateTransactionCategoryAllocationParams{
		UserID:     userID,
		CategoryID: uint32(cadExpenseCategoryID),
		Amount:     3000,
		Currency:   money.CAD,
	})
	if err != nil {
		return err
	}

	// Create Scheduled Transactions
	// Monthly income schedule (CAD)
	cadIncomeScheduledTransactionResult, err := s.querier.CreateScheduledTransaction(ctx, tx, queries.CreateScheduledTransactionParams{
		UserID: uint32(userID),
		CategoryID: sql.NullInt32{
			Int32: int32(cadIncomeCategoryID),
			Valid: true,
		},
		SrcAccountID:  uint32(cadRevenueAccountID),
		DestAccountID: uint32(cadCheckingAccountID),
		Description:   "Monthly salary minus tax",
		Amount:        8000,
		Currency:      money.CAD,
	})
	if err != nil {
		return err
	}

	cadIncomeScheduledTransactionID, err := cadIncomeScheduledTransactionResult.LastInsertId()
	if err != nil {
		return err
	}

	cadIncomeRecurrenceRuleResult, err := s.querier.CreateRecurrenceRule(ctx, tx, queries.CreateRecurrenceRuleParams{
		Freq:     rrule.MONTHLY.String(),
		Dtstart:  now.AddDate(0, -11, 0),
		Count:    0,
		Interval: 1,
		Until:    now.AddDate(5, 0, 0),
	})
	if err != nil {
		return err
	}

	cadIncomeRecurrenceRuleID, err := cadIncomeRecurrenceRuleResult.LastInsertId()
	if err != nil {
		return err
	}

	_, err = s.querier.CreateScheduledTransactionRecurrenceRuleRelationship(ctx, tx, queries.CreateScheduledTransactionRecurrenceRuleRelationshipParams{
		ScheduledTransactionID: uint32(cadIncomeScheduledTransactionID),
		RecurrenceRuleID:       uint32(cadIncomeRecurrenceRuleID),
	})
	if err != nil {
		return err
	}

	// Monthly loan payment schedule (CAD)
	cadLoanPaymentScheduledTransactionResult, err := s.querier.CreateScheduledTransaction(ctx, tx, queries.CreateScheduledTransactionParams{
		UserID: uint32(userID),
		CategoryID: sql.NullInt32{
			Int32: int32(cadLoanPaymentCategoryID),
			Valid: true,
		},
		SrcAccountID:  uint32(cadCheckingAccountID),
		DestAccountID: uint32(cadLiabilitiesAccountID),
		Description:   "Car loan payment",
		Amount:        880,
		Currency:      money.CAD,
	})
	if err != nil {
		return err
	}

	cadLoanPaymentScheduledTransactionID, err := cadLoanPaymentScheduledTransactionResult.LastInsertId()
	if err != nil {
		return err
	}

	cadLoanPaymentRecurrenceRuleResult, err := s.querier.CreateRecurrenceRule(ctx, tx, queries.CreateRecurrenceRuleParams{
		Freq:     rrule.MONTHLY.String(),
		Dtstart:  now.AddDate(0, -8, -15),
		Count:    0,
		Interval: 1,
		Until:    now.AddDate(0, 36, 0),
	})
	if err != nil {
		return err
	}

	cadLoanPaymentRecurrenceRuleID, err := cadLoanPaymentRecurrenceRuleResult.LastInsertId()
	if err != nil {
		return err
	}

	_, err = s.querier.CreateScheduledTransactionRecurrenceRuleRelationship(ctx, tx, queries.CreateScheduledTransactionRecurrenceRuleRelationshipParams{
		ScheduledTransactionID: uint32(cadLoanPaymentScheduledTransactionID),
		RecurrenceRuleID:       uint32(cadLoanPaymentRecurrenceRuleID),
	})
	if err != nil {
		return err
	}

	// Recurring expense schedule (Groceries) (CAD)
	cadExpenseScheduledTransactionResult, err := s.querier.CreateScheduledTransaction(ctx, tx, queries.CreateScheduledTransactionParams{
		UserID: uint32(userID),
		CategoryID: sql.NullInt32{
			Int32: int32(cadExpenseCategoryID),
			Valid: true,
		},
		SrcAccountID:  uint32(cadCheckingAccountID),
		DestAccountID: uint32(cadExpenseAccountID),
		Description:   "Groceries",
		Amount:        250,
		Currency:      money.CAD,
	})
	if err != nil {
		return err
	}

	cadExpenseScheduledTransactionID, err := cadExpenseScheduledTransactionResult.LastInsertId()
	if err != nil {
		return err
	}

	cadExpenseRecurrenceRuleResult, err := s.querier.CreateRecurrenceRule(ctx, tx, queries.CreateRecurrenceRuleParams{
		Freq:     rrule.WEEKLY.String(),
		Dtstart:  now.AddDate(0, -11, 2),
		Count:    0,
		Interval: 2,
		Until:    time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		return err
	}

	cadExpenseRecurrenceRuleID, err := cadExpenseRecurrenceRuleResult.LastInsertId()
	if err != nil {
		return err
	}

	_, err = s.querier.CreateScheduledTransactionRecurrenceRuleRelationship(ctx, tx, queries.CreateScheduledTransactionRecurrenceRuleRelationshipParams{
		ScheduledTransactionID: uint32(cadExpenseScheduledTransactionID),
		RecurrenceRuleID:       uint32(cadExpenseRecurrenceRuleID),
	})
	if err != nil {
		return err
	}

	// Initial capital data (JPY)
	if _, err = s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(jpyRevenueAccountID),
		DestAccountID: uint32(jpyCheckingAccountID),
		Description:   "Initial capital",
		Amount:        5600000,
		Currency:      money.JPY,
		TransactedAt:  now.AddDate(-1, 0, 0),
	}); err != nil {
		return err
	}

	// Loan (CAD)
	if _, err = s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(cadLiabilitiesAccountID),
		DestAccountID: uint32(cadCheckingAccountID),
		Description:   "Car loan minus down payment",
		Amount:        32000,
		Currency:      money.CAD,
		TransactedAt:  now.AddDate(0, -9, 0),
	}); err != nil {
		return err
	}

	// Big payment made from loan (CAD)
	if _, err = s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(cadCheckingAccountID),
		DestAccountID: uint32(cadBigExpenseAccountID),
		Description:   "Bought a car!",
		Amount:        32000,
		Currency:      money.CAD,
		TransactedAt:  now.AddDate(0, -9, 0),
	}); err != nil {
		return err
	}

	// 1 year monthly income data (CAD)
	for i := 0; i < 12; i++ {
		if _, err := s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
			UserID:                 uint32(userID),
			ScheduledTransactionID: sql.NullInt32{Valid: true, Int32: int32(cadIncomeScheduledTransactionID)},
			CategoryID: sql.NullInt32{
				Int32: int32(cadIncomeCategoryID),
				Valid: true,
			},
			SrcAccountID:  uint32(cadRevenueAccountID),
			DestAccountID: uint32(cadCheckingAccountID),
			Description:   "Monthly salary minus tax",
			Amount:        8000,
			Currency:      money.CAD,
			TransactedAt:  now.AddDate(0, -i, 0),
		}); err != nil {
			return err
		}
	}

	// Loan payment (CAD)
	cadLoanPaymentDate := now.AddDate(0, -8, -15)
	for cadLoanPaymentDate.Before(now) {
		if _, err := s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
			UserID:                 uint32(userID),
			ScheduledTransactionID: sql.NullInt32{Valid: true, Int32: int32(cadLoanPaymentScheduledTransactionID)},
			CategoryID: sql.NullInt32{
				Int32: int32(cadLoanPaymentCategoryID),
				Valid: true,
			},
			SrcAccountID:  uint32(cadCheckingAccountID),
			DestAccountID: uint32(cadLiabilitiesAccountID),
			Description:   "Loan payment",
			Amount:        880,
			Currency:      money.CAD,
			TransactedAt:  cadLoanPaymentDate,
		}); err != nil {
			return err
		}
		cadLoanPaymentDate = cadLoanPaymentDate.AddDate(0, 1, 0)
	}

	// 1 year expsense data (Groceries) (CAD)
	cadExpenseDate := now.AddDate(0, -11, 2)
	for cadExpenseDate.Before(now) {
		if _, err := s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
			UserID:                 uint32(userID),
			ScheduledTransactionID: sql.NullInt32{Valid: true, Int32: int32(cadExpenseScheduledTransactionID)},
			CategoryID: sql.NullInt32{
				Int32: int32(cadExpenseCategoryID),
				Valid: true,
			},
			SrcAccountID:  uint32(cadCheckingAccountID),
			DestAccountID: uint32(cadExpenseAccountID),
			Description:   "Groceries",
			Amount:        250,
			Currency:      money.CAD,
			TransactedAt:  cadExpenseDate,
		}); err != nil {
			return err
		}
		cadExpenseDate = cadExpenseDate.AddDate(0, 0, 14)
	}

	// Discretionary expense (JPY)
	if _, err := s.querier.CreateTransaction(ctx, tx, queries.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(jpyCheckingAccountID),
		DestAccountID: uint32(jpyExpenseAccountID),
		Description:   "Cool stuff bought online!",
		Amount:        130000,
		Currency:      money.JPY,
		TransactedAt:  now.AddDate(0, -4, -12),
	}); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) createCashbunnyPublicUserAccounts(ctx context.Context, tx db.DB, userID uint32) (
	jpyCheckingAccountID int64,
	cadCheckingAccountID int64,
	jpyExpenseAccountID int64,
	cadExpenseAccountID int64,
	cadBigExpenseAccountID int64,
	jpyRevenueAccountID int64,
	cadRevenueAccountID int64,
	cadLiabilitiesAccountID int64,
	err error,
) {
	// Assets (JPY)
	jpyCheckingAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryAssets),
			Name:        "Checking account (ぷるるん銀行)",
			Description: "Money to order goods online.",
			Currency:    money.JPY,
		},
	)
	if err != nil {
		return
	}
	jpyCheckingAccountID, err = jpyCheckingAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Assets (CAD)
	cadCheckingAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryAssets),
			Name:        "Checking account (Northern Horizons Bank)",
			Description: "Store money for buying groceries, paying bills etc.",
			Currency:    money.CAD,
		},
	)
	if err != nil {
		return
	}
	cadCheckingAccountID, err = cadCheckingAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Expenses (JPY)
	jpyExpenseAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryExpenses),
			Name:        "半地下レトロ工場",
			Description: "Online store for Retro goods",
			Currency:    money.JPY,
		},
	)
	if err != nil {
		return
	}
	jpyExpenseAccountID, err = jpyExpenseAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Expenses (CAD)
	cadExpenseAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryExpenses),
			Name:        "Fresh Fields Market",
			Description: "Fresh, high-quality produce and general groceries",
			Currency:    money.CAD,
		},
	)
	if err != nil {
		return
	}
	cadExpenseAccountID, err = cadExpenseAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Big Expense requiring loan (CAD)
	cadBigExpenseAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryExpenses),
			Name:        "Happy road cars",
			Description: "Car dealership",
			Currency:    money.CAD,
		},
	)
	if err != nil {
		return
	}
	cadBigExpenseAccountID, err = cadBigExpenseAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Revenues (JPY)
	jpyRevenueAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryRevenues),
			Name:        "Initial capital",
			Description: "Source for money I have, but have not been tracking",
			Currency:    money.JPY,
		},
	)
	if err != nil {
		return
	}
	jpyRevenueAccountID, err = jpyRevenueAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Revenues (CAD)
	cadRevenueAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryRevenues),
			Name:        "NovaTech Solutions",
			Description: "My workplace",
			Currency:    money.CAD,
		},
	)
	if err != nil {
		return
	}
	cadRevenueAccountID, err = cadRevenueAccountResult.LastInsertId()
	if err != nil {
		return
	}

	// Liabilities (CAD)
	cadLiabilitiesAccountResult, err := s.querier.CreateAccount(
		ctx,
		tx,
		queries.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryLiabilities),
			Name:        "Car loans (Northern Horizons Bank)",
			Description: "Money I borrowed to purchase my car",
			Currency:    money.CAD,
		},
	)
	if err != nil {
		return
	}
	cadLiabilitiesAccountID, err = cadLiabilitiesAccountResult.LastInsertId()
	if err != nil {
		return
	}

	return
}
