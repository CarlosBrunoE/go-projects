package main

import (
	"fmt"
	//"main/slice"
)

type Record struct {
	Name string
	Age  int
}

func changeName(r Record) {
	r.Name = "Peter"
	fmt.Println("inside change name: ", r.Name)
}

func main() {

	rec := Record{Name: "John"}
	changeName(rec)
	fmt.Println("after changing value: ", rec.Name)

}
