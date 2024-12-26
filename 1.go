// This solution intend to be as simple as possible, for educational purposes.
// Although there are plenty of ways to optimize this algorithm,
// these optimizations are out of the scope of this solution.

package main

import "fmt"

func main() {
	var grade1, grade2, grade3, grade4, grade_mean int

	fmt.Println("Enter first grade:")
	fmt.Scanln(&grade1)

	fmt.Println("Enter second grade:")
	fmt.Scanln(&grade2)

	fmt.Println("Enter third grade:")
	fmt.Scanln(&grade3)

	fmt.Println("Enter forth grade:")
	fmt.Scanln(&grade4)

	grade_mean = (grade1 + grade2 + grade3 + grade4) / 4
	fmt.Println("Arithmetic mean:", grade_mean)
}
