package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func solution(n int64) []int {

	str := strings.Split(strconv.Itoa(int(n)), "")
	arr := []int{}

	for i := len(str) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(str[i])
		arr = append(arr, num)
	}

	return arr
}

func TestSolution(t *testing.T) {

	t1 := solution(12345)
	assert.Equal(t, t1, []int{5, 4, 3, 2, 1})
}
