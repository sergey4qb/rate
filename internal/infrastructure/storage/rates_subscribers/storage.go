package rates_subscribers

import (
	"github.com/sergey4qb/rate/internal/infrastructure/storage/rates_subscribers/db"
	"github.com/sergey4qb/rate/pkg/postgresql_database"
)

type Storage struct {
	Db *db.Db
}

func NewStorage(database *postgresql_database.DB) *Storage {
	return &Storage{Db: db.New(database)}
}
