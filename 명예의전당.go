package main

// import (
// 	"fmt"
// 	"math"
// 	"sort"
// )

// func main(){
// 	solution(3,[]int{10,100,20,150,1,100,200})
// 	solution(4,[]int{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000})
// 	// solution(9,[]int{10, 30, 40, 3, 0, 20, 4})
// }

// func solution(k int, score []int) []int {
// 	result := make([]int,len(score))

// 	arr := make([]int,k)
// 	for i,v := range score {

// 		if i < k {
// 			result[i] = score[0]
// 			arr[i] = v
// 		}else{
// 			result[i] = diffNum(arr, v)

// 		}
// 	}

// 	fmt.Println(result)
//     return result
// }

// func diffNum(score []int, value int) int{
// 	sort.Ints(score)
// 	max := int(math.Max(float64(score[0]), float64(value)))

// 	if score[0] <= max {
// 		score[0] = max
// 	}

// 	sort.Ints(score)
// 	return score[0]
// }
