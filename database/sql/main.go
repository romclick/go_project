package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("connect database failed", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("ping database failed", err)
	}
	sqlStmt := "INSERT INTO students (name, age , grade) values ('张三',20,'三年级')"
	result, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("insert failed", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("get last insert id failed", err)
	}
	fmt.Printf("insert success, id=%d", id)

}
