package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main(){

// 	dp := make([]int, 11)

// 	dp[0] = 0
// 	dp[1] = 1
// 	dp[2] = 2
// 	dp[3] = 4

// 	for i,_ := range dp {

// 		if i > 3 {
// 			dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
// 		}
// 	}

// 	for _, v := range input() {
// 		fmt.Println(dp[v])
// 	}

// }

// func input() []int{
// 	var input string
// 	fmt.Scanln(&input)
// 	len, _ := strconv.Atoi(input)

// 	nums := make([]int, len)

// 	for i:=0; i< len; i++ {
// 		scn := bufio.NewScanner(os.Stdin)
// 		scn.Scan()

// 		v, _ := strconv.Atoi(strings.Split(scn.Text(), " ")[0])
// 		nums[i] = v
// 	}

// 	return nums
// }
