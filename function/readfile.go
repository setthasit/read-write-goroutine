package function

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

// ReadDigitsFromFile - read the file and print the even/odd result
func ReadDigitsFromFilePosition(filePath string, position int64, mu *sync.RWMutex) int {
	mu.RLock()
	// Open the digits file with Read-only permission
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file.Seek(position, 0)

	// Use scanner to read file content
	scanner := bufio.NewScanner(file)
	// Scan rune will read only 1 digit
	scanner.Split(bufio.ScanRunes)
	scanner.Scan()

	mu.RUnlock()

	// convert digit string into int
	str := scanner.Text()
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}

	return n
}
