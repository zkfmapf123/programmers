package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func solution(x int, n int) []int64 {

// 	arr := make([]int64, n)
// 	arr[0] = int64(x)

// 	for i := 1; i < n; i++ {
// 		arr[i] = arr[i-1] + int64(x)
// 	}

// 	return arr
// }

// func TestSolution(t *testing.T) {
// 	assert.Equal(t, solution(2, 5), []int64{2, 4, 6, 8, 10})
// 	assert.Equal(t, solution(4, 3), []int64{4, 8, 12})
// 	assert.Equal(t, solution(-4, 2), []int64{-4, -8})
// }
