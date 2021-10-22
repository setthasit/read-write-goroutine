package function

import (
	"bufio"
	"os"
)

// ReadDigitsFromFile - read the file and print the even/odd result
func ReadDigitsFromFile(filePath string) (*os.File, *bufio.Scanner) {
	// Open the digits file with Read-only permission
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err.Error())
	}

	// Use scanner to read file content
	scanner := bufio.NewScanner(file)

	return file, scanner
}
