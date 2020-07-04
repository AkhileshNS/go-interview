package main

import "fmt"

func reverse(str []string) {
	for i := 0; i < len(str)/2; i++ {
		// Equivalent index backwards : 0 - (length - 1)
		j := len(str) - (i + 1)
		str[i], str[j] = str[j], str[i]
	}
}

func main() {
	q := []string{"h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"}
	reverse(q)
	fmt.Print(q)
}
