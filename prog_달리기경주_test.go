package main

// import (
// 	"fmt"
// 	"testing"
// )

// func solution(players []string, callings []string) []string {

// 	m := make(map[string]int, len(players))
// 	for i, v := range players {
// 		m[v] = i
// 	}

// 	// callings
// 	for _, v := range callings {
// 		// 선두를 앞지른다
// 		lastOrder, lastOrderOk := m[v]

// 		// 이미 선두라면?
// 		if lastOrder == 0 {
// 			continue
// 		}

// 		if lastOrderOk {
// 			lastOrder--
// 			m[v] = lastOrder
// 		}

// 		// 해당 선두는 앞자리를 내준다
// 		firstOrder, firstOrderOk := m[players[lastOrder]]

// 		if firstOrderOk {
// 			firstOrder++
// 			m[players[lastOrder]] = firstOrder
// 		}

// 		players[firstOrder], players[lastOrder] = players[lastOrder], players[firstOrder]
// 	}

// 	return players
// }

// // go test -run ./prog_달리기경주_test.go
// func TestSolution(t *testing.T) {

// 	v := solution([]string{"mumu", "soe", "poe", "kai", "mine"}, []string{"kai", "kai", "mine", "mine"})

// 	fmt.Println(v)

// 	// assert.Equal(t, v, []string{"mumu", "kai", "mine", "soe", "poe"})

// }
