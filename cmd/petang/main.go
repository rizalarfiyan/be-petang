package main

import (
	"fmt"

	"github.com/rizalarfiyan/be-petang/config"
	"github.com/rizalarfiyan/be-petang/database"
)

func main() {
	conf := config.Get()
	db := database.Postgres()
	defer db.Close()

	fmt.Println("App Running!")
	fmt.Println(conf)
}
