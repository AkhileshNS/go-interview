package main

import (
	"fmt"
	"math"
)

func getMid(start, end int) int {
	return int(math.Round(float64(end+start) / 2))
}

func findRotatePoint(words []string) int {
	// Linear Approach
	/* for i := 1; i < len(words); i++ {
			if words[i] < words[i-1] {
				return i
			}
	    } */

	// [7, 8, 9, 2, 3, 4, 5]

	// Binary Approach
	res := -1
	start := 0
	end := len(words)
	for start < end {
		mid := getMid(start, end)

		if words[mid] < words[mid-1] {
			return mid
		}

		if words[mid] < words[0] {
			end = mid
		} else {
			start = mid
		}
	}

	return res
}

func main() {
	words := []string{
		"ptolemaic",
		"retrograde",
		"supplant",
		"undulate",
		"xenoepist",
		"asymptote", // <-- rotates here!
		"babka",
		"banoffee",
		"engender",
		"karpatka",
		"othellolagkage",
	}

	fmt.Println(findRotatePoint(words))
}
