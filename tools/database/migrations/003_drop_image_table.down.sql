CREATE TABLE IF NOT EXISTS "image" (
    id_room UUID PRIMARY KEY,
    image_url TEXT NOT NULL,
    FOREIGN KEY (id_room) REFERENCES room(id)
);