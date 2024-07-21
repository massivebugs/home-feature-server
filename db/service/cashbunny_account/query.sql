-- name: CreateAccount :execresult
INSERT INTO
  cashbunny_accounts (
    user_id,
    name,
    description,
    balance,
    currency,
    type,
    order_index
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?);

-- name: ListAccounts :many
SELECT
  *
FROM
  cashbunny_accounts
WHERE
  user_id = ?
ORDER BY
  order_index;

-- name: IncrementIndex :exec
UPDATE cashbunny_accounts
SET
  order_index = order_index + 1
WHERE
  user_id = ?
  AND order_index >= ?;
