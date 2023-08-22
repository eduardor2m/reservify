-- name: CreateUser :exec

INSERT INTO "user" ("id", "name", "cpf", "email", "password", "phone", "date_of_birth", "admin", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: Login :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: FindUserByEmail :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: FindUserByID :one

SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;

-- name: ListUsersByName :many

SELECT * FROM "user" WHERE "name" LIKE $1 ORDER BY "id";

-- name: ListAllUsers :many

SELECT * FROM "user" ORDER BY "id";

-- name: UpdateUserByEmail :exec

UPDATE "user" SET "name" = $1, "cpf" = $2, "email" = $3, "password" = $4, "phone" = $5, "date_of_birth" = $6, "admin" = $7, "created_at" = $8, "updated_at" = $9 WHERE "email" = $10;

-- name: DeleteUserByEmail :exec

DELETE FROM "user" WHERE "email" = $1;
