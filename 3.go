// This solution intend to be as simple as possible, for educational purposes.
// Although there are plenty of ways to optimize this algorithm,
// these optimizations are out of the scope of this solution.

package main

import "fmt"

func main() {
	var number, fatorial int = 0, 1
	fmt.Print("Enter the number: ")
	fmt.Scanln(&number)
	for counter := 1; counter <= number; counter++ {
		fatorial = fatorial * counter
	}
	fmt.Println("Factorial of", number, "is", fatorial)

}
