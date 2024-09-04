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

-- name: CreateUserRefreshToken :execresult
INSERT INTO
  user_refresh_tokens (user_id, value, expires_at)
VALUES
  (?, ?, ?);

-- name: GetUserRefreshTokenByValue :one
SELECT
  *
FROM
  user_refresh_tokens
WHERE
  user_id = ?
  AND value = ?
LIMIT
  1;

-- name: UpdateUserRefreshTokenExpiresAt :exec
UPDATE user_refresh_tokens
SET
  expires_at = ?
WHERE
  id = ?;

-- name: DeleteUserRefreshToken :exec
DELETE FROM user_refresh_tokens
WHERE
  user_id = ?
  AND value = ?;