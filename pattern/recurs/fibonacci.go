package recurs

// func fib(n int) int {
// 	if n < 2 {
// 		return n
// 	}

// 	return fib(n-1) + fib(n-2)
// }

// 优化版，缓存可能会重复计算的数
func fib(n int) int {
	m := make(map[int]int)

	return dfs(n, m)
}

func dfs(n int, m map[int]int) int {
	if n < 2 {
		return n
	}

	if result, ok := m[n]; ok {
		return result
	}

	return dfs(n-1, m) + dfs(n-2, m)
} 