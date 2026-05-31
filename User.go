package main

import (
	"database/sql"
	"log"
)

type User struct {
	ID        int
	user_name string
	is_admin  bool
}

func create_user(db *sql.DB, username string) {
	_, err := db.Exec("INSERT INTO users (username, isadmin) VALUES ($1, false)", username)
	if err != nil {
		log.Fatal(err)
	}

}

func (u *User) promote_user() {
	u.is_admin = true
	u.update_db()
}

func (u *User) update_db() {
	db.Exec("UPDATE users SET username=$1, isadmin=$2 WHERE userid=$3", u.user_name, u.is_admin, u.ID)
}
