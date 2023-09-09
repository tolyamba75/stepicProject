package main

import "fmt"

func main() {
	var number, maxNumber, countMax int
	countMax, maxNumber = 0, 0

	for fmt.Scan(&number); number != 0; fmt.Scan(&number) {

		switch {
		case number > maxNumber:
			maxNumber = number
			countMax = 1
		case number == maxNumber:
			countMax++
		}

	}
	fmt.Print(countMax)

}
