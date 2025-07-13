package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func cleanInput(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func palindrome(word string) bool {
	start := 0
	end := len(word) - 1

	for start < end {
		if word[start] != word[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	fmt.Println("Please enter a sentence:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.ToLower(input)
	cleaned := cleanInput(input)

	if palindrome(cleaned) {
		fmt.Println("It is a palindrome.")
	} else {
		fmt.Println("It is not a palindrome.")
	}
}
