package slice

import "fmt"

func DoAppend(sl []int) {
	sl = append(sl, 100)
	fmt.Println("After appending:", sl)

}

func GetSlice(someSlice []int) []int {
	for i := 0; i < len(someSlice); i++ {
		fmt.Printf("slice entry %d: %s\n", i, someSlice[i])
	}
	return someSlice

}

func GetSliceRange(someSlice []int) []int {
	for index, val := range someSlice {
		fmt.Printf("slice entry %d: %s\n", index, val)
	}
}
