package main

import (
	"log"
	"math/rand"
	"os"
	"read-digit/function"
	"sync"
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
	// Mutex use for preventing race condition
	var mu sync.RWMutex
	// Channel that contain position of new digits
	changes := make(chan int64, 2048)

	// Write the digits to file concurently
	function.WriteRandomDigitsToFile(filePath, 2048, changes, &mu)
}
