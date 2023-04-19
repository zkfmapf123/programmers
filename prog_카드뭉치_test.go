package main

// func solution(cards1 []string, cards2 []string, goal []string) string {

// 	isFirstExists, isSecondExists := true, true

// 	for _, str := range goal {

// 		isFirstExists, cards1 = getReduceCard(cards1, str)
// 		isSecondExists, cards2 = getReduceCard(cards2, str)

// 		if !isFirstExists && !isSecondExists {
// 			return "No"
// 		}
// 	}

// 	return "Yes"

// }

// func getReduceCard(cards []string, goal string) (bool, []string) {

// 	if len(cards) == 0 {
// 		return false, cards
// 	}

// 	if cards[0] == goal {

// 		return true, cards[1:]
// 	}

// 	return false, cards
// }

// func TestSolution(t *testing.T) {

// 	v1 := solution([]string{"i", "drink", "water"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"})
// 	v2 := solution([]string{"i", "water", "drink"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"})

// 	assert.Equal(t, v1, "Yes")
// 	assert.Equal(t, v2, "No")
// }
