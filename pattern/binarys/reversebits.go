package binarys

func reverseBits(num uint32) uint32 {
	var res uint32

	//这个地方可以优化，因为num可能不需要检测全部32位就已经为0了
	//for i := 0; i < 32; i++ {
	//	res = (res << 1) | (num & 1)
	//	num >>= 1
	//}

	for i := 0; i < 32 && num > 0; i++ {
		res |= (num & 1) << (31 - i)
		num >>= 1
	}

	return res
}
