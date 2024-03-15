package main

import (
	"fmt"
	//"main/slice"
)

func main() {
	modelToMake := map[string]string{
		"name": "John Doe",
		"age":  "30",
	}
	modelToMake["outback"] = "subaru"
	if carMake, ok := modelToMake["outback"]; ok {
		fmt.Printf("car model has %q\n", carMake)
	} else {
		fmt.Printf("car model has unknown make\n")
	}
	// Get all values from map:
	for key, val := range modelToMake {
		fmt.Printf("car model %q has make %q\n", key, val)
	}

}
