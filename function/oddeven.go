package function

import (
	"log"
)

// printEvenOdd - Check and print if the number is even or odd
func printEvenOdd(n *int) {
	// Check if text is nil
	if n == nil {
		log.Print("cannot print even/odd: argument is nil")
		return
	}

	// Check if the number is even or odd and print the number
	if *n == 0 || *n%2 == 0 {
		log.Printf("%d is is an even number\n", *n)
		return
	}
	log.Printf("%d is is an odd number\n", *n)
}
