package main

// 재귀로 푸니까 시간초과 남
// func solution(n int) int {

// 	arr := make([]int, n+1)

// 	for i := 0; i < n; i++ {
// 		if i <= 1 {
// 			arr[i] = 1 % 1234567
// 			continue
// 		}

// 		arr[i] = (arr[i-1] + arr[i-2]) % 1234567
// 	}

// 	return arr[n-1]
// }

// func TestSolution(t *testing.T) {

// 	t1 := solution(3)
// 	t2 := solution(5)

// 	assert.Equal(t, t1, 2)
// 	assert.Equal(t, t2, 5)
// }
