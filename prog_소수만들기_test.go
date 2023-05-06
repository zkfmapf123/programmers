package main

// func solution(nums []int) int {

// 	result := 0

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				sum := nums[i] + nums[j] + nums[k]
// 				if isPrime(sum) {
// 					result++
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

// // 소수 판별
// func isPrime(num int) bool {
// 	if num < 2 {
// 		return true
// 	}

// 	for i := 2; i < num; i++ {

// 		if num%i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func TestSolution(t *testing.T) {
// 	v1 := solution([]int{1, 2, 3, 4})
// 	// v2 := solution([]int{1, 2, 7, 6, 4})

// 	assert.Equal(t, v1, 1)
// 	// assert.Equal(t, v2, 4)
// }
