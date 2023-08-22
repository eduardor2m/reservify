// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package bridge

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :exec

INSERT INTO "user" ("id", "name", "cpf", "email", "password", "phone", "date_of_birth", "admin", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type CreateUserParams struct {
	ID          uuid.UUID
	Name        string
	Cpf         string
	Email       string
	Password    string
	Phone       string
	DateOfBirth string
	Admin       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Cpf,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.DateOfBirth,
		arg.Admin,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec

DELETE FROM "user" WHERE "email" = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteUserByEmail, email)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "email" = $1 LIMIT 1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Email,
		&i.Phone,
		&i.DateOfBirth,
		&i.Password,
		&i.Admin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "id" = $1 LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Email,
		&i.Phone,
		&i.DateOfBirth,
		&i.Password,
		&i.Admin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllUsers = `-- name: ListAllUsers :many

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" ORDER BY "id"
`

func (q *Queries) ListAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listAllUsers)
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
			&i.Cpf,
			&i.Email,
			&i.Phone,
			&i.DateOfBirth,
			&i.Password,
			&i.Admin,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByName = `-- name: ListUsersByName :many

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "name" LIKE $1 ORDER BY "id"
`

func (q *Queries) ListUsersByName(ctx context.Context, name string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByName, name)
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
			&i.Cpf,
			&i.Email,
			&i.Phone,
			&i.DateOfBirth,
			&i.Password,
			&i.Admin,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const login = `-- name: Login :one

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "email" = $1 LIMIT 1
`

func (q *Queries) Login(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, login, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpf,
		&i.Email,
		&i.Phone,
		&i.DateOfBirth,
		&i.Password,
		&i.Admin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserByEmail = `-- name: UpdateUserByEmail :exec

UPDATE "user" SET "name" = $1, "cpf" = $2, "email" = $3, "password" = $4, "phone" = $5, "date_of_birth" = $6, "admin" = $7, "created_at" = $8, "updated_at" = $9 WHERE "email" = $10
`

type UpdateUserByEmailParams struct {
	Name        string
	Cpf         string
	Email       string
	Password    string
	Phone       string
	DateOfBirth string
	Admin       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email_2     string
}

func (q *Queries) UpdateUserByEmail(ctx context.Context, arg UpdateUserByEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserByEmail,
		arg.Name,
		arg.Cpf,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.DateOfBirth,
		arg.Admin,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Email_2,
	)
	return err
}
