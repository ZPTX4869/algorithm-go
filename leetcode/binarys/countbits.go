package binarys

func countBits(n int) []int {
	res := make([]int, 0)

	for i := 0; i <= n; i++ {
		cur := i
		cnt := 0
		for cur != 0 {
			cur = (cur - 1) & cur
			cnt += 1
		}
		res = append(res, cnt)
	}

	return res
}

//func countBits(n int) []int {
//	res := make([]int, 0)
//
//	for i := 0; i <= n; i++ {
//		cnt := 0
//		for j := 0; j < 32; j++ {
//			if i&(1<<j) != 0 {
//				cnt += 1
//			}
//		}
//		res = append(res, cnt)
//	}
//
//	return res
//}
