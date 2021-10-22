package function

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDigit(t *testing.T) {
	filePath := "TestReadDigit.txt"
	testText := "987654321"

	file, err := os.Create(filePath)
	if err != nil {
		t.Error(err.Error())
	}
	file.Close()
	defer os.Remove(filePath)

	err = ioutil.WriteFile(filePath, []byte(testText), 0222)
	if err != nil {
		t.Error(err.Error())
	}

	file, scanner := ReadDigitsFromFile(filePath)
	defer file.Close()

	scanner.Scan()
	assert.Equal(t, testText, scanner.Text())
}

func TestReadFileNotfound(t *testing.T) {
	assert.Panics(
		t,
		func() { ReadDigitsFromFile("randomfile.txt") },
		"code should be panic",
	)
}
