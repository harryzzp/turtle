package db

import (
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLClient struct {
	db *sql.DB
}

var instance *MySQLClient = nil

func Connect() (db *sql.DB, err error) {
	if instance == nil {
		instance = new(MySQLClient)
		var err error
		instance.db, err = sql.Open("mysql",
			"root:123456@/turtle?timeout=90s&collation=utf8mb4_unicode_ci")
		if err != nil {
			return nil, err
		}
	}
	return instance.db, nil
}

func Close() {
	if instance != nil {
		instance.db.Close()
	}
}

