-- name: CreateTransaction :execresult
INSERT INTO
  cashbunny_transactions (user_id, description, transacted_at)
VALUES
  (?, ?, ?);

-- name: ListTransactionsByUserID :many
SELECT
  *
FROM
  cashbunny_transactions
WHERE
  user_id = ?;