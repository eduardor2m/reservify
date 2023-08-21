CREATE TABLE IF NOT EXISTS "reservation" (
    id UUID PRIMARY KEY,
    id_user UUID NOT NULL,
    id_room UUID NOT NULL,
    check_in VARCHAR(10) NOT NULL,
    check_out VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (id_user) REFERENCES "user"(id),
    FOREIGN KEY (id_room) REFERENCES room(id)
);