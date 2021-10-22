package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	// Declare the seed value in order to make the random number chage overtime
	rand.Seed(time.Now().UTC().UnixNano())

	// Check if the file exist, if not will create the new file
	if _, err := os.Stat("digits.txt"); err != nil {
		file, err := os.Create("digits.txt")
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		return
	}
	// If file exist, clear the content in the file
	err := os.Truncate("digits.txt", 0)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	writeRandomDigitsToFile(2048)
	readDigitsFromFile()
}

// writeRandomDigitsToFile - This function will write random number in to the file
func writeRandomDigitsToFile(length int) {
	// Open the digits file with Write-only permission
	file, err := os.OpenFile("digits.txt", os.O_WRONLY, 0222)
	if err != nil {
		log.Fatal(err.Error())
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

// readDigitsFromFile - read the file and print the even/odd result
func readDigitsFromFile() {
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
	return
}
