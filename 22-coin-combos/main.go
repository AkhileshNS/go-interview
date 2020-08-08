package main

import (
	"fmt"
)

// Recursive
/* func getCombos(res *map[string][]int, total int, coins, combo []int, cache *map[string]bool) {
	currTotal := 0
	comboStr := ""
	for i, count := range combo {
		currTotal += (coins[i] * count)
		comboStr += strconv.Itoa(count)
	}
	(*cache)[comboStr] = true
	if currTotal == total {
		(*res)[comboStr] = combo
		return
	} else if currTotal > total {
		return
	}

	for i := range coins {
		newCombo := make([]int, len(combo))
		copy(newCombo, combo)
		newCombo[i]++
		newComboStr := ""
		for _, count := range newCombo {
			newComboStr += strconv.Itoa(count)
        }
        if
		getCombos(res, total, coins, newCombo, cache)
	}
}

func getCoins(coins []int, total int) []string {
	res := map[string][]int{}
	combo := []int{}

	for i := 0; i < len(coins); i++ {
		combo = append(combo, 0)
	}

	getCombos(&res, total, coins, combo, &map[string]bool{})
	converted := []string{}

	for _, val := range res {
		str := ""

		for i, count := range val {
			for j := 0; j < count; j++ {
				str += strconv.Itoa(coins[i]) + "+"
			}
		}

		if len(str) > 0 {
			converted = append(converted, str[:len(str)-1])
		}
	}

	return converted
} */

func getCoins(coins []int, total int) int {
	// Each index in ways represents the amount to reach
	// The value at each index represents the number of ways to reach (index) amount
	ways := []int{0}
	for i := 0; i < total; i++ {
		ways = append(ways, 0)
	}

	for _, coin := range coins {
		for i := coin; i <= total; i++ {
			remainder := i - coin

			if remainder == 0 {
				ways[i]++
			} else if remainder > 0 {
				ways[i] += ways[remainder]
			}
		}
	}

	return ways[total]
}

func main() {
	denominations := []int{1, 3, 5}
	total := 5
	fmt.Println(getCoins(denominations, total))
}
