package database

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type DataBase struct {
	Conn *pgx.Conn
}

func New() (*DataBase, error) {
	connUrl := createConnectionUrl()
	connConfig, _ := pgx.ParseConfig(connUrl)

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)

	if err != nil {
		return nil, fmt.Errorf("Unable to connection to database: %v\n", err)
	}

	log.Printf("Postgres connected")

	db := DataBase{
		Conn: conn,
	}

	m, err := migrate.New(
		"file://src/database/migrations",
		connUrl,
	)

	if err != nil {
		return nil, fmt.Errorf("Unable to migrate to database: %v\n", err)
	}

	err = m.Up()

	if err != nil {
		log.Printf("Unable to migrate to database: %v\n", err)
	}

	return &db, nil
}

func (db *DataBase) Close() {
	if db.Conn != nil {
		db.Conn.Close(context.Background())
	}
}

func createConnectionUrl() string {
	host := "127.0.0.1"
	port := 5432
	user := "postgres"
	password := "123456"
	database := "app"

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
}
