package binary

func hammingWeight(num uint32) int {
	var cnt int

	for num != 0 {
		num = (num - 1) & num
		cnt++
	}

	return cnt
}

//func hammingWeight(num uint32) int {
//	var cnt int
//
//	for i := 0; i < 32; i++ {
//		if num&(uint32(1)<<uint32(i)) != uint32(0) {
//			cnt++
//		}
//	}
//
//	return cnt
//}
