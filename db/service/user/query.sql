-- name: GetUser :one
SELECT
  *
FROM
  users
WHERE
  id = ?
LIMIT
  1;

-- name: GetUserByName :one
SELECT
  *
FROM
  users
WHERE
  name = ?
LIMIT
  1;

-- name: CreateUser :execresult
INSERT INTO
  users (name)
VALUES
  (?);

-- name: UpdateUser :exec
UPDATE users
SET
  name = ?
WHERE
  id = ?;

-- name: DeleteUser :exec
UPDATE users
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  id = ?;