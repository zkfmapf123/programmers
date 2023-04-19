package main

// func solution(keymap []string, targets []string) []int {

// 	touchPad := getTouchPad(keymap)

// 	ret := make([]int, len(targets))
// 	for i, target := range targets {
// 		ret[i] = getMinTouchPadNum(touchPad, target)
// 	}

// 	return ret

// }

// func getTouchPad(keymaps []string) map[string]int {

// 	m := make(map[string]int)
// 	for _, keys := range keymaps {

// 		for i, key := range keys {
// 			str := string(key)

// 			// touchPad가 없을 때
// 			if m[str] == 0 {
// 				m[str] = i + 1
// 				continue
// 			}

// 			m[str] = int(math.Min(float64(m[str]), float64(i+1)))
// 		}
// 	}

// 	return m
// }

// func getMinTouchPadNum(m map[string]int, target string) int {

// 	sum := 0
// 	for _, ru := range target {
// 		str := string(ru)

// 		if m[str] == 0 {
// 			return -1
// 		}

// 		sum += m[str]
// 	}

// 	return sum
// }

// func TestSolution(t *testing.T) {
// 	v1 := solution([]string{"ABACD", "BCEFD"}, []string{"ABCD", "AABB"})
// 	v2 := solution([]string{"AA"}, []string{"B"})
// 	v3 := solution([]string{"AGZ", "BSSS"}, []string{"ASA", "BGZ"})

// 	assert.Equal(t, v1, []int{9, 4})
// 	assert.Equal(t, v2, []int{-1})
// 	assert.Equal(t, v3, []int{4, 6})
// }
