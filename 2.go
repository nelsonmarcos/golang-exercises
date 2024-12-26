// This solution intend to be as simple as possible, for educational purposes.
// Although there are plenty of ways to optimize this algorithm,
// these optimizations are out of the scope of this solution.

package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)
	fmt.Print("Enter the third number: ")
	fmt.Scanln(&number3)

	if number1 > number2 && number1 > number3 {
		fmt.Println(number1, "is the biggest number")
	} else if number2 > number1 && number2 > number3 {
		fmt.Println(number2, "is the biggest number")
	} else if number3 > number1 && number3 > number2 {
		fmt.Println(number3, "is the biggest")
	}

}
