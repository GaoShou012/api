package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func SqlMysql(dns string, poolMin uint64, poolMax uint64) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(int(poolMax))
	db.SetMaxIdleConns(int(poolMin))
	db.SetConnMaxLifetime(time.Minute * 10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
