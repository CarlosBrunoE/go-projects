package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./readme.json")
	if err != nil {
		panic(err)
	}
	s := string(data)
	fmt.Println(s)
}
