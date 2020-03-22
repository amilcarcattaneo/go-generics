package libs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"generics/src/libs"
)

func Test_Map_Success(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When

	users := []struct {
		ID int
	}{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
		{
			ID: 3,
		},
	}

	var ids []int
	libs.Map(len(users), func(i int) {
		ids = append(ids, users[i].ID)
	})

	// Then
	ass.Len(ids, 3)
	for i := 0; i < 3; i++ {
		ass.Equal(i+1, ids[i])
	}
}
