package main

import (
	"fmt"
)

//Problem #3
//Assign variables to your first name, last name, age, and full name
//Then ptint the information

func main() {
	fname := "Coleman"
	fmt.Println("First Name:", fname)
	lname := "Word"
	fmt.Println("Last Name:", lname)

	fullname := fname + " " + lname
	fmt.Println("Full Name: ", fullname)

	age := 25
	fmt.Println("Age:", age)
}
