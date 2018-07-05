// Pretty much straightforward solution
package main

import "fmt"

func searchInsert(nums []int, target int) int {

	for i:=0;i<len(nums);i++{
		//fmt.Println("val: ",nums[i])
		if  target<= nums[i]{
			return i
		}
	}
	return len(nums)
}

func main() {
	val_ip := []int{1,3,5,6}

	lenn := searchInsert(val_ip, 7 )
	fmt.Println("Final answer: " , lenn)

}
