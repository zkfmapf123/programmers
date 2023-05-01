package main

// func solution(n int) int64 {

// 	arr := make([]int, n)

// 	for i := 0; i < n; i++ {

// 		if i <= 1 {
// 			arr[i] = (i + 1) % 1234567
// 			continue
// 		}

// 		arr[i] = (arr[i-1] + arr[i-2]) % 1234567

// 	}

// 	return int64(arr[n-1])
// }

/*
	2. 재귀 방법 (시간초과)
*/
// func recur(n int) int {

// 	if n == 1 || n == 0 {
// 		return 1
// 	}

// 	return recur(n-1) + recur(n-2)
// }

/*
	1. DFS 방법 (시간초과)
*/
// func dfs(len, sum int) {

// 	if sum > len {
// 		return
// 	}

// 	if len == sum {
// 		ans++
// 		return
// 	}

// 	dfs(len, sum+1)
// 	dfs(len, sum+2)
// }

// func TestSolution(t *testing.T) {

// 	t1 := solution(4)
// 	t2 := solution(3)
// 	t3 := solution(1)
// 	t4 := solution(2)

// 	assert.Equal(t, t1, int64(5))
// 	assert.Equal(t, t2, int64(3))
// 	assert.Equal(t, t3, int64(1))
// 	assert.Equal(t, t4, int64(2))
// }
