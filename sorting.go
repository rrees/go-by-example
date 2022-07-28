package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := []int{3, 7, 1, 19, 23, 8, 3, 14}

	sort.Ints(numbers)

	fmt.Println("Numbers: ", numbers)
}