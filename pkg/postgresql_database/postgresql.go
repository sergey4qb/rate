package postgresql_database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	conf Config
	conn *sqlx.DB
}

func New(conf *Config) (*DB, error) {
	psqlConn := &DB{
		conf: *conf,
	}
	return psqlConn, nil
}

func (psql *DB) Close() {
	if psql.conn != nil {
		psql.conn.Close()
		log.Println("Disconnected from PostgresDB")
	}
}

func (psql *DB) GetClient() *sqlx.DB {
	if psql.conn == nil {
		psql.connectToPostgres()
	}
	return psql.conn
}

func (psql *DB) connectToPostgres() {
	connSettings := psql.conf.ConnectionString()
	db := sqlx.MustConnect(psql.conf.DriverName, connSettings)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to PostgresDB:", err)
	}
	psql.conn = db
	log.Println("Connected to PostgreSQL")
}

func (psql *DB) Reconnect() {
	psql.Close()
	psql.connectToPostgres()
}
