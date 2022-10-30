package main

import (
	"fmt"
)

func main(){
	r1 :=solution(10,2)
	fmt.Println(r1)

	r2 :=solution(8,1)
	fmt.Println(r2)

	r3 :=solution(24,24)
	fmt.Println(r3)

	r4 :=solution(12,4)
	fmt.Println(r4)

}


	func solution(brown int, yellow int) []int {
		xpy := brown / 2 + 2
		for x := 3; x <= xpy; x++ {
			for y := 3; y <= x; y++ {
				if yellow == (x - 2) * (y - 2) && brown == (x + y) * 2 - 4 {
					return []int { x, y }
				}
			}
		}
	
		return []int { 0, 0 } 
	}