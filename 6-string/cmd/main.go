package main

import (
	"strconv"
	"strings"
	"time"
)

func main() {
	// str := ""
	// for i := 0; i < 100000; i++ {
	// 	str += fmt.Sprintf("%d", i)
	// }

	var builder strings.Builder

	timeStart := time.Now()
	for i := 0; i < 100_000; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	println(time.Since(timeStart).Milliseconds())
	//println(builder.String())
}
