package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())

	rp := &r

	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perimeter())
}