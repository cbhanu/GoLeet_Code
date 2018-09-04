package main

import "fmt"


func Fibo( n int)int{
	memo := make([]int,n+1)
	memo[0] , memo[1] = 0,1

	for i:= range memo[2:]{
		memo[i+2] = memo[i+1] + memo[i]

	}
	fmt.Println(memo)
	return memo[n]
}
func main() {
	ip := 5

	fmt.Println(Fibo(ip))
}
