package main

import "fmt"

func main() {
	var n int
	var count int = 0

	fmt.Scan(&n)
	slice := make([]int, n)

	for i := range slice {
		fmt.Scan(&slice[i])
		if slice[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}
