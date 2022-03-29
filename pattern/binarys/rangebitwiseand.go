package binarys

func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right &= right - 1
	}

	return right
}

//func rangeBitwiseAnd(left int, right int) int {
//	shift := 0
//	for right > left {
//		left, right = left>>1, right>>1
//		shift++
//	}
//
//	return left << shift
//}

// 执行超时
//func rangeBitwiseAnd(left int, right int) int {
//	res := left
//	for i := left + 1; i <= right; i++ {
//		res = res & i
//	}
//
//	return res
//}
