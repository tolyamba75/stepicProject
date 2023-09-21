package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	result := ""
	for _, lit := range str {
		if strings.Count(str, string(lit)) == 1 {
			result += string(lit)
		}
	}
	fmt.Println(result)
}
