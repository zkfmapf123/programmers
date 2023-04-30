package main

// func solution(s string, n int) string {
// 	r := make([]rune, len(s))
// 	offset := rune(n)

// 	for i, v := range s {

// 		if 'A' <= v && v <= 'Z' {
// 			r[i] = ((v + offset - 'A') % 26) + 'A'
// 			continue
// 		}

// 		if 'a' <= v && v <= 'z' {
// 			r[i] = ((v + offset - 'a') % 26) + 'a'
// 			continue
// 		}

// 		r[i] = v
// 	}

// 	return string(r)
// }

// func TestSolution(t *testing.T) {

// 	t1 := solution("AB", 1)
// 	t2 := solution("z", 1)
// 	t3 := solution("a B z", 4)

// 	assert.Equal(t, t1, "BC")
// 	assert.Equal(t, t2, "a")
// 	assert.Equal(t, t3, "e F d")
// }
