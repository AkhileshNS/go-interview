package main

import "fmt"

func reverseLetters(s []string, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseWords(s []string) {
	reverseLetters(s, 0, len(s)-1)
	spaceIndex := -1

	for i, letter := range s {
		if letter == " " {
			reverseLetters(s, spaceIndex+1, i-1)
			spaceIndex = i
		}
	}

	reverseLetters(s, spaceIndex+1, len(s)-1)
}

func main() {
	q := []string{"c", "a", "k", "e", " ", " p", "o", "u", "n", "d", " ", "s", "t", "e", "a", "l"}
	reverseWords(q)
	fmt.Print(q)
}
