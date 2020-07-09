package main

import "fmt"

func checkFCFS(takeOut, dineIn, served []int) bool {
	i := 0
	j := 0

	for _, order := range served {
		if order != dineIn[j] && order != takeOut[i] {
			return false
		}
		if order == takeOut[i] && i != len(takeOut)-1 {
			i++
		}
		if order == dineIn[j] && j != len(dineIn)-1 {
			j++
		}
	}

	return true
}

func main() {
	takeOut := []int{1, 3, 5}
	dineIn := []int{2, 4, 6}
	served := []int{1, 2, 4, 6, 5, 3}
	fmt.Println(checkFCFS(takeOut, dineIn, served))

	takeOut = []int{17, 8, 24}
	dineIn = []int{12, 19, 2}
	served = []int{17, 8, 12, 19, 24, 2}
	fmt.Println(checkFCFS(takeOut, dineIn, served))
}
