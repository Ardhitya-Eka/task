package main

import (
	"encoding/json"
	"os"
)

func main() {

	// 1. Load data
	var data Data
	file, err := os.ReadFile("task.json")
	if err == nil {
		json.Unmarshal(file, &data)
	}

	// 2. Run CLI
	RunCLI(&data, os.Args)
}
