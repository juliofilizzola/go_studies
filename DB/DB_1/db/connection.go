package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connection() {
	url := "root:123456@/book?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", url)

	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection...")
}
