package day1

import (
	"fmt"
	"regexp"
	"strings"
)

type Digit struct {
	Forward     string
	Backward    string
	StringValue string
	Value       int
}

type Digits []Digit

func (digits Digits) ForwardPattern() string {
	groups := make([]string, 0)
	for _, digit := range digits {
		groups = append(
			groups,
			fmt.Sprintf("(%s)", digit.StringValue),
			fmt.Sprintf("(%s)", digit.Forward),
		)
	}
	return strings.Join(groups, "|")
}

func (digits Digits) BackwardPattern() string {
	groups := make([]string, 0)
	for _, digit := range digits {
		groups = append(
			groups,
			fmt.Sprintf("(%s)", digit.StringValue),
			fmt.Sprintf("(%s)", digit.Backward),
		)
	}
	return strings.Join(groups, "|")
}

func (digits Digits) Map() map[string]int {
	result := make(map[string]int)
	for _, digit := range digits {
		result[digit.StringValue] = digit.Value
		result[digit.Forward] = digit.Value
		result[digit.Backward] = digit.Value
	}
	return result
}

func Reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FindValue(input string) int {
	if len(input) == 0 {
		return 0
	}

	digits := Digits{
		{"one", "eno", "1", 1},
		{"two", "owt", "2", 2},
		{"three", "eerht", "3", 3},
		{"four", "ruof", "4", 4},
		{"five", "evif", "5", 5},
		{"six", "xis", "6", 6},
		{"seven", "neves", "7", 7},
		{"eight", "thgie", "8", 8},
		{"nine", "enin", "9", 9},
	}

	validDigit := regexp.MustCompile(digits.ForwardPattern())
	validReverseDigit := regexp.MustCompile(digits.BackwardPattern())

	digitMap := digits.Map()
	firstDigit := digitMap[validDigit.FindString(input)]
	lastDigit := digitMap[validReverseDigit.FindString(Reverse(input))]

	return firstDigit*10 + lastDigit
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += FindValue(input)
	}
	return sum
}
