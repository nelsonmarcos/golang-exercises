package main

import "fmt"

func main() {
	var text, reverse string
	fmt.Println("Enter a string:")
	fmt.Scanln(&text)

	for index := 0; index < len(text); index++ {
		reverse = text[index]
	}

	if text == reverse {
		fmt.Println("It's a palindrome")
	}

}
