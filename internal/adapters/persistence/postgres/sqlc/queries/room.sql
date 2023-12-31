-- name: CreateRoom :exec

INSERT INTO room (id, cod, number, vacancies, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: ListAllRooms :many

SELECT id, cod, number, vacancies, price, created_at, updated_at FROM room ORDER BY created_at DESC;

-- name: FindRoomById :one

SELECT id, cod, number, vacancies, price, created_at, updated_at FROM room WHERE id = $1;