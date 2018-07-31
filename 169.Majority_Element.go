package main

import "fmt"

func majorityElement(nums []int) int {

	count:= 0
	if len(nums)%2 == 0{
		count = len(nums)/2
	}else{
		count = (len(nums)-1)/2
	}
	// make a map of value and its count
	m := make(map[int]int)
	for i := range(nums){
		m[nums[i]] += 1
		if m[nums[i]] > count{
			return nums[i]
		}
	}
	return 0

}
func main() {
	ip := []int {3,2,3,3}

	ans := majorityElement(ip)

	fmt.Println(ans)
}
