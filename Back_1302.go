package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// var (
// 	w = bufio.NewWriter(os.Stdout)
// 	r = bufio.NewReader(os.Stdin)
// )

// func main(){
// 	defer w.Flush()

// 	N := getN()

// 	m := make(map[string]int)

// 	// o(n)
// 	for i:=0; i<N; i++ {
// 		str, _, _ := r.ReadLine()
// 		value := string(str)

// 		v, isOk := m[value]
// 		if isOk {
// 			v++
// 			m[value] = v
// 		}else{
// 			m[value] = 1
// 		}
// 	}

// 	// o(n)
// 	max := -1
// 	for _,v := range m {
// 		if v > max {
// 			max =v
// 		}
// 	}

// 	// o(n)
// 	book := []string{}
// 	for k,v := range m {

// 		if v == max {
// 			book = append(book, k)
// 		}
// 	}

// 	sort.Strings(book)
// 	fmt.Println(book[0])
// }

// func getN() int{
// 	var N int
// 	fmt.Fscanf(r, "%d\n", &N)
// 	return N
// }
