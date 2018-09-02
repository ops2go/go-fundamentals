package main

// packages dont need main
// funcmain and packagemain in same file

import "fmt"

/*
Functions are central in Go.
We’ll learn about functions with a few different examples.

func {receiver} {name}({parameter(s)}) return {
	{function}
}
*/
func plus(a int, b int) int {
	//requires a and be integer, returns an integer
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	res = minus(1, 2)
	fmt.Println("1-2 =", res)
	res = multiply(1, 2)
	fmt.Println("1*2 =", res)
	res = divide(1, 2)
	fmt.Println("1/2 =", res)

}
