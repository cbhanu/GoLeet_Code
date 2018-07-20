package main

import "fmt"

func maxSubArray(nums []int) int {

	sum  := make([]int,0)
	running_sum:=0
	max_ans:=0
	max:=0
	for i:=0 ; i<len(nums);i++{


		//fmt.Println(i)
		sum = append(sum,nums[i])
		running_sum = Max(sum[i],running_sum+sum[i])
		max_ans = Max(max_ans,running_sum)
		for j:=i+1; j<len(nums)-1;j++{
			sum[i] = sum[i] + nums[j]
			max = Max(max,sum[i])
		}

	}
	// return the max out of the sums
	return max//max_ans//max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	ip := []int {-2,1,-3,4,-1,2,1,-5,4}

	fmt.Println(maxSubArray(ip))
}
