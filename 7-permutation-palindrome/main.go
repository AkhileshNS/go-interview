package main

import "fmt"

func myFunction(str string) bool {
	letters := make(map[string]int)

	for _, letter := range str {
		if _, ok := letters[string(letter)]; ok {
			letters[string(letter)]++
		} else {
			letters[string(letter)] = 1
		}
	}

	noOfOdds := 0

	for _, count := range letters {
		if count%2 == 1 {
			noOfOdds++
		}
		if noOfOdds > 1 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(myFunction("ivicl"))
}
