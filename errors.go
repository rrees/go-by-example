package main

import "errors"
import "fmt"

func f1(i int) (int, error) {
	if i == 42 {
		return -1, errors.New("Cannot work with 42")
	}

	return i + 3, nil

}

type argError struct {
	arg int
	problem string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Cannot work with 42 (sorry)"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.problem)
	}
}