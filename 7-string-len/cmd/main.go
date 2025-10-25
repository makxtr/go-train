package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "ddЯЙ描"

	fmt.Println(len(str))                    // bytes not runes
	fmt.Println(utf8.RuneCountInString(str)) //5
}
