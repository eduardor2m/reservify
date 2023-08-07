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
    "id" UUID PRIMARY KEY,
                                    cod VARCHAR(255) NOT NULL,
                                    number VARCHAR(255) NOT NULL,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS guest (
                                     id UUID PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
                                     cpf VARCHAR(255) NOT NULL,
                                     create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "reservation" (
                                             id UUID PRIMARY KEY,
                                             user_cpf VARCHAR(255) NOT NULL,
                                             room_cod VARCHAR(255) NOT NULL,
                                             guest_cpf VARCHAR(255) NOT NULL,
                                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
