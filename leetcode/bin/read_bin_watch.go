package bin

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)

	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			nHour := bits.OnesCount8(uint8(i))
			nMinute := bits.OnesCount8(uint8(j))

			if nHour+nMinute == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}

	return res
}
