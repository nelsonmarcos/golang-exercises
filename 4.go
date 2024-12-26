// This solution intend to be as simple as possible, for educational purposes.
// Although there are plenty of ways to optimize this algorithm,
// these optimizations are out of the scope of this solution.

package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter the number of terms:")
	fmt.Scanln(&number)

	sequence := make([]int, number)

	if number > 1 {
		sequence[1] = 1
		for index := 2; index < number; index++ {
			sequence[index] = sequence[index-1] + sequence[index-2]
		}

	}

	fmt.Println(sequence)
}
