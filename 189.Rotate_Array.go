package main

import "fmt"

func rotate(nums []int, k int)  {

	// 1) O(k) space and O(n) timee
	nums1 := []int{}
	nums1 = append(nums1,nums[len(nums)-k:]...)
	nums1 = append(nums1,nums[:len(nums)-k]... )

	copy (nums, nums1)
}
func main() {
	ip := []int{1,2,3,4,5,6,7}

	rotate(ip, 3)

	fmt.Println(ip)

}
