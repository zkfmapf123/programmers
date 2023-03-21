
/*
	1. 맨 앞의 값을 삭제 
	2. 맨 앞의 값을 맨 뒤로 보내기

	=> o(n) * o(n) => 2n
*/


// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var (
// 	w = bufio.NewWriter(os.Stdout)
// 	r = bufio.NewReader(os.Stdin)
// )

// func main(){
// 	defer w.Flush()
	
// 	n := getN()
// 	var arr = make([]int, n)

// 	for i:=1; i<=n; i++ {
// 		arr[i-1] = i
// 	}

// 	for len(arr) > 1{
		
// 		// 하나 버림
// 		que := arr[2:]

// 		// 한장은 뒤로돌림
// 		que = append(que, arr[1])
		
// 		arr =que
// 	}

// 	fmt.Println(arr[0])
// }

// func getN() int{
// 	var n int
// 	fmt.Fscanf(r, "%d\n", &n)
// 	return n
// }