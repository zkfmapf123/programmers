package main

// func solution(arr []int) int {

// 	num := arr[0]
// 	for i, _ := range arr {

// 		if i == 0 {
// 			continue
// 		}
// 		num = lcm(num, arr[i])
// 	}

// 	return num
// }

// // 최대 공배수
// func lcm(a, b int) int {
// 	return a * b / gcd(a, b)
// }

// // 최소 공약수 구하기
// func gcd(a, b int) int {

// 	if b == 0 {
// 		return a
// 	}

// 	return gcd(b, a%b)
// }

// func TestSolution(t *testing.T) {
// 	t1 := solution([]int{2, 6, 8, 14})

// 	assert.Equal(t, t1, 168)
// }
