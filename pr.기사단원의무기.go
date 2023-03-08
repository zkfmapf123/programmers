package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(solution(5,3,2))
}

func get(n int ) int {
	v := make(map[int]struct{})
	for i:=1; i<= int(math.Sqrt(float64(n))); i++ {
		if n %i == 0 {
			v[i] = struct{}{}
			v[n/i] = struct{}{}
		}
	}

	fmt.Println(v)
	return len(v)
}

func solution(number int, limit int, power int) int {

	iron := 0

	for i:=1; i<= number; i++ {

		w := get(i)

		if w > limit {
			iron += power
		}else{
			iron += w
		}
	}

	return iron
}