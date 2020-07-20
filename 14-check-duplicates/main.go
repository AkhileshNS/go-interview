package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func myFunction(arr []int) int {
	start := 1          // range starts from 1 to (n - 1)
	end := len(arr) - 1 // where n is length of the array

	for start <= end {
		mid := int(math.Round(float64(end+start) / 2))
		if mid == end || mid == start {
			return mid
		}

		lesserCount := 0

		for _, num := range arr {
			if num <= mid {
				lesserCount++
			}
		}

		if lesserCount == mid {
			start = mid
		} else {
			end = mid
		}
	}

	return start
}

func main() {
	nums := []int{1, 9, 8, 7, 6, 6, 5, 4, 3, 2}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	fmt.Println(nums)
	fmt.Println(myFunction(nums))
}
