package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	fmt.Println(palindrome("ibu ratna antar ubi"))
	fmt.Println(palindrome("kasur ini rusak"))
	fmt.Println(palindrome("A nut for a jar of tuna."))
	fmt.Println(palindrome("Borrow or rob?"))
	fmt.Println(palindrome("Was it a car or a cat I saw?"))
	fmt.Println(palindrome("Yo, banana boy!"))
	fmt.Println(palindrome("UFO tofu?"))

}

func palindrome(str string) bool {
	param := regexp.MustCompile(`[\W_]`)
	clean_lowered_str := param.ReplaceAllString(strings.ToLower(str), "")

	if clean_lowered_str == reverseString(clean_lowered_str) {
		return true
	}

	return false
}

func reverseString(str string) string {
	res := ""
	for _, i := range str {
		res = string(i) + res
	}

	return res
}
