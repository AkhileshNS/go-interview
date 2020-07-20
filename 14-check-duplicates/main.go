package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// O(nlog(n)) : O(1)
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

// O(n) : O(1)
func checkDuplicates(arr []int) int {
	curr := arr[len(arr)-1]

	for i := 0; i < len(arr); i++ {
		curr = arr[curr-1]
	}

	newCurr := arr[curr-1]
	size := 1

	for newCurr != curr {
		newCurr = arr[newCurr-1]
		size++
	}

	curr = arr[len(arr)-1]
	follower := arr[len(arr)-1]
	for i := 0; i < len(arr); i++ {
		if i >= size {
			follower = arr[follower-1]
		}
		curr = arr[curr-1]
		if follower == curr {
			return follower
		}
	}

	return -1
}

func main() {
	nums := []int{4, 3, 1, 1, 4}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	fmt.Println(nums)
	fmt.Println(checkDuplicates(nums))
}
