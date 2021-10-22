package main

import (
	"log"
	"math/rand"
	"os"
	"read-digit/function"
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
	function.WriteRandomDigitsToFile(2048)
	function.ReadDigitsFromFile()
}
