package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1)+fibonacci(n-2);
}
func memoize(fn func(int)int) func(int)int{
	cache := map[int]int{}
	return func(n int) int {
		if cache[n] == 0 {
			cache[n] = fn(n)
		}
		return cache[n]
	}
}
func main(){	
	newFib := memoize(fibonacci)
	fmt.Println(newFib(6))
	
	// fmt.Println(fibonacci(6))
}