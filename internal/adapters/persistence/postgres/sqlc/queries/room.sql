-- name: CreateRoom :exec

INSERT INTO room (id, name, description, cod, number, vacancies, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: ListAllRooms :many

SELECT id, name, description, cod, number, vacancies, price, created_at, updated_at FROM room ORDER BY created_at DESC;

-- name: FindRoomById :one

SELECT id, name, description, cod, number, vacancies, price, created_at, updated_at FROM room WHERE id = $1;

-- name: FindRoomByCod :one

SELECT id, name, description, cod, number, vacancies, price, created_at, updated_at FROM room WHERE cod = $1;

-- name: DeleteRoomById :exec

DELETE FROM room WHERE id = $1;