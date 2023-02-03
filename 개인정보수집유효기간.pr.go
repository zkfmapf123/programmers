package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	r1 := solution("2022.5.19", []string{"A 6","B 12","C 3"}, []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"})
	r2 := solution("2020.01.01", []string{"Z 3", "D 5"}, []string{"2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"})
	fmt.Println(r1, r2)
}

func solution(today string, terms []string, privacies []string) []int {	
	tm, todayNum := getTermsMap(terms), makeDateForNum(splitYearMonthDate(today))
	
	var destroyPv []int
	for i, pv := range privacies {
		splitTerms := strings.Split(pv, " ")
		convertPv := addMonthPv(splitTerms[0], tm[splitTerms[1]])

		// Destroy
		if convertPv <= todayNum {
			destroyPv = append(destroyPv, i+1)
		}
	}

    return destroyPv
}

func getTermsMap(terms []string) map[string]int{
	m := make(map[string]int)

	for _, term := range terms{
		splitTerms := strings.Split(term," ")
		num, _ := strconv.Atoi(splitTerms[1])
		m[splitTerms[0]] =  num
	}

	return m
}

func convertDate(ds string) int {
	num,_ :=  strconv.Atoi(strings.Replace(ds,".","",-1))
	return num 
}

func splitYearMonthDate(pv string) (int, int, int){
	splitPv := strings.Split(pv,".")

	year, _ := strconv.Atoi(splitPv[0])
	month, _ := strconv.Atoi(splitPv[1])
	date, _ := strconv.Atoi(splitPv[2])

	return year, month, date
}

// 2021.05.02
// 2021.07.01
func addMonthPv(pv string, addMonth int) int {
	year, month, date := splitYearMonthDate(pv)
	month = month + addMonth

	if month > 12 {
		year, month = year + (month / 12), (month % 12)
		// 12월
		if month == 0 {
			year -= 1
			month = 12
		}
 	}

	return makeDateForNum(year, month, date)
}

func makeDateForNum(year, month, date int) int {
	var _year, _month, _date string

	_year = strconv.Itoa(year)
	if month < 10 {
		_month = fmt.Sprintf("0%d", month)
	}else {
		_month = strconv.Itoa(month)
	}

	_date = strconv.Itoa(date)
	if date < 10 {
		_date = fmt.Sprintf("0%d", date)
	}else{
		_date = strconv.Itoa(date)
	}
	
	convertStr := strings.Join([]string{_year, _month, _date}, "")
	convertNum , _ := strconv.Atoi(convertStr)

	return convertNum
}