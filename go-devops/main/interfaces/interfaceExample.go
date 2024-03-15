package main

import (
	"fmt"
	"strings"
)

// Stringer is a real type defined in the standard library's fmt package.
type Stringer interface {
	Ztring() string
}

type Person struct {
	First, Last string
}

// This is a way to bind the function to the struct. So the function can be called with the same name.
func (p Person) Ztring() string {
	return fmt.Sprintf("%s %s", p.First, p.Last)
}

type StrList []string

func (s StrList) Ztring() string {
	return strings.Join(s, ",")
}

// PrintStringer prints the value of a Stringer to stdout.
func PrintStringer(s Stringer) {
	fmt.Println(s.Ztring())
}

func main() {
	john := Person{First: "John", Last: "Doak"}
	var nameList Stringer = StrList{"David", "Sarah"}

	PrintStringer(john)
	PrintStringer(nameList)

}
