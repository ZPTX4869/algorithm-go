package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2}
	sort.Ints(nums)
	fmt.Println(nums)
}
