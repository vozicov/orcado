package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	init_db()
	defer db.Close()

	fmt.Println("Welcome!\nEnter your username: ")

	var username string
	fmt.Scanln(&username)
	current_user, err := find_by_username(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Welcome " + current_user.user_name + "!\nAvaible users are:\n")
	users := get_all_users()
	for _, user := range users {
		println("Username: ", user.user_name, "\n - Is Admin?: ", user.is_admin)
	}

	for true {
		fmt.Print(current_user.user_name + "@todo:~$ ")

		var cmd string
		fmt.Scanln(&cmd)

		if cmd == "exit" {
			break
		}

		if current_user.is_admin == false {
			fmt.Printf("You must be administrator to run commands!\n - Run exit to exit\n")
			continue
		}

		if cmd == "moew" {
			fmt.Printf("meow!\n")
		}
	}

}

type Todo struct {
	title      string
	descripton string
	is_done    bool
}
