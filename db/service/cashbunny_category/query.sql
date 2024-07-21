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