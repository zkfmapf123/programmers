// package main

// import (
// 	"fmt"
// )

// func main(){
// 	// fmt.Println(isSubsequence("abc","ahbgdc"))
// 	// fmt.Println(isSubsequence("axc","ahdgdc"))
// 	// fmt.Println(isSubsequence("acb","ahbgdc"))
// 	// fmt.Println(isSubsequence("ab","baab"))
// 	fmt.Println(isSubsequence("abc","acb"))
// }

// func isSubsequence(s string, t string) bool {

// 	sp, tp := 0, 0
// 	for sp < len(s) && tp < len(t) {
// 		if s[sp] == t[tp] {
// 			sp++
// 		}
// 		tp++
// 	}

// 	return sp == len(s)
// }