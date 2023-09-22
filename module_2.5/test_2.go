package main

import (
	"fmt"
)

func main() {

	var str string
	rts := ""

	fmt.Scan(&str)

	for _, v := range str {
		rts = string(v) + rts
	}

	// fmt.Println(rts == str)
	if rts == str {
		fmt.Println("Палиндром")
		return
	}
	fmt.Println("Нет")
}
