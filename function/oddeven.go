package function

import (
	"log"
)

// PrintEvenOdd - Check and print if the number is even or odd
func PrintEvenOdd(n *int) {
	// Check if text is nil
	if n == nil {
		log.Print("cannot print even/odd: argument is nil")
		return
	}

	// Check if the number is even or odd and print the number
	if *n == 0 || *n%2 == 0 {
		log.Printf("%d is an even number\n", *n)
		return
	}
	log.Printf("%d is an odd number\n", *n)
}
