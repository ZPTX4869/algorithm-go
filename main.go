package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var c chan int
	c = make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 1

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sc.Text()
}
