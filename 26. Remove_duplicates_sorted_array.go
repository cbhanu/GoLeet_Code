/* This solution of mine was the fastest as of 29th June, feel free to suggest improvements!*/
/* This solution compares 1st value with second, then 3rd etc until it ceases to be the same value and slices the array
 from the last similar element*/

package main
import "fmt"

func removeDuplicates(nums []int) int {

	for i:=0;i<len(nums)-1;i++{
		cnt:=0
		//fmt.Println(nums[i],nums[i+1])
		if nums[i+1] == nums[i] {
			cnt+=1
			for j:=i+1;j<len(nums)-1;j++{
				if nums[j+1] != nums[i] {
					break
				}else{cnt += 1}

		}
		fmt.Println(nums)
		nums = append(nums[:i], nums[i+cnt:]...)
		}
	}
	return len(nums)
}

func main() {
    val_ip := []int{0,0,0,1,1,1,2,2,3,3,4,5,5, 6, 7, 7, 8}

	lenn := removeDuplicates(val_ip)
	fmt.Println("Final answer: ")
	for i:= 0; i < lenn; i++{
		fmt.Print(val_ip[i]);
	}
	fmt.Println("\nLength is:",lenn)

}

/* Output:
[0 0 0 1 1 1 2 2 3 3 4 5 5 6 7 7 8]
[0 1 1 1 2 2 3 3 4 5 5 6 7 7 8]
[0 1 2 2 3 3 4 5 5 6 7 7 8]
[0 1 2 3 3 4 5 5 6 7 7 8]
[0 1 2 3 4 5 5 6 7 7 8]
[0 1 2 3 4 5 6 7 7 8]
Final answer:
012345678
Length is: 9

 */
