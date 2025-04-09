package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		formatter func(string) string
		messages  []string
		expected  []string
	}

	runCases := []testCase{
		{
			formatter: addSignature,
			messages:  []string{"Hello, how are you?", "I hope you are well,"},
			expected:  []string{"Hello, how are you? kind regards.", "I hope you are well, kind regards."},
		},
		{
			formatter: addGreeting,
			messages:  []string{"I'm doing well.", "Love your hair!"},
			expected:  []string{"Hello! I'm doing well.", "Hello! Love your hair!"},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			formatter: addGreeting,
			messages:  []string{"", ""},
			expected:  []string{"Hello! ", "Hello! "},
		},
		{
			formatter: addGreeting,
			messages:  []string{"I'm so sick of this crap.", "I need a change.", "Maybe I should go touch grass."},
			expected:  []string{"Hello! I'm so sick of this crap.", "Hello! I need a change.", "Hello! Maybe I should go touch grass."},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		messages := getFormattedMessages(test.messages, test.formatter)
		for i, message := range messages {
			expected := test.expected[i]
			input := test.messages[i]
			if message != expected {
				failCount++
				t.Errorf(`---------------------------------
Test Failed:
  input:     %v
  Expecting: %v
  Actual:    %v
  Fail
`, input, expected, message)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed:
  input:     %v
  Expecting: %v
  Actual:    %v
  Pass
`, input, expected, message)
			}
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func addSignature(message string) string {
	return message + " kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
