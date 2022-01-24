package main

import (
	"fmt"
)

func A(nums []int) {
	nums[0] = 0
}

func main() {
	nums := []int{1, 2, 3}
	A(nums)

	fmt.Println(nums)
}
