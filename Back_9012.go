// package main

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// var (
// 	w = bufio.NewWriter(os.Stdout)
// 	r = bufio.NewReader(os.Stdin)
// )

// func main(){
// 	defer w.Flush()
// 	n := getN()

// 	for i:=0; i<n; i++ {
// 		txt, _ := r.ReadString('\n')
// 		str := strings.TrimSuffix(txt, "\n")
// 		s := newStack()
// 		isMatch := true

// 		for _,v := range strings.Split(str, ""){
			
// 			if v == "(" {
// 				s.push(v)
// 			}

// 			prefix, err := s.peek()

// 			if err != nil {
// 				isMatch = false
// 				break	
// 			}

// 			if err == nil && prefix == "(" && v == ")" {
// 				s.pop()
// 			}
// 		}

// 		if isMatch && s.idx == -1 {
// 			fmt.Println("YES")
// 		}else{
// 			fmt.Println("NO")
// 		}
// 	}
// }

// func getN() int{
// 	var n int
// 	fmt.Fscanf(r, "%d\n", &n)
// 	return n
// }

// type Stack struct {
// 	list []string
// 	idx int
// }

// func newStack() *Stack	{
// 	return &Stack{
// 		list : []string{},
// 		idx : -1,
// 	}
// }

// func (s *Stack) push(v string){
// 	s.idx++
// 	s.list = append(s.list, v)
// }

// func (s *Stack) pop(){
// 	s.list = s.list[0:s.idx]
// 	s.idx--
// }

// func (s *Stack) peek() (string, error){
// 	if s.idx == -1 {
// 		return "", errors.New("s.idx is nil")
// 	}

// 	return s.list[s.idx], nil
// }