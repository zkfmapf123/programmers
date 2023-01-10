// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main(){
// 	v, _ := strconv.Atoi(input())

// 	dp := make([]int,v+1)
// 	dp[0] = 0
// 	dp[1] = 1

// 	if v <= 1 {
// 		fmt.Println(dp[v])
// 		return
// 	}

// 	for i := 2; i<=v; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}

// 	fmt.Print(dp[v])	
// }

// func input() string{
// 	scn := bufio.NewScanner(os.Stdin)

// 	scn.Scan()
// 	return strings.Split(scn.Text(), " ")[0]
// }