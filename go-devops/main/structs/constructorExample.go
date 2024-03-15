package main

import (
	"fmt"
	//"main/slice"
)

type Record struct {
	Name string
	Age  int
}

func NewRecord(name string, age int) (*Record, error) {
	if name == "" || age <= 0 {
		return nil, fmt.Errorf("name can't be empty and  age must be greater than zero")
	}
	return &Record{Name: name, Age: age}, nil
}

func main() {
	rec, err := NewRecord("John Doe", 100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rec)
	}

}
