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

func fibBottomUp(n int) int {
	if n < 2 {
		return n
	}
	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1
	table[2] = 1
	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

func main() {
	// fmt.Println(fibonacci(6))

	// fmt.Println(fibMemoized(6, map[int]int{}))
	fmt.Println(fibBottomUp(6))
}
