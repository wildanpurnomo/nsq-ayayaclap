package maindb

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func InitPostgre() {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN"))
	if err != nil {
		log.Panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Panic(err)
	}

	PostgresDB = db
}

func MigratePostgre() {
	if err := PostgresDB.Ping(); err != nil {
		log.Panic(err)
	}

	m, err := migrate.New(
		"file://db/migrations/postgre",
		os.Getenv("POSTGRES_CONN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
