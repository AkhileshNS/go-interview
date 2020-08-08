package main

import (
	"fmt"
)

type cake struct {
	w int
	v int
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func maxDuffelBag(cakes []cake, capacity int) int {
	maxValues := []int{0}

	for i := 1; i <= capacity; i++ {
		maxValue := 0

		for _, cake := range cakes {
			currMax := 0

			if cake.w <= i {
				currMax += cake.v

				remainder := i - cake.w
				if remainder >= 0 && remainder < len(maxValues) {
					currMax += maxValues[remainder]
				}
			}

			maxValue = max(maxValue, currMax)
		}

		maxValues = append(maxValues, maxValue)
	}

	return maxValues[capacity]
}

func main() {
	cakes := []cake{cake{w: 7, v: 160}, cake{w: 3, v: 90}, cake{w: 2, v: 15}}

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(maxDuffelBag(cakes, 20))
}
