package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//here i don't use any password, but if you want to use
	// password you should use a user:password@
	// check this link: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	db, err := sql.Open("mysql", "@tcp(127.0.0.1:9306)/")

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal(err)
	}

	insert, err := db.Query("INSERT INTO test (id, version) VALUES (1, 'RT')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Inserted")

}
