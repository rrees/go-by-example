package main 

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Map: ", m)

	v1 := m["k1"]
	fmt.Println("Get: ", v1)

	fmt.Println("Length: ", len(m))

	delete(m, "k2")
	fmt.Println("Delete: ", m)

	_, present := m["k2"]
	fmt.Println("prs: ", present)

	n := map[string]int{"foo" : 1, "bar" : 2}
	fmt.Println("Map n: ", n)
}