package main

import "fmt"

// FAILED because GoLang's objects don't have order

func mergeSortedArrays(arrays ...[]int) []int {
	res := []int{}
	holder := make(map[int]int)

	for _, arr := range arrays {
		for _, item := range arr {
			if _, ok := holder[item]; ok {
				holder[item]++
			} else {
				holder[item] = 1
			}
		}
	}

	for key, val := range holder {
		for i := 0; i < val; i++ {
			res = append(res, key)
		}
	}

	return res
}

func main() {
	myArray := []int{3, 4, 6, 10, 11, 15}
	alicesArray := []int{1, 5, 8, 12, 14, 19}
	fmt.Print(mergeSortedArrays(myArray, alicesArray))
}
