package main

// go test -run ./proc_둘만의암호_test.go
// func solution(s string, skip string, index int) string {

// 	result := ""
// 	alphbet := getAlphbetMap(skip)

// 	for _, r := range s {
// 		str := string(r)
// 		result += findAlphbet(alphbet, str, index)
// 	}

// 	return result

// }

// func findAlphbet(alphbet string, target string, index int) string {

// 	alphabetList := strings.Split(alphbet, "")

// 	idx := -1
// 	for i, v := range alphabetList {

// 		if v == target {
// 			idx = i
// 		}
// 	}

// 	selectIdx := idx + index

// 	for selectIdx > len(alphabetList) {
// 		selectIdx = (idx + index) - len(alphabetList)
// 	}

// 	fmt.Println(alphabetList, (idx + index), len(alphabetList))
// 	return alphabetList[idx+index]
// }

// func getAlphbetMap(skip string) string {

// 	alphabet := "abcdefghijklmnopqrstuvwxyz"

// 	for _, c := range skip {
// 		alphabet = strings.Replace(alphabet, string(c), "", 1)
// 	}

// 	return alphabet
// }

// func TestSolution(t *testing.T) {
// 	v1 := solution("aukks", "wbqd", 5)
// 	// v2 := solution("zzzzzz", "abcdefghijklmnopqrstuvwxy", 6)

// 	assert.Equal(t, v1, "happy")
// 	// assert.Equal(t, v2, "zzzzzz")
// }
