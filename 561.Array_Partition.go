package main

import (
	"sort"
	"fmt"
)

func arrayPairSum(nums []int) int {

	// first sort the array - O(nlogn)
	sort.Ints(nums)
	//fmt.Println(nums)
	sum := 0
	for i:=0;i<len(nums);i=i+2{
		sum += nums[i]
	}
	return sum
}

func main() {
	ip := []int{1,4,3,2}

	fmt.Println("result: ", arrayPairSum(ip))
}


/* JAVA SOLution in O(n) : from discussions

public class Solution {

	public int arrayPairSum(int[] nums) {
		int[] exist = new int[20001];
		for (int i = 0; i < nums.length; i++) {
			exist[nums[i] + 10000]++;
		}
		int sum = 0;
		boolean odd = true;
		for (int i = 0; i < exist.length; i++) {
			while (exist[i] > 0) {
				if (odd) {
					sum += i - 10000;
				}
				odd = !odd;
				exist[i]--;
			}
		}
		return sum;
	}

}

*/