package main

import (
	"fmt"
	"os"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/seeders"
)

// HELP message. For spacing and alignment use tabs
const HELP = `Developer CLI

Usage:

go run cmd/dev/main.go <command>

Available commands:

run:seeders		- Run seeders
help 			- Show this help message`

func main() {
	config.SetVariable("DATABASE_HOST", "localhost")
	config.LoadConfig(".")

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Missing command to execute")
		fmt.Println(HELP)
		return
	}

	switch args[0] {

	case "run:seeders":
		fmt.Println("Running seeders")
		if err := seeders.Run(); err != nil {
			fmt.Println("Error running seeders")
			fmt.Println(err)
			return
		}
		fmt.Println("Seeders ran successfully")

	case "help":
		fmt.Println(HELP)

	default:
		fmt.Println("Invalid command")
		fmt.Println(HELP)
	}
}
