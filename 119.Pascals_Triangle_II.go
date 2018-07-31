package main

import (
	"fmt"
)

//func generate(numRows int) [][]int
func getRow(rowIndex int) []int {

	rowIndex +=1

	// To do : Optimize the algorithm to take O(k) space only, where we are returning kth index row starting from zero.

	//  ensure that numRows>0
	if rowIndex==0{
		return []int{}
	}
	// initializing the entire structure with ones
		// Kinda using O(k^2) space , 1 + 2 + 3 + ... + k = k(k+1)/2
	result := make([][]int,rowIndex)//[][]int{}
	for i:=0;i<rowIndex;i++{
		result[i] = make([]int,i+1)
		for j:= range result[i]{
			result[i][j] = 1
		}
	}
	// value of nth element in a row is the addition of the (n-1)th and nth element of previous row
	for p := 2; p < rowIndex; p++{
		for q :=1; q < p; q++{
			result[p][q] = result[p-1][q-1] + result[p-1][q]
		}
	}

	return result[rowIndex-1]

}


func main() {

	ip := 3 // i.e. return 4th row

	answer := getRow(ip)

	fmt.Println(answer)

}
