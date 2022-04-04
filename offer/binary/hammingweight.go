package binary

func hammingWeight(num uint32) int {
    var res int
	
	for num != 0 {
		if num & 1 > 0 {
			res += 1
		}
		num = num >> 1
	}

	return res
}

func hammingWeight2(num uint32) int {
    var res int
	
	for i := 0; i < 31; i++ {
		if 1 << i & num != 0 {
			res += 1
		}
	}

	return res
}