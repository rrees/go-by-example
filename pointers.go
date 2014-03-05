package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroOperator(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroOperator(&i)
	fmt.Println("zeroOperator:",i)

	fmt.Println("Pointer:", &i)
}