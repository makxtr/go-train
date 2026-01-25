package main

import (
	"strconv"
	"strings"
)

/**

Have the function ReversePolishNotation(str) read str which will be an arithmetic expression
composed of only integers and the operators: +,-,* and / and the input expression will be in postfix notation (Reverse Polish notation),
an example: (1 + 2) * 3 would be 1 2 + 3 * in postfix notation.
Your program should determine the answer for the given postfix expression.
For example: if str is 2 12 + 7 / then your program should output 2.
Examples
Input: "1 1 + 1 + 1 +"
Output: 4
Input: "4 5 + 2 1 + *"
Output: 27
*/

func ReversePolishNotation(str string) int {
	tokens := strings.Fields(str)
	var stack []int

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				// Should not happen for valid RPN
				return 0
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			stack = append(stack, res)
		}
	}

	if len(stack) == 0 {
		return 0
	}
	return stack[0]
}
