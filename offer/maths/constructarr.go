package maths

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return []int{}
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	
	dp1, dp2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp1[i] = 1
		dp2[i] = 1
	}

	for i := 1; i < n; i++ {
		dp1[i] = dp1[i-1] * a[i-1]
		dp2[n-1-i] = dp2[n-i] * a[n-i]
	}

	for i := 0; i < n; i++ {
		res[i] = dp1[i] * dp2[i]
	}

	return res
}

// 双循环
// func constructArr(a []int) []int {
// 	n := len(a)
// 	if n == 0 {
// 		return []int{}
// 	}

// 	res := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		res[i] = 1
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n ; j++ {
// 			if j  == i {
// 				continue
// 			}
// 			res[i] *= a[j]
// 		}
// 	}

// 	return res
// }
