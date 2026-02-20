package mysql_config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func MustCreateConnection(user, pass, net, host string, port int, dbname string) *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = pass
	cfg.Net = net
	cfg.Addr = fmt.Sprintf("%s:%d", host, port)
	cfg.DBName = dbname

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	return db
}
