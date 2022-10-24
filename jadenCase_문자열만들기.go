package main

import "strings"

/*
	Camel Case 전략
	그외에은 소문자
*/

func solution(s string) string{
	var convertStr []string
	str := strings.Split(s," ")

	for _, s := range str {
		convertStr = append(convertStr,strings.Title(s))
	}

	return strings.Join(convertStr," ")
}


