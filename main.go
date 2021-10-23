package main

import (
	"log"
	"math/rand"
	"os"
	"read-digit/function"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var filePath = "digits.txt"
var digitsLength int64

func init() {
	// Check if CPU >= 2 cores
	if runtime.NumCPU() >= 2 {
		runtime.GOMAXPROCS(2)
	}

	// Remove prefix & flag form log output
	log.SetFlags(0)

	// Declare the seed value in order to make the random number chage overtime
	rand.Seed(time.Now().UTC().UnixNano())

	// Check if the file exist, if not will create the new file
	// If file already exist, will clear the content in the file
	if _, err := os.Stat(filePath); err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		return
	} else {
		err := os.Truncate(filePath, 0)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// Check if there command argument
	// argument should be number
	digitsLength = 2048
	if os.Args[1] != "" {
		argLength, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		if argLength > 0 {
			digitsLength = int64(argLength)
		}
	}
}

func main() {
	// Mutex use for preventing race condition
	var mu sync.RWMutex
	// Channel that contain position of new digits
	changes := make(chan int64, digitsLength)

	// Write the digits to file concurently
	go function.WriteRandomDigitsToFile(filePath, digitsLength, changes, &mu)

	// Read the digits and print to console
	counts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for change := range changes {
		n := function.ReadDigitsFromFilePosition(filePath, change, &mu)
		function.PrintEvenOdd(&n)

		counts[n] += 1
	}

	// Summary of digits in the file
	log.Println("#############################")
	log.Println("Each digit count in the file")
	log.Println("#############################")
	var totalCount = 0
	for idx, count := range counts {
		log.Printf("%d has %d\n", idx, count)
		totalCount += count
	}
	log.Printf("Total digits counts in the file is %d\n", totalCount)
}
