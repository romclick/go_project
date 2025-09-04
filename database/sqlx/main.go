package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type employee struct {
	ID         int     `db: id`
	name       string  `db: name`
	department string  `db: department`
	salary     float32 `db: salary`
}

func initDB(dsn string) (*sql.DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("db connect failed", err)
	}
	if err := db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("db ping failed", err)
	}
	return db, nil
}

func querytechEmploees(db *sql.DB) ([]employee, error) {
	sql := "select id,name,department,salary from employees where department = ?"

	var techEmployees []employee

	err := db.SelectContext(
		context.Background(),
		&techEmployees,
		sql,
		"技术部",
	)
	if err != nil {
		return nil, fmt.Errorf("query tech employees failed", err)
	}
	return techEmployees, nil

}

func queryhighSalary(db *sql.DB) (employee, error) {
	sql := "select id,name,department,salary from employees order by salary desc limit 1"
	var topEmployee employee
	err := db.GetContext(
		context.Background(),
		&topEmployee,
		sql,
	)
	if err != nil {
		return employee{}, fmt.Errorf("query high salary failed", err)
	}
	return topEmployee, nil
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err := initDB(dsn)
	if err != nil {
		log.Fatalf("init db failed", err)
	}
	defer db.Close()

	techEmploees, err := querytechEmploees(db)
	if err != nil {
		log.Fatalf("query tech employees failed", err)
	} else {
		fmt.Printf("技术部共%v\n人", len(techEmploees))
		for _, employee := range techEmploees {
			fmt.Printf("ID: %d,姓名：%s ,部门：%s,工资：%.2f\n", employee.id, employee.name, employee.department, employee.salary)
		}
	}
	topEmployee, err := queryhighSalary(db)
	if err != nil {
		log.Fatalf("query high salary failed", err)
	} else {
		fmt.Printf("ID: %d, 姓名： %s, 部门： %s,工资： %.2f\n", topEmployee.id, topEmployee.name, topEmployee.department, topEmployee.salary)
	}

}
