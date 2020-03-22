package libs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"generics/src/libs"
)

func Test_Reduce_Success_WithArgs(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	items := []struct {
		Amount int
	}{
		{
			Amount: 2,
		},
		{
			Amount: 6,
		},
	}

	sum := 2
	libs.Reduce(len(items), func(i int) {
		sum += items[i].Amount
	}, 0)

	// Then
	ass.Equal(10, sum)
}

func Test_Reduce_Success_WithoutArgs(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	items := []struct {
		Amount int
	}{
		{
			Amount: 2,
		},
		{
			Amount: 6,
		},
	}

	sum := 0
	libs.Reduce(len(items), func(i int) {
		sum += items[i].Amount
	})

	// Then
	ass.Equal(8, sum)
}

func Test_Reduce_Success_With_Index_Arg(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	items := []struct {
		Amount int
	}{
		{
			Amount: 2,
		},
		{
			Amount: 6,
		},
	}

	sum := 0
	libs.Reduce(len(items), func(i int) {
		sum += items[i].Amount
	}, 1)

	// Then
	ass.Equal(6, sum)
}
