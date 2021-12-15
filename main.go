package main

import "fmt"

type A interface{}
type B struct{}

var a *B

func main() {
	fmt.Println(a == nil)
	fmt.Println(a == (*B)(nil))
	fmt.Println((A)(a) == nil)
}
