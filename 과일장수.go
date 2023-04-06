// package main

// import (
// 	"fmt"
// 	"math"
// 	"sort"
// )

// func main(){
// 	fmt.Println(solution(3,4,[]int{1,2,3,1,2,3,1}))
// 	fmt.Println(solution(4,3,[]int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}))

// }

// func solution(k int, m int, score []int) int {

// 	sum := 0

// 	// O(n)
// 	sort.Slice(score, func(i , j int) bool {
// 		return score[i] > score[j]
// 	})

// 	// O(n)
// 	for i:=0; i<=len(score); i+=m {

// 		if i+m > len(score) {
// 			break
// 		}

// 		max := math.Min(float64(k), float64(score[i:i+m][m-1]))
// 		sum += int(max) * m * 1
// 	}

// 	return sum
// }