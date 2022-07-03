package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./go.mod")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
