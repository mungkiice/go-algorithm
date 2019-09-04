package main

import (
	"fmt"
	"regexp"
)

func firstRepeatedWord(s string) string {
	var temp = map[string]bool{}
	words := regexp.MustCompile(`\s+`).Split(s, -1)
	for _,word := range words {
		if _,ok := temp[word];ok{
			return word
		}
		temp[word] = true
	}
	return ""
}

func main() {
	fmt.Println(firstRepeatedWord("He had had quite enough of this nonsense."))
}