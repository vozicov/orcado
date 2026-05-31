package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bash() {
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
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print(current_user.user_name + "@todo:~$ ")

		cmd, _ := reader.ReadString('\n')
		// fmt.Scanf("%s\n", &cmd)
		fmt.Printf("DEBUG: Command was: '%s'\n", cmd)
		cmd = strings.TrimSuffix(cmd, "\n")
		if cmd == "exit" {
			break
		}

		if current_user.is_admin == true {
			fmt.Printf("You must be administrator to run commands!\n - Run exit to exit\n")
			continue
		}

		if cmd == "moew" {
			fmt.Printf("meow!\n")
		}

		if strings.HasPrefix(cmd, "user ") {
			cmd = strings.TrimPrefix(cmd, "user ")
			if strings.HasPrefix(cmd, "new ") {
				name := strings.TrimPrefix(cmd, "new ")
				var created_user User = create_user(name)
				fmt.Println("User " + created_user.user_name + " with id " + strconv.Itoa(created_user.ID) + " created with no errors\n")
			}

			if strings.HasPrefix(cmd, "promote ") {
				name := strings.TrimPrefix(cmd, "promote ")
				fmt.Println("'" + name + "'")
				promoted_user, err := find_by_username(name)
				if err != nil {
					fmt.Println(err)
					return
				}
				promoted_user.promote_user()
				fmt.Println("User " + promoted_user.user_name + " has been promoted")
			}
		}
	}
}
