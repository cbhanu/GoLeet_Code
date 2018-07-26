package main

import (
	"fmt"
)

func generate(numRows int) [][]int {

	//  ensure that numRows>0
	if numRows==0{
		return [][]int{}
	}

	// initializing the entire structure with ones
	result := make([][]int,numRows)//[][]int{}
	for i:=0;i<numRows;i++{
		result[i] = make([]int,i+1)
		for j:= range result[i]{
			result[i][j] = 1
		}
	} //special cases for numRows<=2, keep as it is

	// value of nth element in a row is the addition of the (n-1)th and nth element of previous row
	for p := 2; p < numRows; p++{
		for q :=1; q < p; q++{
			result[p][q] = result[p-1][q-1] + result[p-1][q]
		}
	} 	// if a row has 'n' elements, it needs to modify (n-2) elements

	return result

}

func main() {

	ip := 7

	answer := generate(ip)

	for i:= range answer{
		fmt.Println(answer[i])
	}

}
/*
OUTPUT:
[1]
[1 1]
[1 2 1]
[1 3 3 1]
[1 4 6 4 1]
[1 5 10 10 5 1]
[1 6 15 20 15 6 1]
 */