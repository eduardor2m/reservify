CREATE TABLE IF NOT EXISTS "USER" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50) NOT NULL,
    date_of_birth DATE NOT NULL,
    password VARCHAR(255) NOT NULL,
    admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS room (
    id UUID PRIMARY KEY,
                                    cod VARCHAR(255) NOT NULL UNIQUE,
                                    number VARCHAR(255) NOT NULL,
    vacancies INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "reservation" (
                                             id UUID PRIMARY KEY,
                                             id_user UUID NOT NULL,
                                             id_room UUID NOT NULL,
                                                check_in DATE NOT NULL,
                                                check_out DATE NOT NULL,
                                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                FOREIGN KEY (id_user) REFERENCES user(id),
                                                FOREIGN KEY (id_room) REFERENCES room(id)
);
