package binary

func singleNumber(nums []int) int {
	var res int

	for _, num := range nums {
		res = res ^ num
	}

	return res
}

func singleNumber2(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		var sum int32
		for j := 0; j < len(nums); j++ {
			sum += (int32(nums[j]) >> int32(i)) & int32(1)
		}
		res += (sum % 3) << int32(i)
	}

	return int(res)
}

func singleNumber3(nums []int) []int {
	var xorsum int

	for _, num := range nums {
		xorsum = xorsum ^ num
	}

	lsb := ((xorsum - 1) & xorsum) ^ xorsum

	var res0, res1 int
	for _, num := range nums {
		if num&lsb == 0 {
			res0 = res0 ^ num
		} else {
			res1 = res1 ^ num
		}
	}

	return []int{res0, res1}
}
