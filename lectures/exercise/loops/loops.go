//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//* Print integers 1 to 50, except:
	// for i := 1; i <= 50; i++ {
	// 	//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		// 	if i % 3 == 0 && i % 5 == 0 {
		// 		fmt.Println("FizzBuzz")
		// 	} else if i % 3 == 0 {
			// 		//  - Print "Fizz" if the integer is divisible by 3
			// 		fmt.Println("Fizz")
		// 	} else if i % 5 == 0 {
				// 		//  - Print "Buzz" if the integer is divisible by 5
				// 		fmt.Println("Buzz")
		// 	} else {
		// 		fmt.Println(i)
	// 	}
// }

	for i := 1; i <= 50; i++ {
		divisibleby3 := i%3 == 0
		divisibleby5 := i%5 == 0

		if divisibleby3 && divisibleby5 {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FizzBuzz")
		} else if divisibleby3 {
			// 		//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println("Fizz")
		} else if divisibleby5 {
	// 		//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}

