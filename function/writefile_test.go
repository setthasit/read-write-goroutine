package function

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteLenght(t *testing.T) {
	var testLenght int64 = 10
	changes := make(chan int64, testLenght)
	filePath := "TestWriteLenght.txt"

	checkfile(filePath)
	defer os.Remove(filePath)

	var mu sync.RWMutex
	WriteRandomDigitsToFile(filePath, testLenght, changes, &mu)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	assert.Equal(t, int(testLenght), len(string(file)))
}

func TestWriteFileNotfound(t *testing.T) {
	var mu sync.RWMutex
	changes := make(chan int64)
	assert.Panics(
		t,
		func() { WriteRandomDigitsToFile("randomfile.txt", 1, changes, &mu) },
		"code should be panic",
	)
}
