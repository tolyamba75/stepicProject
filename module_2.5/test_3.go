package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, e string
	fmt.Scan(&str)
	fmt.Scan(&e)
	fmt.Println(strings.Index(str, e))

}
