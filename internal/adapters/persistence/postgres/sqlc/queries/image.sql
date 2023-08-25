-- name: AddImageToRoomByRoomID :exec

INSERT INTO image (id_room, image_url, thumbnail) VALUES ($1,$2,$3);

-- name: ListImagesByRoomID :many

SELECT id_room, image_url, thumbnail FROM image WHERE id_room = $1;