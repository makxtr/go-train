package main

import (
	"strconv"
	"strings"
)

/**
Have the function RunLength(str) take the str parameter being passed and return a compressed version of the string using the Run-length encoding algorithm.
This algorithm works by taking the occurrence of each repeating character and outputting that number along with a single character of the repeating sequence.
For example: "wwwggopp" would return 3w2g1o2p. The string will not contain any numbers, punctuation, or symbols.
Examples
Input: "aabbcde"
Output: 2a2b1c1d1e
Input: "wwwbbbw"
Output: 3w3b1w.
*/

func RunLength(str string) string {
	if len(str) == 0 {
		return ""
	}

	var sb strings.Builder
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(str[i-1])
			count = 1
		}
	}
	sb.WriteString(strconv.Itoa(count))
	sb.WriteByte(str[len(str)-1])

	return sb.String()
}
