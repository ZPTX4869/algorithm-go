package stack

import "math"

// 记录当前元素和全栈最小值的差值

type MinStack struct {
	diffs []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		diffs: make([]int, 0),
		min:   math.MaxInt,
	}
}

func (s *MinStack) Push(val int) {
	if len(s.diffs) == 0 {
		s.diffs = append(s.diffs, 0)
		s.min = val
		return
	}

	diff := val - s.min
	s.diffs = append(s.diffs, diff)
	if diff < 0 {
		s.min = val
	}
}

func (s *MinStack) Pop() {
	diff := s.diffs[len(s.diffs)-1]
	if diff < 0 {
		s.min = s.min - diff
	}

	s.diffs = s.diffs[0 : len(s.diffs)-1]
}

func (s *MinStack) Top() int {
	diff := s.diffs[len(s.diffs)-1]
	if diff < 0 {
		return s.min
	} else {
		return s.min + diff
	}
}

func (s *MinStack) GetMin() int {
	return s.min
}

// 使用辅助栈

//type MinStack struct {
//	diffs []int
//	mins []int
//}
//
//func Constructor() MinStack {
//	return MinStack{
//		diffs: make([]int, 0),
//	}
//}
//
//func (s *MinStack) Push(val int) {
//	s.diffs = append(s.diffs, val)
//
//	if len(s.mins) == 0 {
//		s.mins = append(s.mins, val)
//		return
//	}
//
//	min := s.mins[len(s.mins)-1]
//	if val < min {
//		min = val
//	}
//
//	s.mins = append(s.mins, min)
//}
//
//func (s *MinStack) Pop() {
//	s.diffs = s.diffs[:len(s.diffs)-1]
//	s.mins = s.mins[:len(s.mins)-1]
//}
//
//func (s *MinStack) Top() int {
//	return s.diffs[len(s.diffs)-1]
//}
//
//func (s *MinStack) GetMin() int {
//	if len(s.diffs) == 0 {
//		panic("Stack is empty!")
//	}
//
//	return s.mins[len(s.mins)-1]
//}
