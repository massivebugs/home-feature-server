-- name: GetUserPasswordByUserID :one
SELECT
  *
FROM
  user_passwords
WHERE
  user_id = ?
LIMIT
  1;

-- name: CreateUserPassword :execresult
INSERT INTO
  user_passwords (user_id, password_hash)
VALUES
  (?, ?);

-- name: UpdateUserPassword :exec
UPDATE user_passwords
SET
  password_hash = ?
WHERE
  id = ?;