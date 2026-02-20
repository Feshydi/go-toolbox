package postgresql_config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func MustCreateConnection(host string, port int, user, pass, dbname, sslMode string) *sql.DB {
	cfg := NewConfig()
	cfg.Host = host
	cfg.Port = port
	cfg.User = user
	cfg.Pass = pass
	cfg.DBName = dbname
	cfg.SSLMode = sslMode

	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	return db
}
