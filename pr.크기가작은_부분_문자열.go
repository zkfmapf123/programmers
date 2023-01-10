// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main(){
// 	fmt.Println(solution("3141592","271"))
// 	fmt.Println(solution("500220839878","7"))
// 	fmt.Println(solution("10203","15"))
// }

// func solution(t string, p string) int{
// 	result := 0
// 	pNum, _:= strconv.Atoi(p)
// 	str, tArr := "", strings.Split(t,"")

// 	for i, _:= range tArr{
		
// 		lastIndex := i + len(p)
// 		if lastIndex > len(tArr) {
// 			break	
// 		}

// 		str = strings.Join(tArr[i:i+len(p)], "")

// 		if len(p) == len(str){
// 			num, _:= strconv.Atoi(str);

// 			if pNum >= num {
// 				result += 1
// 			}
// 		}
// 		str = ""
// 	}

// 	return result
// }