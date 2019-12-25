package adjacent_area_test

import (
	"github.com/ivanovaleksey/algogo/adjacent_area"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected int
		name     string
	}{
		{
			matrix: [][]int{
				{0, 1, 1, 0, 1},
				{1, 0, 1, 0, 0},
				{0, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 0, 1, 1},
			},
			expected: 4,
			name:     "with plain square matrix",
		},
		{
			matrix:   [][]int{},
			expected: 0,
			name:     "with empty matrix",
		},
		{
			matrix: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 1,
			name:     "with all 1",
		},
		{
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
			name:     "with all 0",
		},
		{
			matrix: [][]int{
				{1, 1, 0, 1},
			},
			expected: 2,
			name:     "with single row",
		},
		{
			matrix: [][]int{
				{1},
				{0},
				{1},
			},
			expected: 2,
			name:     "with single column",
		},
		{
			matrix: [][]int{
				{1, 1, 1, 1, 1},
				{0, 0, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 0, 1, 1, 1},
			},
			expected: 2,
			name:     "with a kind of loop",
		},
		{
			matrix: [][]int{
				{1, 0, 1, 0, 1},
				{0, 1, 1, 1},
				{0, 0, 1, 0, 0},
				{1, 1, 0},
			},
			expected: 4,
			name:     "with different sized rows",
		},
		{
			matrix: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
				{0, 1, 0},
			},
			expected: 6,
			name:     "with diagonal adjacency",
		},
	}

	for _, tt := range testCases {
		calc := adjacent_area.NewCalculator(tt.matrix)

		num := calc.Calculate()

		assert.Equal(t, tt.expected, num, tt.name)
	}
}
