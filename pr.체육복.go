package main

import "fmt"

func main(){
	
	fmt.Println(solution(5, []int{2, 4}, []int{1, 3, 5}))
	fmt.Println(solution(5, []int{2, 4}, []int{3}))
	fmt.Println(solution(3, []int{3}, []int{1}))
	fmt.Println(solution(13,[]int{1, 2, 5, 6, 10, 12, 13}, []int{2, 3, 4, 5, 7, 8, 9, 10, 11, 12}))
	fmt.Println(solution(8,[]int{5,6,7,8},[]int{4,7}))
	fmt.Println(solution(7,[]int{2,4,7},[]int{1,3,5}))
}


func solution(n int, lost []int, reserve []int) int {

	reserveCount, result := 0,0
	m := make(map[int]bool, n)

	// all student
	for i:= 1; i<=n; i++ {
		m[i] = true
	}

	// lost student
	for _, v := range lost {
		m[v] = false
	}

	// possible borrow count
	for _,v := range reserve {
		if m[v] == false {
			m[v] = true
			continue
		}

		if m[v] == true {
			reserveCount++
		}
	}

	for _, v := range m {
		if v == true {
			result++
		}
	}

	if result + reserveCount > n {
		return n
	}

	return result + reserveCount
}
