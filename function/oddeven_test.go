package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenNumber(t *testing.T) {
	testCase := []int{-2, 0, 2, 4, 6, 8}
	for _, n := range testCase {
		output := captureOutput(func() {
			PrintEvenOdd(&n)
		})
		expected := fmt.Sprintf("%d is is an even number\n", n)
		assert.Equal(t, expected, output)
	}
}

func TestOddNumber(t *testing.T) {
	testCase := []int{-1, 1, 3, 5, 7, 9}
	for _, n := range testCase {
		output := captureOutput(func() {
			PrintEvenOdd(&n)
		})
		expected := fmt.Sprintf("%d is is an odd number\n", n)
		assert.Equal(t, expected, output)
	}
}

func TestNilNumber(t *testing.T) {
	output := captureOutput(func() {
		PrintEvenOdd(nil)
	})
	expected := "cannot print even/odd: argument is nil\n"
	assert.Equal(t, expected, output)
}
