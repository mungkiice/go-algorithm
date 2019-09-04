package main

import "fmt"

func countPerfectSubstr(s string, k int) int {
	var count int
	for i,v := range s {
		elements := map[string]int{}
		elements[string(v)] = 1
		for j:=i;j<len(s);j++{
			w := string(s[j])
			if _,ok := elements[w]; i!=j && ok{
				elements[w]++
			}else if i != j{
				elements[w] = 1
			}
			if isPerfect(elements, k){
				count++
			}
		}
	}
	return count
}

func isPerfect(elements map[string]int, k int) bool {
	for _,count := range elements {
		if count != k {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(countPerfectSubstr("1102021222", 2))
}