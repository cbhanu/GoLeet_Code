/* This solution checks each value and increments only if an element was not removed otherwise checks in the same index again
Run Time : 0 ms (Well.. :\*/

package main
import "fmt"

func removeElement(nums []int, val int) int {
	i:=0
	LOOP:
	if i<len(nums){
		if nums[i]==val {
			nums = append(nums[:i], nums[i+1:]...)
				fmt.Println(nums)
				goto LOOP

			}else{
				i++
				goto LOOP
			}
		fmt.Println(nums)
		}

	return len(nums)
}

func main() {
	//val_ip := []int{0,0,0,1,1,1,2,2,3,3,4,5,5, 6, 7, 7, 8}
	val_ip := []int{1,0,8,0,1,3,1,1,2,5,0}
	fmt.Println(val_ip)
	lenn := removeElement(val_ip,1)
	fmt.Println("Final answer: ")
	for i:= 0; i < lenn; i++{
		fmt.Print(val_ip[i]);
	}
	fmt.Println("\nLength is:",lenn)

}

/* Output:
[1 0 8 0 1 3 1 1 2 5 0]
[0 8 0 1 3 1 1 2 5 0]
[0 8 0 3 1 1 2 5 0]
[0 8 0 3 1 2 5 0]
[0 8 0 3 2 5 0]
Final answer:
0803250
Length is: 7
 */
