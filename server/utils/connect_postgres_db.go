package utils

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/zcubbs/pulse/server/models"
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

	//entries := []models.PipelineStatusEntry{
	//	{Status: "success", ProjectName: "test", ProjectId: "123456", Origin: "gitlab", OriginUrl: "https://toto.test", EventDate: time.Time{}, Message: "test", Group: "test_group"},
	//	{Status: "failed", ProjectName: "test2", ProjectId: "123457", Origin: "gitlab", OriginUrl: "https://toto.test", EventDate: time.Time{}, Message: "test", Group: "test_group"},
	//	{Status: "running", ProjectName: "test2", ProjectId: "123458", Origin: "gitlab", OriginUrl: "https://toto.test", EventDate: time.Time{}, Message: "test", Group: "test_group"},
	//	{Status: "?", ProjectName: "test2", ProjectId: "123459", Origin: "gitlab", OriginUrl: "https://toto.test", EventDate: time.Time{}, Message: "test", Group: "test_group"},
	//	{Status: "?", ProjectName: "test2", ProjectId: "1234510", Origin: "gitlab", OriginUrl: "https://toto.test", EventDate: time.Time{}, Message: "test", Group: "test_group"},
	//}
	//_, err = pgDatabase.NewInsert().Model(&entries).Exec(ctx)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("Successfully connected to PostgresDB")
}

func getPostgresDBURL() string {
	postgresDBURL := os.Getenv("POSTGRES_DB_URL")
	if postgresDBURL == "" {
		postgresDBURL = "postgres://postgres:postgres@localhost:5432/zrocket?sslmode=disable&application_name=zrocket"
	}

	return postgresDBURL
}

func GetPgDatabase() *bun.DB {
	return pgDatabase
}
