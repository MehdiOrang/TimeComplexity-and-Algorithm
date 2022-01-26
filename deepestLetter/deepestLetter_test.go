package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func getDeepestLetter(s string) rune {
	currentMax := 0
	mainMax := 0
	result := '?'
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] == '(' {
			currentMax++
			if currentMax > mainMax {
				mainMax = currentMax
				result = rune(s[i+1])
			}
		} else if s[i] == ')' {
			if currentMax > 0 {
				currentMax--
			} else {
				return '?'
			}
		}
	}

	if currentMax != 0 {
		return '?'
	}

	return result

}

type letterTestCase struct {
	input    string
	expected rune
}

func TestGetDeepestLetter(t *testing.T) {
	testCases := []letterTestCase{
		{input: "a(b)c", expected: 'c'},                                   //let it fails purposely
		{input: "((a))(((M)))(c)(D)(e)(((f))(((G))))h(i)", expected: 'G'}, //alright
		{input: "((A)(b)c", expected: '?'},                                //alright
		{input: "(a)((G)c)", expected: 'G'},                               //alright
		{input: "(8)", expected: '8'},                                     //alright
		{input: "(!)", expected: '!'},                                     //alright
	}

	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			actual := getDeepestLetter(test.input)
			if !cmp.Equal(test.expected, actual) {
				t.Log(cmp.Diff(string(test.expected), string(actual)))       //first way
				t.Logf("Expected '%c', but got '%c'", test.expected, actual) //second way
				t.Fail()
			}
		})
	}
}
