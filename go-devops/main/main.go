package main

import (
	"fmt"
	"os"
)

func main() {
	// The string to write to the file
	content := "Hello, world!"

	// The file path where you want to write the string
	filePath := "test.txt"

	// Write the string to the file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("String written to file successfully.")
}
