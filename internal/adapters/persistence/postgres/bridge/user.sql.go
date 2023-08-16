// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
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
	DateOfBirth time.Time
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

const deleteByEmail = `-- name: DeleteByEmail :exec

DELETE FROM "user" WHERE "email" = $1
`

func (q *Queries) DeleteByEmail(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteByEmail, email)
	return err
}

const findByEmail = `-- name: FindByEmail :one

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "email" = $1 LIMIT 1
`

func (q *Queries) FindByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
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

const findByID = `-- name: FindByID :one

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "id" = $1 LIMIT 1
`

func (q *Queries) FindByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, findByID, id)
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

const listAll = `-- name: ListAll :many

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" ORDER BY "id"
`

func (q *Queries) ListAll(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listAll)
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

const listByName = `-- name: ListByName :many

SELECT id, name, cpf, email, phone, date_of_birth, password, admin, created_at, updated_at FROM "user" WHERE "name" LIKE $1 ORDER BY "id"
`

func (q *Queries) ListByName(ctx context.Context, name string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listByName, name)
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

const updateByEmail = `-- name: UpdateByEmail :exec

UPDATE "user" SET "name" = $1, "cpf" = $2, "email" = $3, "password" = $4, "phone" = $5, "date_of_birth" = $6, "admin" = $7, "created_at" = $8, "updated_at" = $9 WHERE "email" = $10
`

type UpdateByEmailParams struct {
	Name        string
	Cpf         string
	Email       string
	Password    string
	Phone       string
	DateOfBirth time.Time
	Admin       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email_2     string
}

func (q *Queries) UpdateByEmail(ctx context.Context, arg UpdateByEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateByEmail,
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
