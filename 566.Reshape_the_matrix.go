package main

import "fmt"

func createMatrix(r,c int)[][]int{
	m := make([][]int,r)
	for i := range m {
		m[i] = make([]int, c)
		for j := 0; j < c; j++ {
			m[i][j] = 0

		}
	}
	return m
}
func matrixReshape(nums [][]int, r int, c int) [][]int {

	// first check if it can be reshaped, otherwise return original
	if len(nums)*len(nums[0])!=r*c{
		fmt.Println("Invalid reshape")
		return nums
	}
	rPo :=0
	cPo :=0
	result := createMatrix(r,c)
	for i:=0 ;i< len(nums);i++{
		for j:=0;j<len(nums[0]);j++{
			val := nums[i][j]
			result[rPo][cPo] = val
			cPo++
			if cPo==c{
				cPo =0
				rPo++
			}
		}
	}

	return result
}
func main() {
	row := 3
	col := 2
	ip := make([][]int,row) //	fmt.Println(ip)
	val:=1
	for i := range ip {
		ip[i] = make([]int, col)
		for j := 0; j < col; j++ {
			ip[i][j] = val
			val++
		}

	}
	fmt.Println("Input: ",ip)
	fmt.Println("result: ",matrixReshape(ip,col,row))
}