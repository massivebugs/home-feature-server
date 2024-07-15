-- name: GetUser :one
SELECT
  *
FROM
  users
WHERE
  id = ?
LIMIT
  1;

-- name: CreateUser :execresult
INSERT INTO
  users (name)
VALUES
  (?);

-- name: DeleteUser :exec
UPDATE users
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  id = ?;