package arrays

import (
	"fmt"
)

func main() {

	var a [5]int
	//a is a list of 5 integers
	fmt.Println("list a:", a)

	a[4] = 100
	fmt.Println("POST:", a)
	fmt.Println("GET:", a[4])
	a[4] = 0
	fmt.Println("DELETE:", a)
	fmt.Println("GET:", a[4])

	singledigit := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Single Digits:", singledigit)
}
