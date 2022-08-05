package main

import (
	"fmt"
)

func main() {
	var c chan int
	c = make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 1
}
