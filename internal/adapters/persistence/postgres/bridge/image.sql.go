// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: image.sql

package bridge

import (
	"context"

	"github.com/google/uuid"
)

const addImageToRoomByRoomID = `-- name: AddImageToRoomByRoomID :exec

INSERT INTO image (id_room, image_url) VALUES ($1,$2)
`

type AddImageToRoomByRoomIDParams struct {
	IDRoom   uuid.UUID
	ImageUrl string
}

func (q *Queries) AddImageToRoomByRoomID(ctx context.Context, arg AddImageToRoomByRoomIDParams) error {
	_, err := q.db.ExecContext(ctx, addImageToRoomByRoomID, arg.IDRoom, arg.ImageUrl)
	return err
}

const listImagesByRoomID = `-- name: ListImagesByRoomID :many

SELECT id_room, image_url FROM image WHERE id_room = $1
`

func (q *Queries) ListImagesByRoomID(ctx context.Context, idRoom uuid.UUID) ([]Image, error) {
	rows, err := q.db.QueryContext(ctx, listImagesByRoomID, idRoom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.IDRoom, &i.ImageUrl); err != nil {
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
