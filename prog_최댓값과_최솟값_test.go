package main

// func solution(s string) string {

// 	arr := strings.Split(s, " ")
// 	sort.Slice(arr, func(i, j int) bool {

// 		numI, _ := strconv.Atoi(arr[i])
// 		numJ, _ := strconv.Atoi(arr[j])

// 		return numI < numJ

// 	})

// 	return fmt.Sprintf("%s %s", arr[0], arr[len(arr)-1])
// }

// func TestSolution(t *testing.T) {
// 	t1 := solution("1 2 3 4")
// 	t2 := solution("-1 -2 -3 -4")
// 	t3 := solution("-1 -1")

// 	assert.Equal(t, t1, "1 4")
// 	assert.Equal(t, t2, "-4 -1")
// 	assert.Equal(t, t3, "-1 -1")
// }
