package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3

	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("I don't know that number")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend! Go have fun.")
	default:
		fmt.Println("Get back to work, it's not the weekend!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 8:
		fmt.Println("Coleman is asleep")
	case t.Hour() > 22:
		fmt.Println("Coleman went to bed")
	case t.Hour() > 8 && t.Hour() < 18:
		fmt.Println("Coleman is working")
	default:
		fmt.Println("I dont know what Coleman is doing!")
	}
}
