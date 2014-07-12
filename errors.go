package main

import "errors"
import "fmt"

func f1(i int) (int, error) {
	if i == 42 {
		return -1, errors.New("Cannot work with 42")
	}

	return i + 3, nil

}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}
}