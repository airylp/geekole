package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//geekole.com

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	db, err := sql.Open("mysql", "user01:password01@tcp(localhost:3306)/mydatabase?parseTime=true")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, username, password, email FROM users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(user.Username + " " + user.Password + " " + user.Email)
	}
}
