/*
Problem Statement:
Write a function that accepts a string consisting of only brackets ([](){}) and returns whether
it is balanced. Every "opening" bracket must be followed by a closing bracket of the same type.
There can also be nested brackets, which adhere to the same rule.

Sample Inputs:
#1
Input: ()[]{}(([])){[()][]}
Output: true
#2
Input: ())[]{}
Output: false
#3
Input: [(])
Output: true

Error handling:
1. Empty String
Input:
Output: False

2. String consisting of any character other than brackets
Input: abc
Output: False
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter string (consisting of only brackets):")
	fmt.Scanf("%s", &str)

	result := CheckBalancedBracket(str)

	fmt.Println(result)
}

// CheckBalancedBracket to check input string consist only of brackets and if it is balanced
func CheckBalancedBracket(str string) bool {
	var countRoundBracket, countSquareBracket, countCurlyBracket int
	var result bool

	// check empty string
	if len(strings.TrimSpace(str)) == 0 {
		// empty string
		return false
	}

	for _, v := range str {
		switch string(v) {
		case "(":
			countRoundBracket++
		case "[":
			countSquareBracket++
		case "{":
			countCurlyBracket++
		case ")":
			countRoundBracket--
		case "]":
			countSquareBracket--
		case "}":
			countCurlyBracket--
		default:
			// invalid string. Error: string can consist of only brackets ([](){})"
			return false
		}

		if countRoundBracket < 0 || countSquareBracket < 0 || countCurlyBracket < 0 {
			// invalid string. Error: closing bracket found before opening bracket"
			return false
		}
	}

	// check
	if countRoundBracket == 0 && countSquareBracket == 0 && countCurlyBracket == 0 {
		result = true
	} else {
		result = false
	}

	// return
	return result
}
