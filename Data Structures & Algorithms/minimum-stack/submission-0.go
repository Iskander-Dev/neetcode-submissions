type element struct {
	val int
	min int
}

type MinStack struct {
	elements []element
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minVal := val

	if len(this.elements) > 0 {
		if currentMin := this.elements[len(this.elements)-1].min; currentMin < minVal {
			minVal = currentMin
		}
	}

	this.elements = append(this.elements, element{val: val, min: minVal})
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}

	this.elements = this.elements[:len(this.elements)-1]
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return 0
	}

	return this.elements[len(this.elements)-1].val
}

func (this *MinStack) GetMin() int {
	if this.isEmpty() {
		return 0
	}

	return this.elements[len(this.elements)-1].min
}

func (this *MinStack) isEmpty() bool {
	return len(this.elements) == 0
}