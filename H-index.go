package main

// func main(){
// 	// r1 := solution([]int{3,0,6,1,5})
// 	// r2 := solution([]int{5,5,5,5,5})
// 	// r3 := solution([]int{0,1,1})
// 	// r4 := solution([]int{5,4,3,8,10})
// 	r5 := solution([]int{5,3,3,8,10})

// 	fmt.Println(r5)
// }

// func solution(citations []int) int {

// 	sort.Ints(citations)

// 	for i := 0; i < len(citations); i++ {
// 		countOfOver, countOfDown := 0, 0

// 		for _,v := range citations {
// 			if i >= v {
// 				countOfOver += 1
// 			}

// 			if i <= v {
// 				countOfDown += 1
// 			}
// 		}

// 		fmt.Println(i , countOfDown, countOfOver)
// 	if countOfDown == countOfOver {
// 		return i
// 	}

// 	}

// 	return -1
// }
