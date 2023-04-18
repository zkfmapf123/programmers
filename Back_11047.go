package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var (
// 	r = bufio.NewReader(os.Stdin)
// )

// /*
// 첫째 줄에 N과 K가 주어진다. (1 ≤ N ≤ 10, 1 ≤ K ≤ 100,000,000)
// 둘째 줄부터 N개의 줄에 동전의 가치 Ai가 오름차순으로 주어진다. (1 ≤ Ai ≤ 1,000,000, A1 = 1, i ≥ 2인 경우에 Ai는 Ai-1의 배수) -> Greedy
// */
// func main() {

// 	result := 0
// 	n1, n2 := getN()

// 	arr := make([]int, n1)
// 	for i := 0; i < n1; i++ {
// 		fmt.Scanf("%d\n", &arr[i])
// 	}

// 	for i := n1 - 1; i >= 0; i-- {
// 		if arr[i] > n2 {
// 			continue
// 		}

// 		result += n2 / arr[i]
// 		n2 %= arr[i]

// 		if n2 == 0 {
// 			break
// 		}
// 	}

// 	fmt.Println(result)

// }

// func getN() (int, int) {

// 	var n1, n2 int
// 	fmt.Scanf("%d %d\n", &n1, &n2)

// 	return n1, n2
// }
