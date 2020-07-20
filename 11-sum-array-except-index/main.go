package main

import "fmt"

func myFunc(nums []int) []int {
	res := []int{}
	before := []int{}
	after := []int{}
	currB := 1
	currA := 1

	for i := 0; i < len(nums); i++ {
		res = append(res, 1)
		before = append(before, 1)
		after = append(after, 1)
	}

	before[0] = 1
	after[len(nums)-1] = 1

	for i := 1; i < len(nums)-1; i++ {
		j := len(nums) - i - 1

		currB *= nums[i-1]
		currA *= nums[j+1]
		before[i] = currB
		after[j] = currA
	}

	before[len(nums)-1] = currB * nums[len(nums)-2]
	after[0] = currA * nums[1]

	for i := range nums {
		res[i] = before[i] * after[i]
	}

	return res
}

func main() {
	nums := []int{1, 0, 3, 4}

	fmt.Println(myFunc(nums))
}
