package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int) {
	mu.Lock()
	defer mu.Unlock()

	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
