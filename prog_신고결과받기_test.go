package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type reportAlertParams struct {
	v   int
	arr []string
}

func solution(id_list []string, report []string, k int) []int {

	// 1. 중복된 report를 검사한다
	reportMap := make(map[string]bool)
	for _, v := range report {
		reportMap[v] = true
	}

	// 2. 신고자들의 대한 map을 만든다
	alertMap := make(map[string]int)
	for _, v := range id_list {
		alertMap[v] = 0
	}

	// 3. 신고당한 사람들의 대한 map을 만든다
	reportAlertMap := make(map[string]reportAlertParams)

	// 4. 신고 당한 사람들을 검사한다
	for key, _ := range reportMap {
		v := strings.Split(key, " ")
		alertMan, reportMan := strings.Trim(v[0], " "), strings.Trim(v[1], " ")

		value, isExist := reportAlertMap[reportMan]

		// 값이 존재하지 않는다면
		if !isExist {

			value.v = 1
			value.arr = append(value.arr, alertMan)
			reportAlertMap[reportMan] = value
			continue
		}

		// 값이 존재한다면 && 신고횟수가 총족된다면
		if value.v >= k {

		}

	}

	return []int{}
}

func TestSolution(t *testing.T) {

	// assert.Equal(t, solution([]string{"muzi", "frodo", "apeach", "neo"}, []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}, 2), []int{2, 1, 1, 0})
	assert.Equal(t, solution([]string{"con", "ryan"}, []string{"ryan con", "ryan con", "ryan con", "ryan con"}, 3), []int{0, 0})
}
