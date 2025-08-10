//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:

package main

import "fmt"


func main() {
	switch age := 14; {
	case age == 0:
		//  - "newborn" when age is 0
		fmt.Println("newborn")
	case age < 4:
		//  - "toddler" when age is 1, 2, or 3
		fmt.Println("toddler")
	case age < 13:
		//  - "child" when age is 4 through 12
		fmt.Println("child")
	case age < 18:
		//  - "teenager" when age is 13 through 17
		fmt.Println("teenager")
	default:
		//  - "adult" when age is 18+
		fmt.Println("adult")
	}
}

