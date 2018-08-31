package main

import (
	"fmt"
)

//Problem #4
//Create constants for a number and a result

const c string = "constant"

func main() {
	fmt.Println(c)

	const number = 50000000

	const result = 723 * number

	fmt.Println("number times 723 =", result)
}
