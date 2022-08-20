package debug

import "fmt"

func PrintIndent(content string, numIndent int) {
	for i := 0; i < numIndent; i++ {
		fmt.Print(" ")
	}
	fmt.Println(content)
}
