package offer

func reverseLeftWords(s string, n int) string {
	var res string

	// for i := n; i < len(s); i++ {
	// 	res += string(s[i])
	// }

	// for i := 0; i < n; i++ {
	// 	res += string(s[i])
	// }

	res += s[n:] + s[:n]

	return res
}
