package main

import "fmt"

func main() {
	s := make([]string, 3)

	fmt.Println("Empty: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Set: ", s)
}