package function

import "fmt"

// printEvenOdd - Check and print if the number is even or odd
func printEvenOdd(n *int) {
	// Check if text is nil
	if n == nil {
		print("cannot print even/odd: argument is nil")
		return
	}

	// Check if the number is even or odd and print the number
	if *n == 0 || *n%2 == 0 {
		fmt.Printf("%d is is an even number\n", *n)
		return
	}
	fmt.Printf("%d is is an odd number\n", *n)
}
