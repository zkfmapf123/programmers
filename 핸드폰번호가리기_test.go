package main

// import (
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func solution(phone_number string) string {

// 	if len(phone_number) <= 4 {
// 		return phone_number
// 	}

// 	phoneArr := strings.Split(phone_number, "")
// 	for i := 0; i < len(phoneArr)-4; i++ {
// 		phoneArr[i] = "*"
// 	}

// 	return strings.Join(phoneArr, "")
// }

// func TestSolution(t *testing.T) {

// 	assert.Equal(t, solution("01033334444"), "*******4444")
// 	assert.Equal(t, solution("027778888"), "*****8888")
// }
