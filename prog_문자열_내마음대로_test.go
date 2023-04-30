package main

// func solution(strings []string, n int) []string {

// 	sort.Slice(strings, func(i, j int) bool {
// 		strA, strB := []byte(strings[i]), []byte(strings[j])

// 		if strA[n] == strB[n] {
// 			return strings[i] < strings[j]
// 		}

// 		return strA[n] < strB[n]
// 	})

// 	return strings
// }

// func TestSolution(t *testing.T) {

// 	t1 := solution([]string{"sun", "bed", "car"}, 1)
// 	t2 := solution([]string{"abce", "abcd", "cdx"}, 2)

// 	assert.Equal(t, t1, []string{"car", "bed", "sun"})
// 	assert.Equal(t, t2, []string{"abcd", "abce", "cdx"})
// }
