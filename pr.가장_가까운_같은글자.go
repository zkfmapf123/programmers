// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main(){
// 	fmt.Println(solution("banana"))
// 	fmt.Println(solution("foobar"))
// }

// func solution(s string) []int {
// 	strArr := strings.Split(s, "")
// 	result := make([]int,len(s))	
// 	result[0] = -1

// 	resultLen := len(result)
// 	for i := 1; i<resultLen; i++ {
		
// 		str := strArr[i]
// 		arr := strArr[0:i]

// 		lastIndex := strings.LastIndex(strings.Join(arr,""), str)
// 		if lastIndex == -1 {
// 			result[i] = -1
// 		}else{
// 			result[i] = i - lastIndex
// 		}
// 	}
//     return result
// }