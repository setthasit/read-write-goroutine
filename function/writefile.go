package function

import (
	"math/rand"
	"os"
	"strconv"
)

// WriteRandomDigitsToFile - This function will write random number in to the file
func WriteRandomDigitsToFile(filePath string, length int) {
	// Open the digits file with Write-only permission
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0222)
	if err != nil {
		panic(err.Error())
	}
	// Close the file when the function end
	defer file.Close()

	// To generate the 2048 digits
	for i := 0; i < length; i++ {
		// to random digit in range 0 - 9
		n := strconv.Itoa(rand.Intn(10))
		file.WriteString(n)
	}
}
