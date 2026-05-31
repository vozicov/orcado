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
		rows.Scan(&newuser.ID, &newuser.user_name, &newuser.is_admin)
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
		res.Scan(&found_user.ID, &found_user.user_name, &found_user.is_admin)
		return found_user, nil
	}

	return User{}, fmt.Errorf("No user found")
}
