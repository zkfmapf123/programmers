// package main

// func maxProfit(prices []int) int {
// 	if prices == nil || len(prices) == 0 {
// 		return 0
// 	}

// 	sum := 0
// 	buy := prices[0]
// 	for i,sell := range prices {
// 		if i == 0 {
// 			continue
// 		}

// 		if buy > sell {
// 			buy = sell
// 			continue
// 		}

// 		if buy < sell && sum < sell-buy {
// 			sum = sell-buy
// 		}
// 	}

// 	return sum
// }
