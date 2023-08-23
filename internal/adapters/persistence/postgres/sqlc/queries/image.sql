-- name: AddImageToRoomByRoomID :exec

INSERT INTO image (id_room, image_url) VALUES ($1,$2);

-- name: ListImagesByRoomID :many

SELECT id_room, image_url FROM image WHERE id_room = $1;