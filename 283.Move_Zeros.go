package main

import "fmt"

func moveZeroes(nums []int) {
	cnt:=0
	i := 0
	LOOP:if i < len(nums) {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			cnt++
			fmt.Println("a",nums)
			goto LOOP

		} else {
			i++
			fmt.Println("b",nums)
			goto LOOP
		}

	}
	fmt.Println("removed",nums)
for j:=0;j<cnt;j++{
	//fmt.Println(nums,j)
	nums=append(nums,0)
	}
}
func main() {
	ip  := []int{0,0,1}//0,1,0,2,0,13}
	moveZeroes(ip)

	fmt.Println("ans",ip)
}
