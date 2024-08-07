-- name: CreateAccount :execresult
INSERT INTO
  cashbunny_accounts (
    user_id,
    category_id,
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

-- name: ListAccountsAndCategories :many
SELECT
  sqlc.embed(cashbunny_accounts),
  sqlc.embed(cashbunny_account_categories)
FROM
  cashbunny_accounts
  LEFT JOIN cashbunny_account_categories ON cashbunny_account_categories.id = cashbunny_accounts.category_id
WHERE
  cashbunny_accounts.user_id = ?
  AND cashbunny_accounts.deleted_at IS NULL
ORDER BY
  cashbunny_accounts.order_index;

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

-- name: GetAccountCategoryByID :one
SELECT
  *
FROM
  cashbunny_account_categories
WHERE
  user_id = ?
  AND deleted_at IS NULL
  AND id = ?
LIMIT
  1;

-- name: GetAccountCategoryByName :one
SELECT
  *
FROM
  cashbunny_account_categories
WHERE
  user_id = ?
  AND deleted_at IS NULL
  AND name = ?
LIMIT
  1;

-- name: CreateAccountCategory :execresult
INSERT INTO
  cashbunny_account_categories (user_id, name, description)
VALUES
  (?, ?, ?);

-- name: ListAccountCategoriesByUserID :many
SELECT
  *
FROM
  cashbunny_account_categories
WHERE
  user_id = ?
  AND deleted_at IS NULL;