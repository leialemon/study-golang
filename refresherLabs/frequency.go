package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := map[string]int{"The":0, "quick":0, "brown":0, "fox":0, "jumps":0,
								"over":0, "the":0, "lazy":0, "dog":0}

	for i:=0; i<len(words);i++{
		var word string = words[i]
		var counter int = 0
		for _, element := range words{
			if element == word{
				counter++
			}
		}
		wordCount[word] = counter
	}
	return wordCount
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}