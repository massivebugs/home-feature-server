// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CashbunnyAccount Model representing any financial account such as bank account, credit card, wallet etc.
type CashbunnyAccount struct {
	Amount        *float64 `json:"amount"`
	AmountDisplay *string  `json:"amount_display"`
	Category      string   `json:"category"`
	CreatedAt     string   `json:"created_at"`
	Currency      string   `json:"currency"`
	Description   string   `json:"description"`
	Id            uint32   `json:"id"`
	Name          string   `json:"name"`
	OrderIndex    uint32   `json:"order_index"`
	Type          string   `json:"type"`
	UpdatedAt     string   `json:"updated_at"`
}

// CashbunnyOverview Overview of user's budget keeping.
type CashbunnyOverview struct {
	AssetAccounts []CashbunnyAccount `json:"asset_accounts"`

	// From Overview target start time
	From              string             `json:"from"`
	LiabilityAccounts []CashbunnyAccount `json:"liability_accounts"`

	// NetWorth Net worth per currency
	NetWorth map[string]string `json:"net_worth"`

	// ProfitLossSummary User's profit/loss summary within the timeframe
	ProfitLossSummary map[string]struct {
		Expense string `json:"expense"`
		Profit  string `json:"profit"`
		Revenue string `json:"revenue"`
	} `json:"profit_loss_summary"`

	// To Overview target end time
	To                        string                 `json:"to"`
	Transactions              []CashbunnyTransaction `json:"transactions"`
	TransactionsFromScheduled []CashbunnyTransaction `json:"transactions_from_scheduled"`
}

// CashbunnyRecurrenceRule Model representing iCalendar recurrence rule (https://datatracker.ietf.org/doc/html/rfc5545)
type CashbunnyRecurrenceRule struct {
	Count    int    `json:"count"`
	Dtstart  string `json:"dtstart"`
	Freq     string `json:"freq"`
	Interval int    `json:"interval"`
	Until    string `json:"until"`
}

// CashbunnyScheduledTransaction Model representing a scheduled transaction
type CashbunnyScheduledTransaction struct {
	Amount                 float64 `json:"amount"`
	AmountDisplay          string  `json:"amount_display"`
	CreatedAt              string  `json:"created_at"`
	Currency               string  `json:"currency"`
	Description            string  `json:"description"`
	DestinationAccountId   uint32  `json:"destination_account_id"`
	DestinationAccountName string  `json:"destination_account_name"`
	Id                     uint32  `json:"id"`

	// RecurrenceRule Model representing iCalendar recurrence rule (https://datatracker.ietf.org/doc/html/rfc5545)
	RecurrenceRule    CashbunnyRecurrenceRule `json:"recurrence_rule"`
	SourceAccountId   uint32                  `json:"source_account_id"`
	SourceAccountName string                  `json:"source_account_name"`
	UpdatedAt         string                  `json:"updated_at"`
}

// CashbunnyTransaction Model representing a single atomic financial transaction between accounts
type CashbunnyTransaction struct {
	Amount                 float64 `json:"amount"`
	AmountDisplay          string  `json:"amount_display"`
	CreatedAt              string  `json:"created_at"`
	Currency               string  `json:"currency"`
	Description            string  `json:"description"`
	DestinationAccountId   uint32  `json:"destination_account_id"`
	DestinationAccountName string  `json:"destination_account_name"`
	Id                     uint32  `json:"id"`

	// ScheduledTransaction Model representing a scheduled transaction
	ScheduledTransaction CashbunnyScheduledTransaction `json:"scheduled_transaction"`
	SourceAccountId      uint32                        `json:"source_account_id"`
	SourceAccountName    string                        `json:"source_account_name"`
	TransactedAt         string                        `json:"transacted_at"`
	UpdatedAt            string                        `json:"updated_at"`
}

// CashbunnyUserPreference Model defining user's Cashbunny preferences such as default currency etc.
type CashbunnyUserPreference struct {
	UserCurrencies []string `json:"user_currencies"`
}

// Error Error code and underlying errors
type Error struct {
	// Message A useful message describing the error
	Message            string            `json:"message"`
	ValidationMessages map[string]string `json:"validation_messages"`
}

// User defines model for user.
type User struct {
	CreatedAt  string `json:"created_at"`
	Id         uint32 `json:"id"`
	LoggedInAt string `json:"logged_in_at"`
	Name       string `json:"name"`
}

// UserSystemPreference Model defining user system preferences such as language, time zone etc.
type UserSystemPreference struct {
	Language *string `json:"language"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody struct {
	// Email must be a valid email
	Email string `json:"email" validate:"email"`

	// Password must be between 8 to 72 characters, and contain a letter, number and a special character
	Password string `json:"password" validate:"_password"`

	// Username must be between 3 to 50 alphanumerical characters
	Username string `json:"username" validate:"alphanum,min=3,max=50"`
}

// CreateJWTTokenJSONBody defines parameters for CreateJWTToken.
type CreateJWTTokenJSONBody struct {
	Password string `json:"password"`

	// Username must be alphanumerical characters, with maximum 50 character count
	Username string `json:"username" validate:"alphanum,max=50"`
}

// RepeatJSONBody defines parameters for Repeat.
type RepeatJSONBody struct {
	// Message Some message you want to be sent back
	Message string `json:"message"`
}

// CreateCashbunnyAccountJSONBody defines parameters for CreateCashbunnyAccount.
type CreateCashbunnyAccountJSONBody struct {
	// Category Category of the account.
	Category    string `json:"category" validate:"required,oneof=assets liabilities revenues expenses"`
	Currency    string `json:"currency" validate:"required,_cashbunny_currency"`
	Description string `json:"description" validate:"required,max=200"`

	// Name Name of the account to save
	Name string `json:"name" validate:"required,max=100"`

	// OrderIndex Order of the record to be placed in
	OrderIndex *uint32 `json:"order_index"`
}

// UpdateCashbunnyAccountJSONBody defines parameters for UpdateCashbunnyAccount.
type UpdateCashbunnyAccountJSONBody struct {
	Description string `json:"description" validate:"required,max=200"`

	// Name Name of the account to save
	Name string `json:"name" validate:"required,max=100"`

	// OrderIndex Order of the record to be placed in
	OrderIndex uint32 `json:"order_index" validate:"required,number"`
}

// GetCashbunnyOverviewParams defines parameters for GetCashbunnyOverview.
type GetCashbunnyOverviewParams struct {
	// From Overview timeframe's start date
	From *int64 `form:"from,omitempty" json:"from,omitempty"`

	// To Overview timeframe's end date
	To *int64 `form:"to,omitempty" json:"to,omitempty"`
}

// CreateCashbunnyTransactionJSONBody defines parameters for CreateCashbunnyTransaction.
type CreateCashbunnyTransactionJSONBody struct {
	Amount      float64 `json:"amount" validate:"required,min=0"`
	Currency    string  `json:"currency" validate:"required,_cashbunny_currency"`
	Description string  `json:"description" validate:"required,max=100"`

	// DestinationAccountId Money from
	DestinationAccountId uint32 `json:"destination_account_id" validate:"required"`

	// SourceAccountId Money from
	SourceAccountId uint32 `json:"source_account_id" validate:"required"`

	// TransactedAt ISO8601 compatible time string for transaction datetime
	TransactedAt string `json:"transacted_at" validate:"required,_iso8601"`
}

// UpdateCashbunnyTransactionJSONBody defines parameters for UpdateCashbunnyTransaction.
type UpdateCashbunnyTransactionJSONBody struct {
	Amount      float64 `json:"amount" validate:"required,min=0"`
	Description string  `json:"description" validate:"required,max=100"`

	// TransactedAt ISO8601 compatible time string for transaction datetime
	TransactedAt string `json:"transacted_at" validate:"required,_iso8601"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// CreateJWTTokenJSONRequestBody defines body for CreateJWTToken for application/json ContentType.
type CreateJWTTokenJSONRequestBody CreateJWTTokenJSONBody

// RepeatJSONRequestBody defines body for Repeat for application/json ContentType.
type RepeatJSONRequestBody RepeatJSONBody

// CreateCashbunnyAccountJSONRequestBody defines body for CreateCashbunnyAccount for application/json ContentType.
type CreateCashbunnyAccountJSONRequestBody CreateCashbunnyAccountJSONBody

// UpdateCashbunnyAccountJSONRequestBody defines body for UpdateCashbunnyAccount for application/json ContentType.
type UpdateCashbunnyAccountJSONRequestBody UpdateCashbunnyAccountJSONBody

// CreateCashbunnyTransactionJSONRequestBody defines body for CreateCashbunnyTransaction for application/json ContentType.
type CreateCashbunnyTransactionJSONRequestBody CreateCashbunnyTransactionJSONBody

// UpdateCashbunnyTransactionJSONRequestBody defines body for UpdateCashbunnyTransaction for application/json ContentType.
type UpdateCashbunnyTransactionJSONRequestBody UpdateCashbunnyTransactionJSONBody

// UpdateUserSystemPreferenceJSONRequestBody defines body for UpdateUserSystemPreference for application/json ContentType.
type UpdateUserSystemPreferenceJSONRequestBody = UserSystemPreference
