package main

import "fmt"

func main() {
	var number, maxNumber, countMax int
	countMax, maxNumber = 0, 0

labl:
	for {
		fmt.Scan(&number)

		switch {
		case number == 0:
			break labl
		case number > maxNumber:
			maxNumber = number
			countMax = 1
		case number == maxNumber:
			countMax++
		}

	}
	fmt.Print(countMax)

}
