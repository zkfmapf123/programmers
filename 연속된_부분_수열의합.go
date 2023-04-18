package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(solution([]int{1, 2, 3, 4, 5}, 7))
// 	// fmt.Println(solution([]int{1, 1, 1, 2, 3, 4, 5}, 5))
// 	// fmt.Println(solution([]int{2, 2, 2, 2, 2}, 6))
// }

// func solution(sequence []int, k int) []int {

// 	answer := make([]int, len(sequence))
// 	forth, sum, _len := 0, 0, 1000000

// 	for i := 0; i < len(sequence); i++ {

// 		sum += sequence[i]
// 		if sum == k {
// 			if _len > i-forth {
// 				answer[0] = forth
// 				answer[1] = i
// 				_len = i - forth
// 			}
// 			sum -= sequence[forth]
// 			forth++
// 			continue
// 		}

// 		if sum > k {
// 			sum -= sequence[forth]
// 			forth++
// 			sum -= sequence[i]
// 			i--
// 		}
// 	}

// 	return []int{answer[0], answer[1]}
// }

// n log(n) 시간초과
// n 으로 풀어야 함
// func solution(sequence []int, k int) []int {

// 	isInit, arr := false, make([]int, 1)

// 	for i, _ := range sequence {
// 		r := iteratorSequence(sequence, k, i, []int{}, 0)

// 		if r == nil {
// 			continue
// 		}

// 		if !isInit {
// 			arr = r
// 			isInit = true
// 			continue
// 		}

// 		if len(arr) > len(r) {
// 			arr = r
// 			continue
// 		}

// 		if len(arr) == len(r) && arr[0] > r[0] {
// 			arr = r
// 			continue
// 		}

// 	}

// 	if len(arr) == 1 {
// 		return []int{arr[0], arr[0]}
// 	}

// 	if len(arr) >= 2 {
// 		return []int{arr[0], arr[len(arr)-1]}
// 	}

// 	return arr
// }

// func iteratorSequence(se []int, k int, idx int, re []int, sum int) []int {

// 	if sum == k {
// 		return re
// 	}

// 	// Failed
// 	if idx >= len(se) {
// 		return nil
// 	}

// 	_sum := sum + se[idx]
// 	re = append(re, idx)
// 	return iteratorSequence(se, k, idx+1, re, _sum)
// }
