package main

import "fmt"

func findMaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := prices[1] - prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < minPrice {
			minPrice = prices[i]
		} else {
			currentProfit := prices[i] - minPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}

func main() {
	prices1 := []int{12, 10, 9, 8, 7, 6}
	prices2 := []int{10, 7, 5, 8, 11, 9}
	prices3 := []int{10, 10, 10, 10, 10, 10}

	fmt.Println(findMaxProfit(prices1))
	fmt.Println(findMaxProfit(prices2))
	fmt.Println(findMaxProfit(prices3))
}
