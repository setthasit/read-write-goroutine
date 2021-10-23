package function

import (
	"sync"
	"testing"

	"io/ioutil"
	"os"

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

	var mu sync.RWMutex
	for idx, txt := range testText {
		n := ReadDigitsFromFilePosition(filePath, int64(idx), &mu)
		// int(txt-'0') is for convert rune to int
		assert.Equal(t, int(txt-'0'), n)
	}

}

func TestReadFileNotfound(t *testing.T) {
	var mu sync.RWMutex
	assert.Panics(
		t,
		func() { ReadDigitsFromFilePosition("randomfile.txt", 0, &mu) },
		"code should be panic",
	)
}
