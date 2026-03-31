package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RunCLI(data *Data, args []string) {
	if len(args) < 2 {
		fmt.Println("Command not provided")
		return
	}

	command := args[1]
	argsJoin := args[3:]
	result := strings.Join(argsJoin, " ")
	switch command {

	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: add <description>")
			return
		}

		err := add(result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	case "delete":
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		delete(id)

	case "update":

		if len(args) < 4 {
			fmt.Println("Usage: update <id> <description>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		description := result

		err = updateDescription(id, description)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Updating Task")

	case "update-status":
		id, _ := strconv.Atoi(args[2])
		status := Status(args[3])
		changeStatus(id, status)

	default:
		fmt.Println("Unknown command")
	}
}
