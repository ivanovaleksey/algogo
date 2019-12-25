package quicksort_test

import (
	"github.com/ivanovaleksey/algogo/quicksort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		array    []int
		expected []int
		name     string
	}{
		{
			array:    []int{3, 2, 7, 10, 8, 5, 19, 4, 3},
			expected: []int{2, 3, 3, 4, 5, 7, 8, 10, 19},
			name:     "case 1",
		},
		{
			array:    []int{45, 49, 18, 98, 56, 63, 61, 1, 97},
			expected: []int{1, 18, 45, 49, 56, 61, 63, 97, 98},
			name:     "case 2",
		},
		{
			array:    []int{1, 20, 40, 84, 16, 6, 90, 77, 29},
			expected: []int{1, 6, 16, 20, 29, 40, 77, 84, 90},
			name:     "case 3",
		},
	}

	for _, tt := range testCases {
		quicksort.Sort(tt.array)

		assert.Equal(t, tt.expected, tt.array, tt.name)
	}
}
