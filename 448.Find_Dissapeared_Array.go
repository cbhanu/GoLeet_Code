package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {

	result := make([]int, len(nums))

	for _, num:= range nums{
		result[num-1] = 1
	}
	for j:= range result{
		if result[j]!=1{
			result = append(result,j+1)
		}
	}
	return result[len(nums):]
}

func main() {
	ip := []int{4,3,2,7,8,2,3,1}

	fmt.Println("\nMIA : ", findDisappearedNumbers(ip))
}
