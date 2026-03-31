package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadTask() ([]Todo, error) {
	var todos []Todo

	file, err := os.ReadFile("task.json")
	if err != nil {
		return todos, nil
	}

	err = json.Unmarshal(file, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func writeFile(data Data) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		fmt.Println("Error Data :", err)
		return err
	}

	filePath := "task.json"
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error Write Data : ")
		return err
	}

	return nil
}
