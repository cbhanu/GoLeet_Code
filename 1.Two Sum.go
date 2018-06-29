package main

import "fmt"

/* My Brute Force approach*/
func twoSum(nums []int, target int) []int {
	//var ans []int
	for i, first := range nums{
		for j:=i+1;j<len(nums);j++{   // or for j, second := range nums[1:] , return j+1 instead
			if first + nums[j] == target{
				//return append(ans, i, j)
				return []int{i,j}
			}
		}

	}
	return nil
}

/* The following is the best solution in terms of runtime for Go*/

func twoSum_best(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if _, ok := m[comp]; ok {
			return []int {m[comp], i}
		}
		m[nums[i]] = i
	}
	return nil
}


func main() {
	input:= []int{2,8,7,11}
	//twoSum_best(input,9)
	fmt.Printf("Answer : %d", twoSum(input,9))
}
