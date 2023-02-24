package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	url := "root:123456@/book?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
