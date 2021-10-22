package function

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func checkfile(filePath string) {
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

func TestWriteLenght(t *testing.T) {
	testLenght := 10
	filePath := "TestWriteLenght.txt"

	checkfile(filePath)
	defer os.Remove(filePath)

	WriteRandomDigitsToFile(filePath, testLenght)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	assert.Equal(t, testLenght, len(string(file)))
}

func TestFileNotfound(t *testing.T) {
	assert.Panics(
		t,
		func() { WriteRandomDigitsToFile("randomfile.txt", 1) },
		"code should be panic",
	)
}
