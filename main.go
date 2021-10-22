package main

import (
	"log"
	"math/rand"
	"os"
	"read-digit/function"
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
	function.WriteRandomDigitsToFile(filePath, 2048)
	function.ReadDigitsFromFile()
}
