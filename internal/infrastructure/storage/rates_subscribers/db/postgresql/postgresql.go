package postgresql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sergey4qb/rate/pkg/postgresql_database"
)

type PostgreQuery struct {
	conn *sqlx.DB
}

func New(db *postgresql_database.DB) *PostgreQuery {
	q := &PostgreQuery{
		conn: db.GetClient(),
	}
	return q
}

func (r *PostgreQuery) Close() error {
	return r.conn.Close()
}
