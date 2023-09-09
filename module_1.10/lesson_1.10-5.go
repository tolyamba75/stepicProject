package main

import "fmt"

func main() {
	var number, c, d int

	fmt.Scan(&number)
	fmt.Scan(&c) // кратное c
	fmt.Scan(&d) // НЕ кратное d

	for i := c; i <= number; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
