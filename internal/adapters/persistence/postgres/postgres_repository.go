package postgres

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://tools/database/migrations",
		"postgres", driver)
	
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm *DatabaseConnectorManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()

	if err != nil {
		panic(err)
	}
}
