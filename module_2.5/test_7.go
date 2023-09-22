package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str string

	fmt.Scan(&str)
	validMask := regexp.MustCompile(`^[a-zA-Z0-9]{5,}$`)

	if validMask.MatchString(str) {
		fmt.Println("Ok")
		return
	}
	fmt.Println("Wrong password")
}
