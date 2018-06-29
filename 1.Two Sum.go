/*
The earlier easier codes contain some over detailed comments for a comprehensive read
 */

package main

import "fmt"

/* My Brute Force approach*/
func twoSum(nums []int, target int) []int {
	//var ans []int
	for i, first := range nums{
		for j, second := range nums[1:]{
			if first + second == target{
				//return append(ans, i, j)
				return []int{i,j+1}
			}
		}

	}
	return nil
}

/* The following is the best solution in terms of runtime for Go*/

func twoSum_best(nums []int, target int) []int {
	m := make(map[int]int)
	for i, curr := range nums{
		comp := target - curr
		_, ok := m[comp]
		if ok {
			return []int {m[comp], i}
			// When ok is true, the comparison value, comp is going to be the first number, since curr is the second number
		}
		m[nums[i]] = i // current difference value is the key , index is the dictionary value
	}
	return nil
}

func main() {
	input:= []int{2,8,7,11}
	//twoSum(input,9)
	fmt.Printf("\nAnswer : %d", twoSum_best(input,9))
}