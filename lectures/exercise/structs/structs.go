//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//* Create a rectangle structure containing a length and width field
type Rectangle struct {
	length int
	width int
}

func area(rect Rectangle) int {
	return rect.length * rect.width
}

func perimeter(rect Rectangle) int {
	return (2 * rect.length) + (2 * rect.width)
}

func printInfo(rect Rectangle) {
	fmt.Println("The area is", area(rect))
	fmt.Println("The perimeter is", perimeter(rect))
}

func main() {
	var rect = Rectangle{length:3, width: 7}
	printInfo(rect)

	rect.length *= 2
	rect.width *= 2

	printInfo(rect)

}

