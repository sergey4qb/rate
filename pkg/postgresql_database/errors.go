package postgresql_database

import "errors"

var (
	ErrConnect = errors.New(
		"postgres query: error connect to database",
	)
)
