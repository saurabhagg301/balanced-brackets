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

#4
Input: [(()])
Output: true


Error handling:
1. Empty String
Input:
Output: false

2. String consisting of any character other than brackets
Input: abc
Output: false
*/
package main

import (
	"testing"
)

func TestCheckBalancedBracket(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"invalid string", "abc", false},
		{"Sample input #1", "()[]{}(([])){[()][]}", true},
		{"Sample input #2", "())[]{}", false},
		{"Sample input #3", "[(])", true},
		{"Sample input #4", "[((])", false},
	}

	for _, tt := range tests {
		result := CheckBalancedBracket(tt.input)
		if result != tt.expected {
			t.Errorf("Error: TestName: %v, expected: %v, got: %v", tt.testName, tt.expected, result)
		} else {
			t.Logf("Success: TestName: %v, expected: %v, got: %v", tt.testName, tt.expected, result)
		}
	}
}
