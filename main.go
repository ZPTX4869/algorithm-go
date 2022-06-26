package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)

	copy(a[1:], []int{1, 2, 3})
	fmt.Println(a)
}
