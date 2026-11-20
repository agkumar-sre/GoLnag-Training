package main

import (
	"fmt"
	"strings"
)

func main(){
	text := "Go is a simple and effective programming language. Go is designed for building fast software"

	wordCount := make(map[string]int)
	words:= strings.Fields(text)

	for _, word := range words {
		word := strings.ToLower(word)
		wordCount[word]++
	}

	fmt.Println("Paragraph: ")
	fmt.Println(text)

	fmt.Println("\nWord Frequency: ")
	for word, count := range wordCount {
		fmt.Println(word, ":", count)
	}
}