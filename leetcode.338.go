package main

// import (
// 	"fmt"
// )

// func main(){
// 	fmt.Println(countBits(2))
//  	fmt.Println(countBits(5))
// }

// /*
// 	Rule
// 	dp[0] = 0
// 	dp[1] = 01  = 1 + dp[1-1]
// 	dp[2] = 10  = 1 + dp[2-2]
// 	dp[3] = 001 = 1 + dp[3-2]
// 	dp[4] = 100 = 1 + dp[4-4]
// 	...
// 	dp[n] = n = 1 + dp[n - offset]
// */

// func countBits(n int) []int {
// 	dp, offset := make([]int, n+1), 1
// 	for i:=1; i<=n; i++{

// 		// offset => 1, 2, 4, 8 ...
// 		if offset * 2 == i {
// 			offset = i
// 		}

// 		dp[i] = 1+dp[i - offset]
// 	}

// 	return dp
// }
