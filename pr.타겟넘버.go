// package main

// import "fmt"

// func main(){
//     fmt.Println(solution([]int{1, 1, 1, 1, 1}, 3))
//     fmt.Println(solution([]int{4,1,2,1},4))
    
// }

// var profitNums int

// func solution(numbers []int, target int) int {
//     profitNums = 0
//     bfs(len(numbers), 0, 0, target, numbers)

//     return profitNums
// }

// func bfs(maxLen int, idx int, sum int, target int, nums []int) {
    
//     if idx == maxLen {
//         if sum == target {
//             profitNums++
//         }

//         return 
//     }

//     bfs(maxLen, idx+1, sum + nums[idx], target, nums)
//     bfs(maxLen, idx+1, sum - nums[idx], target, nums)
// }