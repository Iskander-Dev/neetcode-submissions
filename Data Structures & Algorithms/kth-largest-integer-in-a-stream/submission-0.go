type KthLargest struct {
    heap []int
	pos int
}


func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		heap: nums,
		pos: k,
	}
}

func (this *KthLargest) heapSort() {
	n := len(this.heap)/2 - 1

	for i := n; i >= 0; i-- {
		this.heapifyDown(len(this.heap), i)
	}

	for i := len(this.heap) - 1; i > 0; i-- {
		this.heap[0], this.heap[i] = this.heap[i], this.heap[0]
		this.heapifyDown(i, 0)
	}
}

func (this *KthLargest) heapifyDown(span, i int) {
	for {
		largest := i
		left := 2*i+1
		right := 2*i+2

		if left < span && this.heap[left] > this.heap[largest] {
			largest = left
		}
		if right < span && this.heap[right] > this.heap[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		this.heap[i], this.heap[largest] = this.heap[largest], this.heap[i]
		i = largest
	}
}

func (this *KthLargest) Add(val int) int {
    this.heap = append(this.heap, val)
	this.heapSort()
	return this.heap[len(this.heap)-this.pos]
}
