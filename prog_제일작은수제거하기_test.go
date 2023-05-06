package main

// func solution(arr []int) []int {

// 	if len(arr) == 1 {
// 		return []int{-1}
// 	}

// 	min, idx := 99999, -1
// 	for i, v := range arr {

// 		if min > v {
// 			min = v
// 			idx = i
// 		}
// 	}

// 	result := []int{}
// 	for i, v := range arr {
// 		if i != idx {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func TestSolution(t *testing.T) {
// 	v1 := solution([]int{4, 3, 2, 1})
// 	v2 := solution([]int{10})

// 	assert.Equal(t, v1, []int{4, 3, 2})
// 	assert.Equal(t, v2, []int{-1})
// }
