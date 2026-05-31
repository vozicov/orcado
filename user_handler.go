package main

import (
	"fmt"
	"log"
)

func get_all_users() []User {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		println(err)
	}
	defer rows.Close()

	var users []User = []User{}

	for rows.Next() {
		newuser := User{}
		rows.Scan(&newuser.ID, &newuser.user_name, &newuser.is_admin, &newuser.password)
		users = append(users, newuser)

	}
	return users
}

func find_by_username(username string) (User, error) {

	res, err := db.Query("SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		log.Println(err)
	}
	defer res.Close()

	found_user := User{}

	for res.Next() {
		res.Scan(&found_user.ID, &found_user.user_name, &found_user.is_admin, &found_user.password)
		return found_user, nil
	}

	return User{}, fmt.Errorf("No user found")
}

func create_user(username string) User {
	_, err := db.Exec("INSERT INTO users (username, isadmin) VALUES ($1, false)", username)
	if err != nil {
		log.Fatal(err)
	}
	created_user, _ := find_by_username(username)
	return created_user
}
