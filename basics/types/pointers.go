package main

import (
	"fmt"
)

func Pointer(x int) (int, *int, int) {
	a := &x
	return x, a, *a
}


func main() {
	fmt.Println("Hello, playground")
	fmt.Println(Pointer(15))
}
