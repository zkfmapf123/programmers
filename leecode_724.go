package main

// import "fmt"

// func main(){
// 	r1 := pivotIndex([]int{1,7,3,6,5,6})
// 	r2 := pivotIndex([]int{1,2,3})
// 	r3 := pivotIndex([]int{2,1,-1})

// 	fmt.Println(r1)
// 	fmt.Println(r2)
// 	fmt.Println(r3)
// }

// func pivotIndex(nums []int) int {
// 	index := -1

// 	for i, _:= range nums {

// 		var ls int
// 		var rs int

// 		arrLen := len(nums)
// 		if i == 0 {
// 			ls = sumArray([]int{0})
// 			rs = sumArray(nums[1:arrLen])
// 		}else if i == arrLen {
// 			ls = sumArray(nums[0:arrLen])
// 			rs = sumArray([]int{0})
// 		}else{
// 			ls = sumArray(nums[0:i])
// 			rs = sumArray(nums[i+1:arrLen])
// 		}

// 		if ls == rs {
// 			return i
// 		}

// 	}

// 	return index
// }

// func sumArray(nums []int) int{

// 	sum := 0
// 	for _, v:= range nums {
// 		sum += v
// 	}
// 	return sum
// }