package stacks

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minStack) == 0 || val < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	n := len(this.stack)

	if n == 0 {
		return
	}

	this.stack = this.stack[:n-1]
	this.minStack = this.minStack[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.stack)

	if n == 0 {
		return 0
	}

	return this.stack[n-1]
}

func (this *MinStack) GetMin() int {
	n := len(this.stack)

	if n == 0 {
		return 0
	}

	return this.minStack[n-1]
}
