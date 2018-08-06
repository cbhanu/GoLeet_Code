package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {

	prev_result := 0
	cnt := 0

	for i:= range nums{
		if nums[i]==1{
			cnt++
			if cnt > prev_result{
				prev_result = cnt
			}
		}else{
			cnt = 0
		}
	}
	return prev_result
}

func main() {
	ip := []int{1,1,1}//{1,1,1, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1}

	fmt.Println(findMaxConsecutiveOnes(ip))
}
