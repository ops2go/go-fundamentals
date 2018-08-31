package main

import (
	"fmt"
)

func main() {
	//empty slice
	s := make([]string, 3)
	fmt.Println("s:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Update:", s)
	fmt.Println("Length:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
