package main

import (
	_ "github.com/lib/pq"
)

func main() {

	init_db()
	defer db.Close()

	bash()

}

type Todo struct {
	title      string
	descripton string
	is_done    bool
}
