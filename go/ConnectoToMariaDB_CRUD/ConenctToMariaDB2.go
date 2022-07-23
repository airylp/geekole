package main

import (
	"database/sql"
	"log"
	"strconv"

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

	_, errCreate := db.Exec(`CREATE TABLE users (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    PRIMARY KEY (id)
    )`)
	if errCreate != nil {
		panic(err.Error())
	}

	for i := 1; i <= 10; i++ {
		_, errInsert := db.Exec(`INSERT INTO users(username, password, email) VALUES (?, ?, ?)`,
			"user"+strconv.Itoa(i),
			strconv.Itoa(i),
			"myemail"+strconv.Itoa(i)+"@test.com")
		if errInsert != nil {
			panic(err.Error())
		}
	}

	_, errDelete := db.Exec(`DELETE FROM users WHERE id=?`, 4)
	if errDelete != nil {
		panic(err.Error())
	}

	_, errUPDATE := db.Exec(`UPDATE users SET username=? WHERE id=?`, "useraname", 1)
	if errUPDATE != nil {
		panic(err.Error())
	}

	results, errSelect := db.Query(`SELECT id, username, password, email FROM users`)
	if errSelect != nil {
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
