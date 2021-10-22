package function

import (
	"bufio"
	"log"
	"os"
)

// ReadDigitsFromFile - read the file and print the even/odd result
func ReadDigitsFromFile() {
	// Open the digits file with Read-only permission
	file, err := os.OpenFile("digits.txt", os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Close the file when the function end
	defer file.Close()

	// Use scanner to read file content
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		text := scanner.Text()
		printEvenOdd(&text)
	}
}
