// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main(){
// 	r1 := solution("aukks", "wbqd",5)
// 	r2 := solution("ybcde", "az",1)
// 	fmt.Println(r1, r2)
// }

// func solution(s string, skip string, index int) string {
// 	exceptSkipAlPha := deleteStr(getAlpha(), skip)
// 	exceptSkipAlPhaArr := strings.Split(exceptSkipAlPha,"")
// 	arrLen := len(exceptSkipAlPhaArr)
// 	var result string

// 	for _, v := range strings.Split(s,"") {
// 		nextIndex := strings.Index(exceptSkipAlPha, v) + index

// 		if nextIndex >= arrLen {
// 			if nextIndex % arrLen == 0 {
// 				nextIndex = 0
// 			}else{
// 				nextIndex /= arrLen
// 			}
// 		}

// 		result += exceptSkipAlPhaArr[nextIndex]

// 	}

//     return strings.Trim(result," ")
// }

// func getAlpha() string{
// 	return "abcdefghijklmnopqrstuvwxyz"
// }

// func deleteStr(s string, delStrs string) string{

// 	for _, delStr := range strings.Split(delStrs, "") {
// 		s = strings.Replace(s, delStr, "", -1)
// 	}

// 	return s
// }