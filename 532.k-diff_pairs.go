package main

import "fmt"

func findPairs(nums []int, k int) int {

	if k<0{
		return 0
	}
	myUnique := make(map[int]int)
	// keys will be all unique here
	for _,n:= range nums{
		myUnique[n]+=1
		}
	fmt.Println(myUnique)

	count := 0
	// traverse they keys
	for ip,_ := range myUnique{
		//fmt.Print(ip)
		check := myUnique[ip+k]

		if check==0 || (check==1 && k==0){
			continue
		}
		count++
	}
	return count
}
func main() {
	ip := []int{1,2,3,4,1,5}

	fmt.Println("No of pairs: ", findPairs(ip,2))
}
