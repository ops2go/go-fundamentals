package main

import (
	"fmt"
)

// Problem #2 N/A
func main() {
	fmt.Println("Strings:", "Coleman", "Word")
	fmt.Println("Integers:", "7 times 3 =", 7*3)
	fmt.Println("Floats:", "3.14 times 7.90 =", 3.14*7.90)
	fmt.Println("Booleans:", true && true)
	fmt.Println("Booleans:", true && false)
	fmt.Println("Booleans:", true || false)
	fmt.Println("Booleans:", false || true)
	fmt.Println("Booleans:", !true)
	fmt.Println("Booleans:", !false)
	//string
	name := "Coleman"
	//int
	age := 25
	//float64
	height := 6.10
	//constant
	const gender string = "male"
	//array
	fingers := [5]string{"thumb", "pointer", "middle", "ring", "pinky"}
	//slice
	pets := []string{"Charlie", "Garth"}
	//bool
	iscool := true
	fmt.Println("Hello,"+name+"Your'e", age, "years old now!"+"Your'e ", height, "feet tall!"+"Your gender is: "+gender+"A hand has 5 fingers: ", fingers, "You have 2 pets: ", pets, "Is Coleman cool?: ", iscool)

}
