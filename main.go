package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	bytes := []byte(strconv.Itoa(n))
	ascQueue := make([]int, 0)

	for i := len(bytes) - 1; i >= 0; i-- {
		if len(ascQueue) > 0 && bytes[i] < bytes[ascQueue[len(ascQueue)-1]] {
			for _, idx := range ascQueue {
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

		ascQueue = append(ascQueue, i)
	}

	ans, _ := strconv.ParseInt(string(bytes), 10, 64)

	if len(ascQueue) == len(bytes) || ans > int64(math.MaxInt32) {
		return -1
	}

	return int(ans)
}

func main() {
	n := 21

	ans := nextGreaterElement(n)
	fmt.Println(ans)
}
