package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// var text string
	var ans bool
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// text := "Быть или не быть."
	rtext := []rune(text)
	if unicode.IsUpper((rtext[0])) && rtext[len(rtext)-1] == '.' {
		ans = true
	}
	if ans {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
