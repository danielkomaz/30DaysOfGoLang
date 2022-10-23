package main

import "fmt"

func main() {
	var input string
	fmt.Println("Enter string to reverse:")
	fmt.Scanln(&input)

	chars := []rune(input)
	var reverseString string

	for i := len(chars)-1 ; i >= 0 ; i-- {
		char := string(chars[i])
		reverseString += char
	}

	fmt.Println(reverseString)
}