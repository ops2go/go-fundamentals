package main

import (
	"fmt"
)

/*range iterates over elements in a variety of data structures.
Let’s see how to use range with some of the data structures
we’ve already learned.

range always returns an index which can be ignored with a blank
identifier*/

func main() {

	scores := []int{25, 34, 18}
	sum := 0
	//blank identifier means not using the sum value
	for _, score := range scores {
		sum += score
	}
	fmt.Println("Total Score:", sum)
	fmt.Println("Average Score:", sum/len(scores))
	for i, score := range scores {
		if score == 25 {
			fmt.Println("index:", i)
		}
	}
}
