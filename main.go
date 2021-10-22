package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	// declare the seed value in order to make the random number chage overtime
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
}

// writeRandomDigitsToFile - This function will write random number in to the file
func writeRandomDigitsToFile(length int) {
	// Open the digits file with Write only permission
	file, err := os.OpenFile("digits.txt", os.O_WRONLY, 0700)
	if err != nil {
		log.Fatal(err.Error())
	}
	//
	defer file.Close()

	// to generate the 2048 digits
	for i := 0; i < length; i++ {
		// to random digit in range 0 - 9
		n := strconv.Itoa(rand.Intn(10))
		file.WriteString(n)
	}
}
