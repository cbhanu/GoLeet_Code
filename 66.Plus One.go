package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	for i:=len(digits)-1;i>=0;i--{
		if (digits[i] + 1) == 10{
			digits[i] = 0
			if i==0{
				// all digits(whatever the length) were nine
				digits = append(digits,0)
				digits[0]=1
			}
		}else{
		digits[i] += 1
		return digits
		}
	}
	return digits
}

func main() {
	input := []int{9,9,9}
	ans := []int{}
	ans = plusOne(input)
	fmt.Print(ans)
}
