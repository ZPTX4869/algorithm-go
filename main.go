package main

import (
	"algorithm-go/util"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("go.mod")
	defer f.Close()
	util.CheckOrPanic(err)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}
