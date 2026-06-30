type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: []int{},
	}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return i*2 + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) Peek() int {
	if h.isEmpty() {
		return 0
	}
	return h.heap[0]
}

func (h *MaxHeap) isEmpty() bool {
	return h.Len() == 0
}

func (h *MaxHeap) Len() int {
	return len(h.heap)
}

func (h *MaxHeap) Push(key int) {
	h.heap = append(h.heap, key)
	h.bubbleUp(h.Len() - 1)
}

func (h *MaxHeap) Pop() int {
	if h.isEmpty() {
		return 0
	}

	max := h.heap[0]
	lastIndex := h.Len() - 1
	h.heap[0] = h.heap[lastIndex]
	h.heap = h.heap[:lastIndex]

	h.bubbleDown(0)

	return max
}

func (h *MaxHeap) MaxHeapify(nums []int) {
	h.heap = append([]int(nil), nums...)
	n := h.Len()/2 - 1

	for i := n; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

func (h *MaxHeap) bubbleUp(i int) {
	for i > 0 && h.heap[h.parent(i)] < h.heap[i] {
		parentIdx := h.parent(i)
		h.heap[parentIdx], h.heap[i] = h.heap[i], h.heap[parentIdx]
		i = parentIdx
	}
}

func (h *MaxHeap) bubbleDown(i int) {
	n := h.Len()

	for {
		largest := i
		leftChild := h.leftChild(i)
		rightChild := h.rightChild(i)

		if leftChild < n && h.heap[largest] < h.heap[leftChild] {
			largest = leftChild
		}

		if rightChild < n && h.heap[largest] < h.heap[rightChild] {
			largest = rightChild
		}

		if largest == i {
			break
		}

		h.heap[i], h.heap[largest] = h.heap[largest], h.heap[i]

		i = largest
	}
}


func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap()
	heap.MaxHeapify(stones)

	for heap.Len() > 1 {
		x, y := heap.Pop(), heap.Pop()

		if x == y {
			continue
		}
		heap.Push(x - y)
	}

	return heap.Peek()
}
