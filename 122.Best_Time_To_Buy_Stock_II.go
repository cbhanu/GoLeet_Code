package main

import "fmt"

func maxProfitII(prices []int) int {

	if len(prices)<=1{
		return 0
	}
	result := 0
	for i := 1 ; i < len(prices) ;i++{

		if prices[i] > prices[i - 1]{
			result = result +  prices[i] - prices[i - 1]
		}

	}

	return result
}
func main() {

	ip := []int{2,1,2,0,1}//{2,4,1}//{6,1,3,2,4,7}//{7,2,1,5,3,6,4}//// // {7,1,5,3,6,4} //{3,2,6,5,0,3} //{1,4,2}

	my_profit := maxProfitII(ip)

	fmt.Println(my_profit)
}
