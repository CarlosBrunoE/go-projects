package main

import (
	"fmt"
	//"main/slice"
)

type Record struct {
	Name string
	Age  int
}

func main() {
	david := Record{Name: "David", Age: 25}
	john := Record{Name: "John", Age: 32}
	fmt.Printf("%+v\n", david) // prints Name: David
	fmt.Printf("%+v\n", john)  // prints Name: John

}
