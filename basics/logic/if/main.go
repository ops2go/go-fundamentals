package main

import (
	"fmt"
)

var name string = "Lacee"

func main() {
	if name == "Coleman" {
		fmt.Println("Your'e the coolest guy ever")
	} else if name == "Lacee" {
		fmt.Println("Your'e the coolest girl ever")
	} else {
		fmt.Println("You're not that cool")
	}
}
