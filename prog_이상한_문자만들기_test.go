package main

// func solution(s string) string {

// 	strA := strings.Split(s, " ")

// 	strB := ""
// 	for _, s := range strA {

// 		for j, v := range s {

// 			if j == 0 || j%2 == 0 {
// 				strB += strings.ToUpper(string(v))
// 				continue
// 			}

// 			strB += strings.ToLower(string(v))
// 		}

// 		strB += " "
// 	}

// 	return strB[0 : len(strB)-1]
// }

// func TestSolution(t *testing.T) {
// 	t1 := solution("try hello world")
// 	t2 := solution("  ")
// 	t3 := solution("they are nothing")

// 	assert.Equal(t, t1, "TrY HeLlO WoRlD")
// 	assert.Equal(t, t2, "  ")
// 	assert.Equal(t, t3, "ThEy ArE NoThInG")
// }
