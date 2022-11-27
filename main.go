package main

import (
	"fmt"
)


func main(){
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{2,4,1}))	
	fmt.Println(maxProfit([]int{1,2}))	

	// v := []int{7,1,5,3,6,4}
	// fmt.Println(v)
	// fmt.Println(v[0:5])
}

// DD
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	
	sum := 0
	buy := prices[0]
	for i,sell := range prices {
		if i == 0 {
			continue
		}

		if buy > sell {
			buy = sell
			continue
		}

		if buy < sell && sum < sell-buy {
			sum = sell-buy
		}	
	}

	return sum
}
