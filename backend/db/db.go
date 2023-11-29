package db

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func Init() {
	db = connect()
}

func GetDB() *bun.DB {
	return db
}

func connect() *bun.DB {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	dbConnector := pgdriver.NewConnector(
		pgdriver.WithAddr("postgres:5432"),
		pgdriver.WithDatabase("aikido-db"),
		pgdriver.WithUser(dbUser),
		pgdriver.WithPassword(dbPassword),
		pgdriver.WithInsecure(true),
	)
	if os.Getenv("ENV") == "dev" {
		dbConnector.Config().Addr = "localhost:5432"
	}

	return bun.NewDB(sql.OpenDB(dbConnector), pgdialect.New())
}
