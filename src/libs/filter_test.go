package libs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"generics/src/libs"
)

func Test_Filter_Success(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers []int

	libs.Filter(len(numbers), func(i int) bool {
		return numbers[i]%2 == 0
	}, func(i int) {
		evenNumbers = append(evenNumbers, numbers[i])
	})

	// Then
	ass.Len(evenNumbers, 5)
}

func Test_Filter_NoOP(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	numbers := []int{1, 3, 5, 7, 9}
	var evenNumbers []int

	libs.Filter(len(numbers), func(i int) bool {
		return numbers[i]%2 == 0
	}, func(i int) {
		evenNumbers = append(evenNumbers, numbers[i])
	})

	// Then
	ass.Len(evenNumbers, 0)
}
