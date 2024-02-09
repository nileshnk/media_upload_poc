// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO public.user (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email, password, created_at, updated_at
`

type CreateUserParams struct {
	ID       pgtype.UUID
	Name     pgtype.Text
	Email    pgtype.Text
	Password pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM public.user WHERE id = $1 RETURNING id, name, email, password, created_at, updated_at
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :many
SELECT id, name, email, password, created_at, updated_at FROM public.user
`

func (q *Queries) GetUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, created_at, updated_at FROM public.user WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, password, created_at, updated_at FROM public.user WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserCount = `-- name: GetUserCount :one
SELECT COUNT(*) FROM public.user
`

func (q *Queries) GetUserCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getUserCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE public.user SET name = $2, email = $3, password = $4 WHERE id = $1 RETURNING id, name, email, password, created_at, updated_at
`

type UpdateUserParams struct {
	ID       pgtype.UUID
	Name     pgtype.Text
	Email    pgtype.Text
	Password pgtype.Text
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserName = `-- name: UpdateUserName :one
UPDATE public.user SET name = $2 WHERE id = $1 RETURNING id, name, email, password, created_at, updated_at
`

type UpdateUserNameParams struct {
	ID   pgtype.UUID
	Name pgtype.Text
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserName, arg.ID, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
UPDATE public.user SET password = $2 WHERE id = $1 RETURNING id, name, email, password, created_at, updated_at
`

type UpdateUserPasswordParams struct {
	ID       pgtype.UUID
	Password pgtype.Text
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserPassword, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
