package config

import (
	"database/sql"
	"fmt"
	"golang/product_postgress_api/helper"

	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s sslmode=disabled", host, port, password, dbName)
	db, err := sql.Open("postgres", sqlInfo)

	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database")
	return db
}
