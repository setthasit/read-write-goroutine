package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"read-digit/function"
	"strconv"
	"time"
)

var filePath = "digits.txt"

func init() {
	// remove prefix & flag form log output
	log.SetFlags(0)
	// Declare the seed value in order to make the random number chage overtime
	rand.Seed(time.Now().UTC().UnixNano())

	// Check if the file exist, if not will create the new file
	if _, err := os.Stat(filePath); err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		return
	}
	// If file exist, clear the content in the file
	err := os.Truncate(filePath, 0)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	// Write the digits to file
	function.WriteRandomDigitsToFile(filePath, 2048)

	// Read digit from file
	file, scanner := function.ReadDigitsFromFile(filePath)
	// Close the file when the function end
	defer file.Close()

	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		text := scanner.Text()
		// Convert text into int
		n, err := strconv.Atoi(text)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		function.PrintEvenOdd(&n)
	}
}
