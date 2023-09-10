package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	baseSlice := make([]int, n)

	for i := range baseSlice {
		fmt.Scan(&baseSlice[i])
	}
	fmt.Println(baseSlice[3])
}
