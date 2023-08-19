package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectorManager)(nil)

type DatabaseConnectorManager struct{}

func (dcm *DatabaseConnectorManager) getConnection() (*sqlx.DB, error) {
	uri := "postgres://root:root@localhost:5432/reservify?sslmode=disable"

	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50) NOT NULL,
    date_of_birth VARCHAR(10) NOT NULL,
    password VARCHAR(255) NOT NULL,
    admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS room (
    id UUID PRIMARY KEY,
                                    cod VARCHAR(255) NOT NULL UNIQUE,
                                    number INTEGER NOT NULL,
    vacancies INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
                                    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		    		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_, err = db.Exec(`
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
	`)

		if err != nil {
			log.Fatal(err)
		}
	}()

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm *DatabaseConnectorManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()

	if err != nil {
		panic(err)
	}
}
