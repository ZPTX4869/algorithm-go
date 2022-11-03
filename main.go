package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	bytes := []byte(strconv.Itoa(n))
	asc_queue := make([]int, 0)

	for i := len(bytes) - 1; i >= 0; i-- {
		if len(asc_queue) > 0 && bytes[i] < bytes[asc_queue[len(asc_queue)-1]] {
			for _, idx := range asc_queue {
				if bytes[idx] > bytes[i] {
					bytes[i], bytes[idx] = bytes[idx], bytes[i]
					break
				}
			}

			ss := bytes[i+1:]
			sort.Slice(ss, func(i int, j int) bool {
				return ss[i] < ss[j]
			})

			break
		}

		asc_queue = append(asc_queue, i)
	}

	ans, _ := strconv.ParseInt(string(bytes), 10, 64)

	if len(asc_queue) == len(bytes) || ans > int64(math.MaxInt32) {
		return -1
	}

	return int(ans)
}

func main() {
	n := 21

	ans := nextGreaterElement(n)
	fmt.Println(ans)
}
