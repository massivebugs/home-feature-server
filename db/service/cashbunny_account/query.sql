-- name: CreateAccount :execresult
INSERT INTO
  cashbunny_accounts (user_id, name, description, balance, currency, type, order_index)
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