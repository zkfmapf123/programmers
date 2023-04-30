package main

// func solution(s string) bool {

// 	if len(s) != 4 && len(s) != 6 {
// 		return false
// 	}

// 	for _, v := range s {

// 		r, _ := regexp.Compile("[^0-9]")
// 		isNotMatched := r.MatchString(string(v))

// 		if isNotMatched {
// 			return false
// 		}
// 	}

// 	return true
// }

// func TestSolution(t *testing.T) {
// 	t1 := solution("a234")
// 	t2 := solution("1234")

// 	assert.Equal(t, t1, false)
// 	assert.Equal(t, t2, true)
// }
