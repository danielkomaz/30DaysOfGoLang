package main

import ("fmt"
		"greeting"
)

func main() {
	// Get greeting message and print it.
	message := greeting.Hello("Daniel")
	fmt.Println(message)
}
