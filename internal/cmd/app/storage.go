package app

import (
	"github.com/sergey4qb/rate/pkg/postgresql_database"
	"log"
	"os"
)

type storage struct {
	db *postgresql_database.DB
}

func createStorage() *storage {
	dbConfig := postgresql_database.Config{
		User:                os.Getenv("POSTGRES_USER"),
		Password:            os.Getenv("POSTGRES_PASSWORD"),
		Host:                os.Getenv("POSTGRES_HOST"),
		Port:                os.Getenv("POSTGRES_PORT"),
		DBName:              os.Getenv("POSTGRES_DB"),
		Network:             "tcp",
		SSLMode:             "disable",
		MaxIdleCons:         0,
		RequestTimeout:      0,
		MaxIdleConnDuration: 30,
		DriverName:          "postgres",
	}
	db, err := postgresql_database.New(&dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	return &storage{
		db: db,
	}
}
