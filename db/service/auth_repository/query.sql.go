// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package auth_repository

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO
  users (name)
VALUES
  (?)
`

func (q *Queries) CreateUser(ctx context.Context, db DBTX, name string) (sql.Result, error) {
	return db.ExecContext(ctx, createUser, name)
}

const createUserPassword = `-- name: CreateUserPassword :execresult
INSERT INTO
  user_passwords (user_id, password_hash)
VALUES
  (?, ?)
`

type CreateUserPasswordParams struct {
	UserID       uint32
	PasswordHash string
}

func (q *Queries) CreateUserPassword(ctx context.Context, db DBTX, arg CreateUserPasswordParams) (sql.Result, error) {
	return db.ExecContext(ctx, createUserPassword, arg.UserID, arg.PasswordHash)
}

const createUserRefreshToken = `-- name: CreateUserRefreshToken :execresult
INSERT INTO
  user_refresh_tokens (user_id, value, expires_at)
VALUES
  (?, ?, ?)
`

type CreateUserRefreshTokenParams struct {
	UserID    uint32
	Value     string
	ExpiresAt sql.NullTime
}

func (q *Queries) CreateUserRefreshToken(ctx context.Context, db DBTX, arg CreateUserRefreshTokenParams) (sql.Result, error) {
	return db.ExecContext(ctx, createUserRefreshToken, arg.UserID, arg.Value, arg.ExpiresAt)
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users
SET
  deleted_at = CURRENT_TIMESTAMP()
WHERE
  id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, db DBTX, id uint32) error {
	_, err := db.ExecContext(ctx, deleteUser, id)
	return err
}

const deleteUserRefreshToken = `-- name: DeleteUserRefreshToken :exec
DELETE FROM user_refresh_tokens
WHERE
  user_id = ?
  AND value = ?
`

type DeleteUserRefreshTokenParams struct {
	UserID uint32
	Value  string
}

func (q *Queries) DeleteUserRefreshToken(ctx context.Context, db DBTX, arg DeleteUserRefreshTokenParams) error {
	_, err := db.ExecContext(ctx, deleteUserRefreshToken, arg.UserID, arg.Value)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
  id, name, created_at, updated_at, deleted_at
FROM
  users
WHERE
  id = ?
LIMIT
  1
`

func (q *Queries) GetUser(ctx context.Context, db DBTX, id uint32) (*User, error) {
	row := db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT
  id, name, created_at, updated_at, deleted_at
FROM
  users
WHERE
  name = ?
LIMIT
  1
`

func (q *Queries) GetUserByName(ctx context.Context, db DBTX, name string) (*User, error) {
	row := db.QueryRowContext(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUserPasswordByUserID = `-- name: GetUserPasswordByUserID :one
SELECT
  id, user_id, password_hash, created_at, updated_at, deleted_at
FROM
  user_passwords
WHERE
  user_id = ?
LIMIT
  1
`

func (q *Queries) GetUserPasswordByUserID(ctx context.Context, db DBTX, userID uint32) (*UserPassword, error) {
	row := db.QueryRowContext(ctx, getUserPasswordByUserID, userID)
	var i UserPassword
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUserRefreshTokenByValue = `-- name: GetUserRefreshTokenByValue :one
SELECT
  id, user_id, value, expires_at, created_at, updated_at
FROM
  user_refresh_tokens
WHERE
  user_id = ?
  AND value = ?
LIMIT
  1
`

type GetUserRefreshTokenByValueParams struct {
	UserID uint32
	Value  string
}

func (q *Queries) GetUserRefreshTokenByValue(ctx context.Context, db DBTX, arg GetUserRefreshTokenByValueParams) (*UserRefreshToken, error) {
	row := db.QueryRowContext(ctx, getUserRefreshTokenByValue, arg.UserID, arg.Value)
	var i UserRefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Value,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
  name = ?
WHERE
  id = ?
`

type UpdateUserParams struct {
	Name string
	ID   uint32
}

func (q *Queries) UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) error {
	_, err := db.ExecContext(ctx, updateUser, arg.Name, arg.ID)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE user_passwords
SET
  password_hash = ?
WHERE
  id = ?
`

type UpdateUserPasswordParams struct {
	PasswordHash string
	ID           uint32
}

func (q *Queries) UpdateUserPassword(ctx context.Context, db DBTX, arg UpdateUserPasswordParams) error {
	_, err := db.ExecContext(ctx, updateUserPassword, arg.PasswordHash, arg.ID)
	return err
}
