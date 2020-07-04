package main

import "fmt"

type word struct {
	startIndex int
	length     int
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getWords(s []string) []word {
	res := []word{}
	spaceIndex := -1

	for i, letter := range s {
		if letter == " " {
			res = append(res, word{startIndex: spaceIndex + 1, length: i - spaceIndex - 1})
			spaceIndex = i
		}
	}

	res = append(res, word{startIndex: spaceIndex + 1, length: len(s) - spaceIndex - 1})
	return res
}

func reverseWords(s []string) {
	reverse(s)
	words := getWords(s)
	track := 0

	for i := 0; i < len(s); i++ {
		// j = lastIndex - currentIndex - track's startIndex
		j := (words[track].length + words[track].startIndex - 1) - (i - words[track].startIndex)
		s[i], s[j] = s[j], s[i]
		trackLength := words[track].length / 2
		if words[track].length%2 == 0 {
			trackLength--
		}
		if i-words[track].startIndex == trackLength {
			track++
			if track == len(words) {
				break
			}
			i = words[track].startIndex - 1
		}
	}
}

func main() {
	q := []string{"c", "a", "k", "e", " ", " p", "o", "u", "n", "d", " ", "s", "t", "e", "a", "l"}
	reverseWords(q)
	fmt.Print(q)
}
