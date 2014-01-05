package main

import "fmt"
import "time"

func main() {

	i := 2
	fmt.Print("Writing i = ", i, " as ")

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend!")
		default:
			fmt.Println("It is a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() <= 12:
		fmt.Println("It's the morning")
	default:
		fmt.Println("It's after noon")
	}
}