-- name: ListUserCurrencies :many
SELECT
  *
FROM
  cashbunny_user_currencies
WHERE
  user_id = ?
ORDER BY
  currency_code;

-- name: CreateUserCurrency :execresult
INSERT INTO
  cashbunny_user_currencies (user_id, currency_code)
VALUES
  (?, ?);

-- name: GetUserPreferenceByUserID :one
SELECT
  *
FROM
  cashbunny_user_preferences
WHERE
  user_id = ?
  AND deleted_at IS NULL
LIMIT
  1;

-- name: CreateUserPreferences :execresult
INSERT INTO
  cashbunny_user_preferences (user_id)
VALUES
  (?);

-- name: CreateAccount :execresult
INSERT INTO
  cashbunny_accounts (
    user_id,
    category,
    name,
    description,
    balance,
    currency,
    type,
    order_index
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListAccounts :many
SELECT
  *
FROM
  cashbunny_accounts
WHERE
  user_id = ?
  AND deleted_at IS NULL
ORDER BY
  order_index;

-- name: ListAccountsByIDs :many
SELECT
  *
FROM
  cashbunny_accounts
WHERE
  user_id = ?
  AND deleted_at IS NULL
  AND id IN (sqlc.slice ('IDs'))
ORDER BY
  order_index;

-- name: GetAccountByID :one
SELECT
  *
FROM
  cashbunny_accounts
WHERE
  user_id = ?
  AND id = ?
  AND deleted_at IS NULL
LIMIT
  1;

-- name: UpdateAccountBalance :exec
UPDATE cashbunny_accounts
SET
  balance = ?
WHERE
  user_id = ?
  AND id = ?
  AND deleted_at IS NULL;

-- name: DeleteAccount :exec
UPDATE cashbunny_accounts
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  user_id = ?
  AND id = ?;

-- name: IncrementIndex :exec
UPDATE cashbunny_accounts
SET
  order_index = order_index + 1
WHERE
  user_id = ?
  AND order_index >= ?;

-- name: CreateTransaction :execresult
INSERT INTO
  cashbunny_transactions (
    user_id,
    src_account_id,
    dest_account_id,
    description,
    amount,
    currency,
    transacted_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?);

-- name: ListTransactions :many
SELECT
  *
FROM
  cashbunny_transactions
WHERE
  user_id = ?
  AND deleted_at IS NULL
ORDER BY
  transacted_at;

-- name: GetTransactionByID :one
SELECT
  *
FROM
  cashbunny_transactions
WHERE
  user_id = ?
  AND id = ?
  AND deleted_at IS NULL
LIMIT
  1;

-- name: ListTransactionsBetweenDates :many
SELECT
  *
FROM
  cashbunny_transactions
WHERE
  user_id = ?
  AND deleted_at IS NULL
  AND transacted_at BETWEEN ? AND ?
ORDER BY
  transacted_at;

-- name: DeleteTransaction :exec
UPDATE cashbunny_transactions
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  user_id = ?
  AND id = ?;

-- name: DeleteTransactionsByAccountID :exec
UPDATE cashbunny_transactions
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  user_id = ?
  AND src_account_id = sqlc.arg (account_id)
  OR dest_account_id = sqlc.arg (account_id);