// Branching with if and esle in Go is straight-forward
// There is not ternary if in Go, so you'll need to
// use a full if statement even for basic conditions.

package main

import "fmt"

func main() {
	// Here is a basic example of a if else branch
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can even have an if statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisable by 4")
	}

	// Logic operators like && and || are often useful in
	// conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statment can precede conditionals; any variables
	// declared in this statment are available in the current
	// and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
