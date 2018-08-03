package main

import (
	"fmt"
	"math"
)
func myMax(a,b int)int{
	if a<b{
		return b
	}else{
		return a
	}
}
func thirdMax(nums []int) int {

	if len(nums)<3{
		return myMax(nums[0],nums[1])
	}

	// Solution has to have O(n) time
	first ,second, third := -math.MaxInt64,-math.MaxInt64,-math.MaxInt64 //nums[len(nums)-1],nums[len(nums)-1],nums[len(nums)-1]

	for i:= range(nums){

		if nums[i]!= first && nums[i]!= second && nums[i]!= third{ // better way to do in Go ?

			if nums[i] >= first{
				third = second
				second = first
				first = nums[i]
			}else if nums[i]<first && nums[i]>second{
				third = second
				second = nums[i]

			}else if nums[i]<second && nums[i]>third{
				third = nums[i]
			}
			fmt.Println(first,second,third)
		}
	}

	if third == -math.MaxInt64{
		return first
	}
	return third
}
func main() {
	ip := []int{1,2,-2147483648}//1,2,2,5,3,5}//{1,1,2}//{2,4,1,3,7, 5}

	fmt.Println("Answer",thirdMax(ip))
}
