package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go mysql")
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/go_mysql_example")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("connect ok")

	// insert, err := db.Query("INSERT INTO Users VALUES('hung')")
	// if err != nil{
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("Successfully inserted into user table")

	result, err := db.Query("SELECT name FROM Users")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var user User
		err = result.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}
}
