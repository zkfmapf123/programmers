package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	// fmt.Println(romanToInt("III"))
	// fmt.Println(romanToInt("LVIII"))
	// fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("MCDLXXVI"))
}

func romanToInt(s string) int {
	sum := 0
	r := strings.NewReplacer(
		"CM","900,",
		"CD" , "400,",
		"XC","90,",
		"XL","40,",
		"IV","4,",
		"IX","9,",
		"I","1,",
		"V","5,",
		"X","10,",
		"L","50,",
		"C","100,",
		"D","500,",
		"M","1000,",
	)

	result := strings.Split(r.Replace(s), ",")
	fmt.Println(result)
	for _, r  := range result {
		num,_ := strconv.Atoi(r)
		sum += num
	}

	return sum
}