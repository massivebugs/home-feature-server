-- name: CreateEntry :execresult
INSERT INTO
  cashbunny_entries (
    user_id,
    account_id,
    transaction_id,
    amount,
    currency,
    transacted_at
  )
VALUES
  (?, ?, ?, ?, ?, ?);

-- name: ListEntriesByTransactionID :many
SELECT
  *
FROM
  cashbunny_entries
WHERE
  user_id = ?
  AND transaction_id = ?;