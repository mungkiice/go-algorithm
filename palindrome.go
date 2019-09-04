package main

import "fmt"

func countPalindrome(s string) int {
	palindromes := map[string]bool{}
	for i := range s {
		expand(s, i, i, &palindromes)
		expand(s, i, i+1, &palindromes)
	}
	return len(palindromes)
}

func expand(s string, low, high int, result *map[string]bool){
	for low >= 0 && high < len(s) && s[low] == s[high]{
		palindrome := string(s[low:high+1])
		if _,ok := (*result)[palindrome];!ok {
			(*result)[palindrome] = true
		}
		low--
		high++
	}
}

func main(){
	fmt.Println(countPalindrome("aabaa"))
}