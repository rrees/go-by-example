package main

import "fmt"

func integerSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInteger := integerSequence()

	fmt.Println(nextInteger())
	fmt.Println(nextInteger())
	fmt.Println(nextInteger())

	newNextInteger := integerSequence()

	fmt.Println(newNextInteger())	
}