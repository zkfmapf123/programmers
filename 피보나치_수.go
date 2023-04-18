package main

// func main(){
// 		r1 :=solution(3)
// 		r2 := solution(5)
// 		r3 := solution(10)
// 		fmt.Println(r1)
// 		fmt.Println(r2)
// 		fmt.Println(r3)
// }

// func solution(n int) int {

// 	cache := make([]int, n+1,n+2)
// 	cache[0], cache[1] = 0,1

// 	for i:=2; i<=n; i++ {
// 		cache[i] = (cache[i-2] + cache[i-1]) % 1234567
// 	}

// 	return cache[n]
// }
