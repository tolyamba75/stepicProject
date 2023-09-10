package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	for _, elem := range array {
		if a <= elem {
			a = elem
		}
	}
	fmt.Println(a)
}
