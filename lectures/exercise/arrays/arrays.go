//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//  information again

package main

import "fmt"

//  - Products must include the price and the name
type Products struct {
	price int
	name string
}

func printStats(list [4]Products) {
	var cost, totalItems int

	for i := 0; i < len(list); i++{
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}
	//  - The total cost of the items
	fmt.Println("The total cost is", cost)
	//  - The total number of items
	fmt.Println("The total number of items is", totalItems)
	//  - The last item on the list
	fmt.Println("The last item in the list is", list[totalItems - 1].name)
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//* Insert 3 products into the array
	shoppingList := [4]Products{
		{5, "oranges"},
		{10, "steak"},
		{4, "milk"},
	}

	//* Print to the terminal:
	printStats(shoppingList)

	//* Add a fourth product to the list and print out the
	shoppingList[3]= Products{7, "bread"}

	printStats(shoppingList)
}

