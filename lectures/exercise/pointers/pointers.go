//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

const (
	Active = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("Checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {

	pants := Item{"Pants", Active}
	shirt := Item{"Shirt", Active}
	purse := Item{"Purse", Active}
	watch := Item{"Watch", Active}

	items := []Item{shirt, pants, purse, watch}
	fmt.Println(items)

	deactivate(&items[0].tag)
	fmt.Println(items)

	checkout(items)
	fmt.Println("checked out", items)
}
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
// type Item struct {
// 	name string
// 	active bool
// }

// //* Create functions to activate and deactivate security tags using pointers
// func changeState(item *Item) {
// 	item.active = false
// }

// //* Create a checkout() function which can deactivate all tags in a slice
// func checkout(items []Item) {
// 	fmt.Println("Checking out...")
// 	for _, item := range items {
// 		changeState(&item)
// 	}
// }


// func main() {
// 	//* Perform the following:
// 	//  - Create at least 4 items, all with active security tags
// 	var (
// 		shirt = Item{name: "shirt", active: true}
// 		pants = Item{name: "pants", active: true}
// 		wallet = Item{name: "wallet", active: true}
// 		belt = Item{name: "belt", active: true}
// 	)


// 	//  - Store them in a slice or array
// 	items := []Item{shirt, pants, wallet, belt}
// 	fmt.Println(items)

// 	//  - Deactivate any one security tag in the array/slice
// 	changeState(&items[2])
// 	fmt.Println(items)

// 	checkout(items)
// 	fmt.Println(items)
	//  - Call the checkout() function to deactivate all tags
	//  - Print out the array/slice after each change

// }


