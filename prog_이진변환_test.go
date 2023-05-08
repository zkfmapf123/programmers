package main

// func solution(s string) []int {
// 	r := regexp.MustCompile("0")

// 	totalCount, zeroCount := 0, 0
// 	for s != "1" {
// 		totalCount++

// 		// get zero count
// 		zero := r.FindAllString(s, -1)
// 		if len(zero) > 0 {
// 			zeroCount += len(zero)
// 		}

// 		// replace zero string
// 		s = r.ReplaceAllString(s, "")
// 		s = decimal(len(strings.Split(s, "")))
// 	}

// 	return []int{totalCount, zeroCount}
// }

// func decimal(str int) string {
// 	bin := ""
// 	for str > 0 {
// 		remain := str % 2
// 		bin = fmt.Sprintf("%d%s", remain, bin)
// 		str = str / 2
// 	}

// 	return bin
// }

// func TestSolution(t *testing.T) {
// 	v1 := solution("110010101001")
// 	v2 := solution("01110")
// 	v3 := solution("1111111")

// 	assert.Equal(t, v1, []int{3, 8})
// 	assert.Equal(t, v2, []int{3, 3})
// 	assert.Equal(t, v3, []int{4, 1})
// }
