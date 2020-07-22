package main

import (
	"fmt"
	"math"
)

func genFib(n int) float64 {
	return math.Round(math.Pow(math.Phi, float64(n)) / math.Sqrt(5))
}

func main() {
	fmt.Println(genFib(9))
}
