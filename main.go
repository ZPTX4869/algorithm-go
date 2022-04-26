package main

import (
	"algorithm-go/util"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./main.go")
	util.CheckOrPanic(err)
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}
