package function

import "strconv"

// printEvenOdd - Check and print if the number is even or odd
func printEvenOdd(text *string) {
	// Check if text is nil
	if text == nil {
		print("cannot print even/odd: argument is nil")
		return
	}

	// Convert the text into int
	n, err := strconv.Atoi(*text)
	if err != nil {
		println(*text + " is not a number")
	}

	// Check if the number is even or odd and print the number
	if n == 0 || n%2 == 0 {
		println(*text + " is is an even number")
		return
	}
	println(*text + " is is an odd number")
}
