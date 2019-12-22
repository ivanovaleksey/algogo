package heapsort

type IntHeap []int

func NewIntHeap(a []int) IntHeap {
	return a
}

func (h *IntHeap) Sort() {
	h.init()

	for i := len(*h) - 1; i > 0; i-- {
		h.swap(0, i)
		old := *h
		old = old[:i]
		old.heapify(0)
	}
}

func (h *IntHeap) init() {
	idx := len(*h) - 1
	for i := (idx - 1) / 2; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *IntHeap) heapify(i int) {
	maxIndex := i
	left := leftChild(i)
	right := rightChild(i)
	size := len(*h)

	if left <= size-1 && (*h)[left] > (*h)[maxIndex] {
		maxIndex = left
	}

	if right <= size-1 && (*h)[right] > (*h)[maxIndex] {
		maxIndex = right
	}

	if maxIndex != i {
		h.swap(i, maxIndex)
		h.heapify(maxIndex)
	}
}

func (h *IntHeap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
