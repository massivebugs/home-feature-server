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
LIMIT
  1;

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