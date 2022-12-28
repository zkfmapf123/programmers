// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func input() int{
// 	s := bufio.NewScanner(os.Stdin)
// 	s.Scan()

// 	input, _ := strconv.Atoi(strings.Split(s.Text(), " ")[0])
// 	return input
// }

// func main(){
// 	n := input()

// 	dp := make([]int64, n+1)
// 	dp[0] = 1
// 	dp[1] = 2

// 	for i:= 2; i <= n; i++ {
// 		dp[i] = (dp[i-1] + dp[i-2]) % 10007
		
// 	}

// 	fmt.Println(dp[n-1])
// }