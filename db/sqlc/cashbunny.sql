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


-- name: GetUserPreferenceExistsByUserID :one
SELECT EXISTS (
  SELECT
    *
  FROM
    cashbunny_user_preferences
  WHERE
    user_id = ?
    AND deleted_at IS NULL
  LIMIT
    1
);

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
    currency,
    order_index
  )
VALUES
  (
    sqlc.arg(user_id),
    ?,
    ?,
    ?,
    ?,
    COALESCE(
      sqlc.narg(order_index),
      (
        SELECT
          COALESCE(MAX(t.order_index) + 1, 0)
        FROM
          cashbunny_accounts AS t
        WHERE
          t.user_id = sqlc.arg(user_id)
      )
    )
  );

-- name: UpdateCashbunnyAccount :exec
UPDATE cashbunny_accounts
SET
  name = ?,
  description = ?,
  order_index = ?
WHERE
  user_id = ?
  AND id = ?;

-- name: ListAccountsAndAmountByCategory :many
SELECT
  sqlc.embed(a),
  CAST(
    COALESCE(
      SUM(
        CASE
          WHEN a.category IN ('assets', 'expenses')
          THEN IF(tr.src_account_id = a.id, -tr.amount, tr.amount)
          ELSE IF(tr.src_account_id = a.id, tr.amount, -tr.amount)
        END
      ), 0
    ) AS DECIMAL(19, 4)
  ) AS amount
FROM
  cashbunny_accounts a
  LEFT JOIN cashbunny_transactions tr
    ON (tr.src_account_id = a.id OR tr.dest_account_id = a.id)
    AND tr.deleted_at IS NULL
WHERE
  a.user_id = ?
  AND a.category = ?
  AND a.deleted_at IS NULL
GROUP BY
  a.id
ORDER BY
  a.order_index;

-- name: ListAccountsAndAmountBetweenDates :many
SELECT
  sqlc.embed(a),
  CAST(
    COALESCE(
      SUM(
        CASE
          WHEN a.category IN ('assets', 'expenses')
          THEN IF(tr.src_account_id = a.id, -tr.amount, tr.amount)
          ELSE IF(tr.src_account_id = a.id, tr.amount, -tr.amount)
        END
      ), 0
    ) AS DECIMAL(19, 4)
  ) AS amount
FROM
  cashbunny_accounts a
  LEFT JOIN cashbunny_transactions tr
    ON (tr.src_account_id = a.id OR tr.dest_account_id = a.id)
    AND tr.transacted_at BETWEEN ? AND ? 
    AND tr.deleted_at IS NULL
WHERE
  a.user_id = ?
  AND a.deleted_at IS NULL
GROUP BY
  a.id
ORDER BY
  a.order_index;

-- name: ListAccountsByIDs :many
SELECT
  *
FROM
  cashbunny_accounts
WHERE
  user_id = ?
  AND deleted_at IS NULL
  AND id IN (sqlc.slice('IDs'))
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

-- name: DeleteAccount :exec
UPDATE cashbunny_accounts
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  user_id = ?
  AND id = ?;

-- name: IncrementAccountIndices :exec
UPDATE cashbunny_accounts
SET
  order_index = order_index + 1
WHERE
  user_id = ?
  AND order_index >= ?;

-- name: CreateTransactionCategory :execresult
INSERT INTO
  cashbunny_transaction_categories (
    user_id,
    name
  )
VALUES
  (?, ?);

-- name: ListTransactionCategories :many
SELECT
  *
FROM
  cashbunny_transaction_categories
WHERE
  user_id = ?
  AND deleted_at IS NULL
ORDER BY
  id ASC;

-- name: CreateTransactionCategoryAllocation :execresult
INSERT INTO
  cashbunny_transaction_category_allocations (
    user_id,
    category_id,
    amount,
    currency
  )
VALUES
  (?, ?, ?, ?);

-- name: ListTransactionCategoryAllocations :many
SELECT
  *
FROM
  cashbunny_transaction_category_allocations
WHERE
  user_id = ?
  AND deleted_at IS NULL;

-- name: CreateTransaction :execresult
INSERT INTO
  cashbunny_transactions (
    user_id,
    scheduled_transaction_id,
    category_id,
    src_account_id,
    dest_account_id,
    description,
    amount,
    currency,
    transacted_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCashbunnyTransaction :exec
UPDATE cashbunny_transactions
SET
  description = ?,
  amount = ?,
  transacted_at = ?
WHERE
  user_id = ?
  AND id = ?;

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
  transacted_at ASC;

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
  AND src_account_id = sqlc.arg(account_id)
  OR dest_account_id = sqlc.arg(account_id);

-- name: CreateScheduledTransaction :execresult
INSERT INTO
  cashbunny_scheduled_transactions (
    user_id,
    category_id,
    src_account_id,
    dest_account_id,
    description,
    amount,
    currency
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?);

-- name: ListScheduledTransactionsWithAllRelations :many
SELECT
  sqlc.embed(cashbunny_scheduled_transactions),
  sqlc.embed(cashbunny_recurrence_rules),
  sqlc.embed(categories),
  sqlc.embed(source_account),
  sqlc.embed(destination_account)
FROM
  cashbunny_scheduled_transactions
  LEFT JOIN cashbunny_scheduled_transactions_recurrence_rules as relationship ON relationship.scheduled_transaction_id = cashbunny_scheduled_transactions.id
  LEFT JOIN cashbunny_recurrence_rules ON cashbunny_recurrence_rules.id = relationship.recurrence_rule_id
  LEFT JOIN cashbunny_transaction_categories AS categories ON categories.id = category_id
  LEFT JOIN cashbunny_accounts AS source_account ON source_account.id = src_account_id
  LEFT JOIN cashbunny_accounts AS destination_account ON destination_account.id = dest_account_id
WHERE
  cashbunny_scheduled_transactions.user_id = ?
  AND cashbunny_scheduled_transactions.deleted_at IS NULL
ORDER BY
  cashbunny_scheduled_transactions.created_at ASC;

-- name: CreateRecurrenceRule :execresult
INSERT INTO
  cashbunny_recurrence_rules (freq, dtstart, count, `interval`, `until`)
VALUES
  (?, ?, ?, ?, ?);

-- name: CreateScheduledTransactionRecurrenceRuleRelationship :execresult
INSERT INTO
  cashbunny_scheduled_transactions_recurrence_rules (scheduled_transaction_id, recurrence_rule_id)
VALUES
  (?, ?);