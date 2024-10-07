-- name: CreateUserSystemPreference :execresult
INSERT INTO
  user_system_preferences (user_id, `language`)
VALUES
  (?, ?);

-- name: GetUserSystemPreferenceExists :one
SELECT
  EXISTS (
    SELECT
      *
    FROM
      user_system_preferences
    WHERE
      user_id = ?
      AND deleted_at IS NULL
    LIMIT
      1
  );

-- name: GetUserSystemPreference :one
SELECT
  *
FROM
  user_system_preferences
WHERE
  user_id = ?
  AND deleted_at IS NULL
LIMIT
  1;

-- name: UpdateUserSystemPreference :exec
UPDATE user_system_preferences
SET
  `language` = ?
WHERE
  user_id = ?;