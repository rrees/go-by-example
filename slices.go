package main

import "fmt"

func main() {
	s := make([]string, 3)

	fmt.Println("Empty: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[2])
	fmt.Println("Len: ", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append: ", s)

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println("Copy: ", c)

	l1 := s[2:5]
	fmt.Println("Slice: ", l1)

	l2 := s[:5]
	l3 := s[2:]

	fmt.Println("Slice to: ", l2)
	fmt.Println("Slice from: ", l3)

	t := []string{"g", "h", "i"}
	fmt.Println("Declared slice: ", t)

	twoDee := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoDee[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDee[i][j] = i + j
		}
	}
	fmt.Println("Multi-dimensional slice: ", twoDee)
}