package main

import "fmt"

type A interface{}
type B struct{}

var a *B

func main() {
	fmt.Println(a == nil)            //true
	fmt.Println(a == (*B)(nil))      //true
	fmt.Println((A)(a) == (*B)(nil)) //true

	fmt.Println((A)(a) == nil) //false
}
