package main
import "fmt"

type Part string

//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
func main() {
	//* Using a slice, create an assembly line that contains type Part
	//  - Create an assembly line having any three parts
	assemblyline := make([]Part, 3)

	assemblyline[0] = "Pipe"
	assemblyline[1] = "Bolt"
	assemblyline[2] = "Sheet"

	fmt.Println("3 parts")
	showLine(assemblyline)

	//  - Add two new parts to the line
	assemblyline = append(assemblyline, "Washer", "Wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyline)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyline = assemblyline[3:]
	fmt.Println("\nSliced")
	showLine(assemblyline)
}

//* Perform the following:
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts




//* Create a function to print out the contents of the assembly line
func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}


