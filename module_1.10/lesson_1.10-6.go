package main

import "fmt"

func main() {
	var number int

labelScan:
	for {
		fmt.Scan(&number)
		switch {
		case number > 100:
			break labelScan
		case number < 10:
			break
		default:
			fmt.Println(number)
		}

	}

}
