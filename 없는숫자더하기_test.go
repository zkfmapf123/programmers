package main

// func solution(numbers []int) int {
// 	sort.Ints(numbers)

// 	i, start, sum := 0, 0, 0
// 	for start < 10 {

// 		if i >= len(numbers) {
// 			i++
// 			sum += start
// 			start++
// 			continue
// 		}

// 		// not correct
// 		if start != numbers[i] {
// 			sum += start
// 			start++
// 			continue
// 		}

// 		// correct
// 		if start == numbers[i] {
// 			i++
// 			start++
// 			continue
// 		}
// 	}

// 	return sum
// }

// func TestSolution(t *testing.T) {

// 	assert.Equal(t, solution([]int{1, 2, 3, 4, 6, 7, 8, 0}), 14)
// 	assert.Equal(t, solution([]int{5, 8, 4, 0, 6, 7, 9}), 6)
// }
