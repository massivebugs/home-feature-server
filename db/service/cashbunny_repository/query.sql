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