package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	fmt.Println(solution([]int{3, 76, 24}))
// }

// func solution(emergency []int) []int {

// 	arr := make([]int, len(emergency))
// 	copy(arr, emergency)

// 	m := make(map[int]int)

// 	for i, v := range emergency {
// 		m[v] = i
// 	}

// 	sort.Sort(sort.Reverse(sort.IntSlice(emergency)))

// 	for i, v := range emergency {
// 		m[v] = i
// 	}

// 	result := make([]int, len(emergency))
// 	for i, v := range arr {
// 		result[i] = m[v] + 1
// 	}

// 	return result
// }
