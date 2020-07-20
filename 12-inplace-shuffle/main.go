package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func shuffleInPlace(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		pos := getRand(i, len(arr)-1)
		swap(arr, i, pos)
	}

	return arr
}

func main() {
	arr := []int{1, 5, 7, 9, 2, 8}
	fmt.Println(shuffleInPlace(arr))
}
