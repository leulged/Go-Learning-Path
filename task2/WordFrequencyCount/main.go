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
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func main() {
	fmt.Println("Please enter a sentence:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.ToLower(input)
	cleaned := cleanInput(input)
	words := strings.Fields(cleaned)

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	fmt.Println("Word Frequencies:")
	for word, count := range wordFreq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
