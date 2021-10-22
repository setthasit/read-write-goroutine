package function

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(0)
}

func TestEvenNumber(t *testing.T) {
	testCase := []int{-2, 0, 2, 4, 6, 8}
	for _, n := range testCase {
		output := captureOutput(func() {
			printEvenOdd(&n)
		})
		expected := fmt.Sprintf("%d is is an even number\n", n)
		assert.Equal(t, expected, output)
	}
}

func TestOddNumber(t *testing.T) {
	testCase := []int{-1, 1, 3, 5, 7, 9}
	for _, n := range testCase {
		output := captureOutput(func() {
			printEvenOdd(&n)
		})
		expected := fmt.Sprintf("%d is is an odd number\n", n)
		assert.Equal(t, expected, output)
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
