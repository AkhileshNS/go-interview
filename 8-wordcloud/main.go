package main

import (
	"fmt"
	"strings"
	"unicode"
)

func myFunction(sentence string) map[string]int {
	res := make(map[string]int)

	formattedSentence := ""
	word := ""
	for _, letter := range sentence {
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			word += string(letter)
		} else {
			if word != "" {
				formattedSentence += word + " "
				word = ""
			}
		}
	}

	words := strings.Fields(strings.ToLower(formattedSentence))

	for _, word := range words {
		if _, ok := res[word]; ok {
			res[word]++
		} else {
			res[word] = 1
		}
	}

	return res
}

func main() {
	fmt.Println(myFunction("Hello darkness my old friend, Hello from the other side"))
}
