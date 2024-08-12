package cashbunny

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
	"github.com/massivebugs/home-feature-server/internal/api"
)

var (
	SupportedCurrencyCodes = []string{
		money.AED,
		money.AFN,
		money.ALL,
		money.AMD,
		money.ANG,
		money.AOA,
		money.ARS,
		money.AUD,
		money.AWG,
		money.AZN,
		money.BAM,
		money.BBD,
		money.BDT,
		money.BGN,
		money.BHD,
		money.BIF,
		money.BMD,
		money.BND,
		money.BOB,
		money.BRL,
		money.BSD,
		money.BTN,
		money.BWP,
		money.BYN,
		money.BYR,
		money.BZD,
		money.CAD,
		money.CDF,
		money.CHF,
		money.CLF,
		money.CLP,
		money.CNY,
		money.COP,
		money.CRC,
		money.CUC,
		money.CUP,
		money.CVE,
		money.CZK,
		money.DJF,
		money.DKK,
		money.DOP,
		money.DZD,
		money.EEK,
		money.EGP,
		money.ERN,
		money.ETB,
		money.EUR,
		money.FJD,
		money.FKP,
		money.GBP,
		money.GEL,
		money.GGP,
		money.GHC,
		money.GHS,
		money.GIP,
		money.GMD,
		money.GNF,
		money.GTQ,
		money.GYD,
		money.HKD,
		money.HNL,
		money.HRK,
		money.HTG,
		money.HUF,
		money.IDR,
		money.ILS,
		money.IMP,
		money.INR,
		money.IQD,
		money.IRR,
		money.ISK,
		money.JEP,
		money.JMD,
		money.JOD,
		money.JPY,
		money.KES,
		money.KGS,
		money.KHR,
		money.KMF,
		money.KPW,
		money.KRW,
		money.KWD,
		money.KYD,
		money.KZT,
		money.LAK,
		money.LBP,
		money.LKR,
		money.LRD,
		money.LSL,
		money.LTL,
		money.LVL,
		money.LYD,
		money.MAD,
		money.MDL,
		money.MGA,
		money.MKD,
		money.MMK,
		money.MNT,
		money.MOP,
		money.MRU,
		money.MUR,
		money.MVR,
		money.MWK,
		money.MXN,
		money.MYR,
		money.MZN,
		money.NAD,
		money.NGN,
		money.NIO,
		money.NOK,
		money.NPR,
		money.NZD,
		money.OMR,
		money.PAB,
		money.PEN,
		money.PGK,
		money.PHP,
		money.PKR,
		money.PLN,
		money.PYG,
		money.QAR,
		money.RON,
		money.RSD,
		money.RUB,
		money.RUR,
		money.RWF,
		money.SAR,
		money.SBD,
		money.SCR,
		money.SDG,
		money.SEK,
		money.SGD,
		money.SHP,
		money.SKK,
		money.SLE,
		money.SLL,
		money.SOS,
		money.SRD,
		money.SSP,
		money.STD,
		money.STN,
		money.SVC,
		money.SYP,
		money.SZL,
		money.THB,
		money.TJS,
		money.TMT,
		money.TND,
		money.TOP,
		money.TRL,
		money.TRY,
		money.TTD,
		money.TWD,
		money.TZS,
		money.UAH,
		money.UGX,
		money.USD,
		money.UYU,
		money.UZS,
		money.VEF,
		money.VES,
		money.VND,
		money.VUV,
		money.WST,
		money.XAF,
		money.XAG,
		money.XAU,
		money.XCD,
		money.XDR,
		money.XOF,
		money.XPF,
		money.YER,
		money.ZAR,
		money.ZMW,
		money.ZWD,
		money.ZWL,
	}
)

type Cashbunny struct {
	db            *sql.DB
	cashbunnyRepo cashbunny_repository.Querier
}

func NewCashbunny(db *sql.DB, cashbunnyRepo cashbunny_repository.Querier) *Cashbunny {
	return &Cashbunny{
		db:            db,
		cashbunnyRepo: cashbunnyRepo,
	}
}

func (s *Cashbunny) GetOverview(ctx context.Context, userID uint32, from time.Time, to time.Time) (*Ledger, error) {
	accountListData, err := s.cashbunnyRepo.ListAccounts(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(accountListData))
	for idx, ad := range accountListData {
		a, err := NewAccount(ad)
		if err != nil {
			return nil, err
		}
		accounts[idx] = a
	}

	transactionListData, err := s.cashbunnyRepo.ListTransactionsBetweenDates(
		ctx,
		s.db,
		cashbunny_repository.ListTransactionsBetweenDatesParams{
			UserID:           userID,
			FromTransactedAt: from,
			ToTransactedAt:   to,
		},
	)
	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, len(transactionListData))
	for idx, td := range transactionListData {
		tr, err := NewTransaction(td)
		if err != nil {
			return nil, err
		}
		transactions[idx] = tr
	}

	for _, a := range accounts {
		for _, tr := range transactions {
			if tr.IsSourceAccount(a) {
				a.AddOutgoingTransaction(tr)
			} else if tr.IsDestinationAccount(a) {
				a.AddIncomingTransaction(tr)
			}
		}
	}

	return NewLedger(accounts), nil
}

func (s *Cashbunny) GetAllCurrencies(ctx context.Context) map[string]string {
	currencies := map[string]string{}
	for _, code := range SupportedCurrencyCodes {
		gmCurrency := money.GetCurrency(code)
		currencies[gmCurrency.Code] = gmCurrency.Grapheme
	}
	return currencies
}

func (s *Cashbunny) GetUserPreferences(ctx context.Context, userID uint32) (*UserPreferences, error) {
	ucData, err := s.cashbunnyRepo.ListUserCurrencies(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	upData, err := s.cashbunnyRepo.GetUserPreferenceByUserID(ctx, s.db, userID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, api.NewAPIError(api.CodeNotFound, errors.New("user preferences hasn't been created yet"))
	}

	up := NewUserPreferences(upData, ucData)

	return up, nil
}

func (s *Cashbunny) CreateDefaultUserPreferences(ctx context.Context, userID uint32) (*UserPreferences, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = s.cashbunnyRepo.CreateUserPreferences(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	_, err = s.cashbunnyRepo.CreateUserCurrency(ctx, tx, cashbunny_repository.CreateUserCurrencyParams{
		UserID:       userID,
		CurrencyCode: money.CAD,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.cashbunnyRepo.CreateUserCurrency(ctx, tx, cashbunny_repository.CreateUserCurrencyParams{
		UserID:       userID,
		CurrencyCode: money.JPY,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	ucData, err := s.cashbunnyRepo.ListUserCurrencies(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	upData, err := s.cashbunnyRepo.GetUserPreferenceByUserID(ctx, s.db, userID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, api.NewAPIError(api.CodeNotFound, errors.New("user preferences hasn't been created yet"))
	}

	up := NewUserPreferences(upData, ucData)

	return up, nil
}

func (s *Cashbunny) CreateAccount(ctx context.Context, userID uint32, req *CreateAccountRequestDTO) error {
	accountType, err := GetAccountTypeForCategory(AccountCategory(req.Category))
	if err != nil {
		// TODO: Flag errors
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.cashbunnyRepo.IncrementIndex(
		ctx,
		tx,
		cashbunny_repository.IncrementIndexParams{
			UserID:     userID,
			OrderIndex: req.OrderIndex,
		},
	)
	if err != nil {
		return err
	}

	_, err = s.cashbunnyRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_repository.CreateAccountParams{
			UserID:      userID,
			Category:    req.Category,
			Name:        req.Name,
			Description: req.Description,
			Balance:     0,
			Currency:    req.Currency,
			Type:        string(accountType),
			OrderIndex:  req.OrderIndex,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32) ([]*Account, error) {
	data, err := s.cashbunnyRepo.ListAccounts(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(data))
	for idx, d := range data {
		a, err := NewAccount(d)
		if err != nil {
			return nil, err
		}
		accounts[idx] = a
	}

	return accounts, nil
}

func (s *Cashbunny) DeleteAccount(ctx context.Context, userID uint32, accountID uint32) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.cashbunnyRepo.DeleteAccount(ctx, tx, cashbunny_repository.DeleteAccountParams{
		UserID: userID,
		ID:     accountID,
	})
	if err != nil {
		return err
	}

	err = s.cashbunnyRepo.DeleteTransactionsByAccountID(ctx, tx, cashbunny_repository.DeleteTransactionsByAccountIDParams{
		UserID:    userID,
		AccountID: accountID,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Cashbunny) CreateTransaction(ctx context.Context, userID uint32, req *CreateTransactionRequestDTO) error {
	// Check if source account belong to this user
	saResult, err := s.cashbunnyRepo.GetAccountByID(ctx, s.db, cashbunny_repository.GetAccountByIDParams{
		UserID: userID,
		ID:     req.SourceAccountID,
	})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// TODO: flag action
		return err
	} else if err != nil {
		return err
	}

	sa, err := NewAccount(saResult)
	if err != nil {
		return err
	}

	// Check if destination account belong to this user
	daResult, err := s.cashbunnyRepo.GetAccountByID(ctx, s.db, cashbunny_repository.GetAccountByIDParams{
		UserID: userID,
		ID:     req.DestinationAccountID,
	})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// TODO: flag action
		return err
	} else if err != nil {
		return err
	}

	da, err := NewAccount(daResult)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	transactedAt, err := time.Parse(time.DateTime, req.TransactedAt)
	if err != nil {
		return err
	}

	if sa.Type == AccountTypeCredit {
		m, err := sa.Balance.Add(money.NewFromFloat(req.Amount, req.Currency))
		if err != nil {
			return err
		}
		sa.Balance = m
	} else {
		m, err := sa.Balance.Subtract(money.NewFromFloat(req.Amount, req.Currency))
		if err != nil {
			return err
		}
		sa.Balance = m
	}

	if da.Type == AccountTypeCredit {
		m, err := da.Balance.Subtract(money.NewFromFloat(req.Amount, req.Currency))
		if err != nil {
			return err
		}
		da.Balance = m
	} else {
		m, err := da.Balance.Add(money.NewFromFloat(req.Amount, req.Currency))
		if err != nil {
			return err
		}
		da.Balance = m
	}

	err = s.cashbunnyRepo.UpdateAccountBalance(ctx, tx, cashbunny_repository.UpdateAccountBalanceParams{
		UserID:  userID,
		ID:      sa.ID,
		Balance: sa.Balance.AsMajorUnits(),
	})
	if err != nil {
		return err
	}

	err = s.cashbunnyRepo.UpdateAccountBalance(ctx, tx, cashbunny_repository.UpdateAccountBalanceParams{
		UserID:  userID,
		ID:      da.ID,
		Balance: da.Balance.AsMajorUnits(),
	})
	if err != nil {
		return err
	}

	_, err = s.cashbunnyRepo.CreateTransaction(ctx, tx, cashbunny_repository.CreateTransactionParams{
		UserID:        userID,
		SrcAccountID:  req.SourceAccountID,
		DestAccountID: req.DestinationAccountID,
		Description:   req.Description,
		Amount:        req.Amount,
		Currency:      req.Currency,
		TransactedAt:  transactedAt,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Cashbunny) ListTransactions(ctx context.Context, userID uint32) ([]*Transaction, error) {
	data, err := s.cashbunnyRepo.ListTransactions(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	// Preload accounts
	accountIds := mapset.NewSet[uint32]()
	for _, d := range data {
		accountIds.Add(d.SrcAccountID)
		accountIds.Add(d.DestAccountID)
	}

	accounts, err := s.cashbunnyRepo.ListAccountsByIDs(ctx, s.db, cashbunny_repository.ListAccountsByIDsParams{
		UserID: userID,
		IDs:    accountIds.ToSlice(),
	})
	if err != nil {
		return nil, err
	}

	trs := make([]*Transaction, len(data))
	for idx, d := range data {
		srcAccountDa := SliceFind(accounts, func(a *cashbunny_repository.CashbunnyAccount) bool { return a.ID == d.SrcAccountID })
		destAccountDa := SliceFind(accounts, func(a *cashbunny_repository.CashbunnyAccount) bool { return a.ID == d.DestAccountID })

		a, err := NewTransaction(d)
		if err != nil {
			return nil, err
		}

		sa, err := NewAccount(*srcAccountDa)
		if err != nil {
			return nil, err
		}

		da, err := NewAccount(*destAccountDa)
		if err != nil {
			return nil, err
		}

		a.SourceAccount = sa
		a.DestinationAccount = da

		trs[idx] = a
	}

	return trs, nil
}

func (s *Cashbunny) DeleteTransaction(ctx context.Context, userID uint32, transactionID uint32) error {
	trResult, err := s.cashbunnyRepo.GetTransactionByID(ctx, s.db, cashbunny_repository.GetTransactionByIDParams{
		UserID: userID,
		ID:     transactionID,
	})
	if err != nil {
		return err
	}

	// Check if source account belong to this user
	saResult, err := s.cashbunnyRepo.GetAccountByID(ctx, s.db, cashbunny_repository.GetAccountByIDParams{
		UserID: userID,
		ID:     trResult.SrcAccountID,
	})
	if err != nil {
		return err
	}

	sa, err := NewAccount(saResult)
	if err != nil {
		return err
	}

	// Check if destination account belong to this user
	daResult, err := s.cashbunnyRepo.GetAccountByID(ctx, s.db, cashbunny_repository.GetAccountByIDParams{
		UserID: userID,
		ID:     trResult.DestAccountID,
	})
	if err != nil {
		return err
	}

	da, err := NewAccount(daResult)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if sa.Type == AccountTypeCredit {
		m, err := sa.Balance.Subtract(money.NewFromFloat(trResult.Amount, trResult.Currency))
		if err != nil {
			return err
		}
		sa.Balance = m
	} else {
		m, err := sa.Balance.Add(money.NewFromFloat(trResult.Amount, trResult.Currency))
		if err != nil {
			return err
		}
		sa.Balance = m
	}

	if da.Type == AccountTypeCredit {
		m, err := da.Balance.Add(money.NewFromFloat(trResult.Amount, trResult.Currency))
		if err != nil {
			return err
		}
		da.Balance = m
	} else {
		m, err := da.Balance.Subtract(money.NewFromFloat(trResult.Amount, trResult.Currency))
		if err != nil {
			return err
		}
		da.Balance = m
	}

	err = s.cashbunnyRepo.UpdateAccountBalance(ctx, tx, cashbunny_repository.UpdateAccountBalanceParams{
		UserID:  userID,
		ID:      sa.ID,
		Balance: sa.Balance.AsMajorUnits(),
	})
	if err != nil {
		return err
	}

	err = s.cashbunnyRepo.UpdateAccountBalance(ctx, tx, cashbunny_repository.UpdateAccountBalanceParams{
		UserID:  userID,
		ID:      da.ID,
		Balance: da.Balance.AsMajorUnits(),
	})
	if err != nil {
		return err
	}

	err = s.cashbunnyRepo.DeleteTransaction(ctx, tx, cashbunny_repository.DeleteTransactionParams{
		UserID: userID,
		ID:     transactionID,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}
