package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerState struct {
	id          int64  `json:"id"`
	customer_id string `json:"customer_id"`
	status      int    `json:"status"`
}

func main() {
	fmt.Println("Welcome to MySql Connection with Go")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cmp")
	checkError(err)

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	fmt.Println("Success")

	//Insert
	// insert, err := db.Query("INSERT INTO customer_migration(product, customer_id, status) values('PRD1', 'Cust1', 1)")
	// checkError(err)
	// fmt.Println("Records inserted")
	// defer insert.Close()

	//Fetch by customer id and product id
	result, err := db.Query("SELECT id, customer_id, status FROM customer_migration")
	checkError(err)

	for result.Next() {
		var customer CustomerState
		err = result.Scan(&customer.id, &customer.customer_id, &customer.status)
		checkError(err)

		fmt.Printf("ID %v CustomerId %v Status %d\n", customer.id, customer.customer_id, customer.status)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
