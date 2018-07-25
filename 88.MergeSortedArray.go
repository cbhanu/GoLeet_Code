package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {

	// first add elements in nums2 to the end of nums1, // then sort the entire nums1
	j:=0
	for i:=m;i<len(nums1);i++{
		nums1[i]=nums2[j]
		j++
	}
	// some sorting algorithm - QuickSort built-in
	sort.Ints(nums1)

	/* An inefficient implementation
	if m==0{
		 nums1 = append(nums1[:0],nums2...)
	     return
	}
	cnt:=0
	for j:=0;j<n;j++{
		if nums2[j] < nums1[m+cnt-1] {
			for i := 0; i < m+cnt; i++ {
				if nums2[j] < nums1[i] {
					//insert
					for ii := m + n - 1; ; ii-- {
						nums1[ii] = nums1[ii-1]
						if ii == i+1 {
							break
						}
					}
					nums1[i] = nums2[j]
					cnt++
					break
				}
			}
		}else{
			//append after all non zero elements of num1
			nums1[m+cnt]=nums2[j]
			cnt++
		}
	}
	// Fastest , better implementation
	index1 := m - 1
	index2 := n - 1
	index3 := m + n - 1
	for {
		if index1 < 0 || index2 < 0 {
			break
		}
		if nums1[index1] > nums2[index2] {
			nums1[index3] = nums1[index1]
			index1--
			index3--
		} else {
			nums1[index3] = nums2[index2]
			index2--
			index3--
		}
	}
	for ; index2 >= 0; {
		nums1[index3] = nums2[index2]
		index2--
		index3--
	}
	*/

}

func main() {

	ip1 := []int{4,0,0,0,0,0}
	m := 1
	ip2 := []int{1,2,3,5,6}
	n := len(ip2)

	merge(ip1,m,ip2,n)

	fmt.Println("Final answer")
	for i:=0;i<(len(ip1));i++{
		fmt.Println(ip1[i])
	}
}
