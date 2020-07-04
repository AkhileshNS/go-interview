package main

import "fmt"

func mergeSortedArrays(a1, a2 []int) []int {
	res := []int{}
	i := 0
	j := 0

	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			res = append(res, a1[i])
			i++
		} else {
			res = append(res, a2[j])
			j++
		}
	}

	if i < len(a1) {
		for i < len(a1) {
			res = append(res, a1[i])
			i++
		}
	}

	if j < len(a2) {
		for j < len(a2) {
			res = append(res, a2[j])
			j++
		}
	}

	return res
}

func main() {
	myArray := []int{3, 4, 6, 10, 11, 15}
	alicesArray := []int{1, 5, 8, 12, 14, 19}
	fmt.Print(mergeSortedArrays(myArray, alicesArray))
}
