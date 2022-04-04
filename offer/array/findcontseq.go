package array

// 滑动窗口
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	if target < 3 {
		return res
	}

	left, right, sum := 1, 2, 3
	for left < right {
		//fmt.Printf("%d, %d, %d\n", left, right, sum)

		if sum < target {
			right++
			sum += right
		} else if sum > target {
			sum -= left
			left++
		} else {
			temp := make([]int, 0)
			for i := left; i <= right; i++ {
				temp = append(temp, i)
			}
			res = append(res, temp)

			sum -= left
			left++
		}
	}

	return res
}

// 暴力枚举
// func findContinuousSequence(target int) [][]int {
// 	res := make([][]int, 0)

// 	for i := 1; i <= target/2; i++ {
// 		sum := i
// 		temp := make([]int, 0)

// 		for j := i + 1; j < target; j++ {
// 			sum += j
// 			if sum > target {
// 				break
// 			} else if sum == target {
// 				for k := i; k <= j; k++ {
// 					temp = append(temp, k)
// 				}
// 				res = append(res, temp)
// 			}
// 		}
// 	}

// 	return res
// }
