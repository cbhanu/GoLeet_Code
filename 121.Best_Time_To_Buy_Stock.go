package main

import (
	"fmt"
)

func maxProfit(prices []int) int {

	// 1) brute force - search through all elements and their differences - O(n^2) and choose the max
	// 2 ) Keep a track of the maxProfit as of now, even if it wasn't bought at minimum price
	minPrice := prices[0]
	maxProfit:= 0

	for i:=1 ; i<len(prices);i++{
		if prices[i] < minPrice{
			minPrice = prices[i]

		}else if prices[i] - minPrice > maxProfit{
			maxProfit = prices[i]-minPrice
		}
	}
	return maxProfit
}

func main() {
	ip := []int{2,4,1}//{7,2,1,5,3,6,4}

	my_profit := maxProfit(ip)

	fmt.Println(my_profit)
}
