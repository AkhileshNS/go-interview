package main

import "fmt"

func greatest3(nums []int) int {
	highest := nums[0]
	secondHighest := nums[1]
	thirdHighest := nums[2]

	lowest := nums[0]
	secondLowest := nums[1]

	for _, num := range nums {
		if num > highest {
			thirdHighest = secondHighest
			secondHighest = highest
			highest = num
		} else if num > secondHighest && num < highest {
			thirdHighest = secondHighest
			secondHighest = num
		} else if num > thirdHighest && num < secondHighest {
			thirdHighest = num
		}

		if num < lowest {
			secondLowest = lowest
			lowest = num
		} else if num < secondLowest && num > lowest {
			secondLowest = num
		}
	}

	if highest*secondHighest*thirdHighest > lowest*secondLowest*highest {
		return highest * secondHighest * thirdHighest
	}
	return lowest * secondLowest * highest
}

func main() {
	nums := []int{8, 4, 9, 5, 7, 1}
	nums2 := []int{-10, -10, 1, 3, 2}
	nums3 := []int{-10, -10, -1, -3, -2}
	fmt.Println(greatest3(nums))
	fmt.Println(greatest3(nums2))
	fmt.Println(greatest3(nums3))
}
