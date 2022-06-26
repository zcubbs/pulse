package utils

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/zcubbs/pulse/pipelines/models"
	"log"
	"os"
)

var pgDatabase *bun.DB
var ctx = context.Background()

func ConnectToPostgresDB() {
	postgresDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(getPostgresDBURL())))
	pgDatabase = bun.NewDB(postgresDB, pgdialect.New())
	pgDatabase.RegisterModel((*models.PipelineStatusEntry)(nil))

	_, err := pgDatabase.NewCreateTable().
		Model((*models.PipelineStatusEntry)(nil)).
		Exec(ctx)
	if err != nil {
		log.Println(err)
	}
}

func getPostgresDBURL() string {
	postgresDBURL := os.Getenv("POSTGRES_DB_URL")
	if postgresDBURL == "" {
		postgresDBURL = "postgres://postgres:postgres@localhost:5432/pulse?sslmode=disable&application_name=zrocket"
	}

	return postgresDBURL
}

func GetPgDatabase() *bun.DB {
	return pgDatabase
}
