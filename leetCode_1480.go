package main

// import "fmt"

// func main(){
// 	r1 := runningSum([]int{1,2,3,4})
// 	r2 := runningSum([]int{1,1,1,1,1})
// 	r3 := runningSum([]int{3,1,2,10,1})

// 	fmt.Println(r1)
// 	fmt.Println(r2)
// 	fmt.Println(r3)
// }

// // Time O(n) + Memory O(n)
// // func runningSum(nums []int) []int {
// // 	var result = make([]int, len(nums))
// // 	sum := 0

// // 	for i,v := range nums{
// // 		sum += v
// // 		result[i] = sum
// // 	}

// // 	return result
// // }

// // Time O(n) + Memory (1)
// func runningSum(nums []int) []int {
// 	sum := 0
// 	for i,v := range nums{
// 		sum += v
// 		nums[i] = sum
// 	}

// 	return nums
// }