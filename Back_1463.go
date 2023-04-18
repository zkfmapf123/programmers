package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main(){
// 	n := input()

// 	dp := make([]int, n+1)

// 	// mapping
// 	dp[0] = 0
// 	dp[1] = 0

// 	for i:= 2; i<=n; i++ {

// 		dp[i] = dp[i-1] + 1

// 		if i % 2 == 0 {
// 			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i / 2] + 1)))
// 		}

// 		if i % 3 == 0 {
// 			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i / 3] + 1)))
// 		}
// 	}

// 	fmt.Println(dp[n])
// }

// func input() int{
// 	s := bufio.NewScanner(os.Stdin)
// 	s.Scan()

// 	input, _ := strconv.Atoi(strings.Split(s.Text(), " ")[0])
// 	return input
// }
