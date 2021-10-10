package handler

import (
	"fmt"
	"ginWeb/database"
	"os"
)

func Command() {
	args := os.Args
	fmt.Println(args)
	if len(args) > 1 && args[1] != "main.go" {
		first := args[1]
		second := ""
		if len(args) > 2 {
			second = args[2]
		}

		if first == "migrate-seed" {
			database.InitialMigration()
			database.InitialDBSeeder()
			os.Exit(0)
		} else if first == "migrate" {
			database.InitialMigration()
		} else if first == "seed" {
			database.InitialDBSeeder()
		}

		if first != "" && second == "" {
			os.Exit(0)
		}
	}
}
