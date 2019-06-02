package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 10
	for {
		str := strconv.Itoa(num)
		str2 := fmt.Sprintf("%b", num)
		str8 := fmt.Sprintf("%o", num)

		if palindrome(str) && palindrome(str2) && palindrome(str8) {
			fmt.Println(num)
			break
		}
		num++
	}
}

func palindrome(str string) bool {
	if str == reverse(str) {
		return true
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
