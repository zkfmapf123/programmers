package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var MONTH = []int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
// var WEEK = []string{"FRI", "SAT", "SUN", "MON", "TUE", "WED", "THU"}

// func solution(a int, b int) string {

// 	sum := 0
// 	for _, v := range MONTH[:a+1] {
// 		sum += v
// 	}

// 	sum -= (MONTH[a] - b)
// 	sum %= len(WEEK)

// 	if sum == 0 {
// 		return WEEK[len(WEEK)-1]
// 	}

// 	return WEEK[sum-1]
// }

// func TestSolution(t *testing.T) {
// 	v1 := solution(5, 24)
// 	assert.Equal(t, v1, "TUE")
// }
