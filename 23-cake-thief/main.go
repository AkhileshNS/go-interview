package main

import (
	"fmt"
	"math"
	"sort"
)

type cake struct {
	w int
	v int
}

func (c *cake) valByWeight() float64 {
	return float64(c.v) / float64(c.w)
}

type byValWeight []cake

func (a byValWeight) Len() int           { return len(a) }
func (a byValWeight) Less(i, j int) bool { return a[i].valByWeight() > a[j].valByWeight() }
func (a byValWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func maxDuffelBag(cakes []cake, capacity int) int {
	sortedCakes := make([]cake, len(cakes))
	copy(sortedCakes, cakes)
	sort.Sort(byValWeight(sortedCakes))

	fmt.Println(sortedCakes)

	remaining := capacity
	maxValue := 0

	for j := range sortedCakes {
		if remaining >= sortedCakes[j].w {
			count := int(math.Floor(float64(remaining) / float64(sortedCakes[j].w)))
			maxValue += (count * sortedCakes[j].v)
			remaining -= (count * sortedCakes[j].w)
			fmt.Println(maxValue, remaining, count)
		}
	}

	return maxValue
}

func main() {
	cakes := []cake{cake{w: 7, v: 160}, cake{w: 3, v: 90}, cake{w: 2, v: 15}}

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(maxDuffelBag(cakes, 20))
}
