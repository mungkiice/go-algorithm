package main

import "fmt"

func fibonacci(n int) int {
	fmt.Println("Run fib for", n)
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibMemoized(n int, m map[int]int) int {
	if n < 2 {
		return n
	} else if m[n] == 0 {
		m[n] = fibMemoized(n-1, m) + fibMemoized(n-2, m)
	}
	return m[n]
}

func main() {
	// fmt.Println(fibonacci(6))

	fmt.Println(fibMemoized(6, map[int]int{}))
}
