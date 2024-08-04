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
ORDER BY
  order_index;

-- name: ListAccountsAndCategories :many
SELECT
  sqlc.embed(cashbunny_accounts),
  sqlc.embed(cashbunny_categories)
FROM
  cashbunny_accounts
  LEFT JOIN cashbunny_categories ON cashbunny_categories.id = cashbunny_accounts.category_id
WHERE
  cashbunny_accounts.user_id = ?
ORDER BY
  cashbunny_accounts.order_index;

-- name: IncrementIndex :exec
UPDATE cashbunny_accounts
SET
  order_index = order_index + 1
WHERE
  user_id = ?
  AND order_index >= ?;

-- name: GetCategoryByID :one
SELECT
  *
FROM
  cashbunny_categories
WHERE
  user_id = ?
  AND id = ?
LIMIT
  1;

-- name: CreateCategory :execresult
INSERT INTO
  cashbunny_categories (user_id, name, description)
VALUES
  (?, ?, ?);

-- name: ListCategoriesByUserID :many
SELECT
  *
FROM
  cashbunny_categories
WHERE
  user_id = ?;