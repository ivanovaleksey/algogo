package heapsort_test

import (
	"github.com/ivanovaleksey/sort-go/heapsort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntHeap_Sort(t *testing.T) {
	a := []int{3, 2, 7, 10, 8, 5, 19, 4, 3}
	h := heapsort.NewIntHeap(a)

	h.Sort()

	expected := []int{2, 3, 3, 4, 5, 7, 8, 10, 19}
	assert.Equal(t, heapsort.NewIntHeap(expected), h)
}
