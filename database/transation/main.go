package main

import (
	"database/sql"
	"fmt"
	"log"
)

func transfer(db *sql.DB, a int, b int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("start failed", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var balance float64
	err = tx.QueryRow(
		"SELECT balance FROM balances WHERE id = ?",
		a,
	).Scan(&balance)

	if balance < amount {
		return fmt.Errorf("not enough balance")
	}

	_, err = tx.Exec(
		"update accounts set balance = balance - ? where id = ?",
		amount, a)
	if err != nil {
		return fmt.Errorf("update  minus failed", err)
	}

	_, err = tx.Exec(
		"update accounts set balance = balance + ? where id = ?",
		amount, b)
	if err != nil {
		return fmt.Errorf("update add failed", err)
	}

	_, err = tx.Exec(
		"insert into transactions (from_account_id,to_account_id, amount) values (?, ?, ?)",
		a, b, amount)
	if err != nil {
		return fmt.Errorf("insert log failed", err)
	}
	return nil
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("connect databse failed", err)
	}
	defer db.Close()

	err = transfer(db, 1, 2, 100)
	if err != nil {
		log.Fatal("transfer failed", err)
	}
	fmt.Println("transfer success")
}
