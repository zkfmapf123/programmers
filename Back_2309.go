// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// )

// var (
// 	w = bufio.NewWriter(os.Stdout)
// 	r = bufio.NewReader(os.Stdin)
// )

// // 아홉명의 난쟁이
// // 난쟁이는 총 일곱명 === 100
// // 2명만 빼면 됨
// func main(){
// 	w.Flush()
// 	N := 9

// 	arr, sum := []int{}, 0

// 	for i:=0; i<N; i++ {
// 		b, _, _ := r.ReadLine()
// 		str := string(b)
// 		v, _ := strconv.Atoi(str)
// 		arr = append(arr, v)
// 		sum += v
// 	}

// 	sort.Ints(arr)
// 	for i:=0; i<len(arr)-1; i++ {
// 		for j:=i+1; j<len(arr); j++ {

// 			if arr[i] + arr[j] == (sum - 100) {

// 				for k:=0; k<len(arr); k++ {
// 					if k == i || k ==j {
// 						continue
// 					}

// 					fmt.Println(arr[k])
// 				}
// 				return
// 			}
// 		}
// 	}
// }

// func getN() int{
// 	var N int
// 	fmt.Fscanf(r, "%d\n", &N)
// 	return N
// }
