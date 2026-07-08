type Point struct {
	dis    int
	coords []int
}

type MinHeap struct {
	heap []Point
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		heap: make([]Point, 0, capacity),
	}
}

func (h *MinHeap) parent(i int) int { return (i - 1) / 2 }

func (h *MinHeap) leftChild(i int) int { return i*2 + 1 }

func (h *MinHeap) rightChild(i int) int { return i*2 + 2 }

func (h *MinHeap) Len() int { return len(h.heap) }

func (h *MinHeap) isEmpty() bool { return h.Len() == 0 }

func (h *MinHeap) Push(coords []int) {
	distance := coords[0]*coords[0] + coords[1]*coords[1]
	h.heap = append(h.heap, Point{distance, coords})
	h.bubbleUp(h.Len() - 1)
}

func (h *MinHeap) Pop() Point {
	if h.isEmpty() {
		return Point{}
	}
	min := h.heap[0]

	lastIndex := h.Len() - 1
	h.heap[0] = h.heap[lastIndex]
	h.heap = h.heap[:lastIndex]
	h.bubbleDown(0)

	return min
}

func (h *MinHeap) bubbleUp(i int) {
	for i > 0 && h.heap[h.parent(i)].dis > h.heap[i].dis {
		parentIdx := h.parent(i)
		h.heap[parentIdx], h.heap[i] = h.heap[i], h.heap[parentIdx]
		i = parentIdx
	}
}

func (h *MinHeap) bubbleDown(i int) {
	n := h.Len()

	for {
		smallest := i
		leftChild := h.leftChild(i)
		rightChild := h.rightChild(i)

		if leftChild < n && h.heap[smallest].dis > h.heap[leftChild].dis {
			smallest = leftChild
		}

		if rightChild < n && h.heap[smallest].dis > h.heap[rightChild].dis {
			smallest = rightChild
		}

		if smallest == i {
			break
		}

		h.heap[smallest], h.heap[i] = h.heap[i], h.heap[smallest]
		i = smallest
	}
}

func kClosest(points [][]int, k int) [][]int {
	res := [][]int{}
	pointHeap := NewMinHeap(len(points))

	for _, coords := range points {
		pointHeap.Push(coords)
	}

	for range k {
		res = append(res, pointHeap.Pop().coords)
	}

	return res
}
