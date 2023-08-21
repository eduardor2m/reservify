CREATE TABLE IF NOT EXISTS "image" (
    id_room UUID NOT NULL,
    image_url TEXT PRIMARY KEY,
    FOREIGN KEY (id_room) REFERENCES room(id)
);