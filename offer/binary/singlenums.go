package binary

import (
	"fmt"
)

// 除了两个数字外其他重复2次
func singleNumbers(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor = xor ^ num
	}

	var bit int
	for i := 0; i < 31; i++ {
		fmt.Printf("%032b\n%032b\n\n", xor, 1<<i)
		if xor&(1<<i) != 0 {
			bit = i
			break
		}
	}
	fmt.Println(bit)

	xor1, xor2 := 0, 0
	for _, num := range nums {
		if num&(1<<bit) == 0 {
			xor1 = xor1 ^ num
		} else {
			xor2 = xor2 ^ num
		}
		fmt.Printf("num: %d | xor1: %d | xor2: %d\n", num, xor1, xor2)
	}

	return []int{xor1, xor2}
}
