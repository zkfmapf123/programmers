package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// go test -run ./prog_덧칠하기_test.go
// func solution(n int, m int, section []int) int {

// 	sum, painted := 0, 0

// 	for _, v := range section {

// 		// 해당 구역이 페인트 되지 않았을 때
// 		if v >= painted {
// 			sum++
// 			painted = v + m
// 		}
// 	}

// 	return sum
// }

// func TestSolution(t *testing.T) {
// 	t1 := solution(8, 4, []int{2, 3, 6})
// 	t2 := solution(5, 4, []int{1, 3})
// 	t3 := solution(4, 1, []int{1, 2, 3, 4})

// 	assert.Equal(t, t1, 2)
// 	assert.Equal(t, t2, 1)
// 	assert.Equal(t, t3, 4)
// }
