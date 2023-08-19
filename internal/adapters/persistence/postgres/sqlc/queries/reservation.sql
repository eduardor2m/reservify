-- name: CreateReservation :exec

INSERT INTO "reservation" (id, id_user, id_room, check_in, check_out, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: ListAllReservations :many

SELECT id, id_user, id_room, check_in, check_out, created_at, updated_at FROM "reservation" ORDER BY created_at DESC;