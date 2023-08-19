-- name: CreateReservation :exec

INSERT INTO "reservation" (id, id_user, id_room, check_in, check_out, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: ListAllReservations :many

SELECT id, id_user, id_room, check_in, check_out, created_at, updated_at FROM "reservation" ORDER BY created_at DESC;

-- name: GetReservationByID :one

SELECT id, id_user, id_room, check_in, check_out, created_at, updated_at FROM "reservation" WHERE id = $1 LIMIT 1;

-- name: GetReservationByIDRoom :many

SELECT id, id_user, id_room, check_in, check_out, created_at, updated_at FROM "reservation" WHERE id_room = $1 ORDER BY created_at DESC;

-- name: GetReservationByIDUser :many

SELECT id, id_user, id_room, check_in, check_out, created_at, updated_at FROM "reservation" WHERE id_user = $1 ORDER BY created_at DESC;

-- name: DeleteReservation :exec

DELETE FROM "reservation" WHERE id = $1;