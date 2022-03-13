package queues

type CQueue struct {
	s1 []int
	s2 []int
}

func NewCQueue() CQueue {
	return CQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	res := -1

	if len(this.s2) != 0 {
		res := this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]

		return res
	}

	if len(this.s1) == 0 {
		return res
	}

	for len(this.s1) > 0 {
		this.s2 = append(this.s2, this.s1[len(this.s1)-1])
		this.s1 = this.s1[:len(this.s1)-1]
	}

	res = this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]

	return res
}
