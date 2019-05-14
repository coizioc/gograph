package main

// Stack is the struct representing the stack.
type Stack struct {
	slist []int
}

// NewStack initializes a new stack.
func NewStack() *Stack {
	return &Stack{[]int{}}
}

// Push pushes a new int to the stack.
func (s *Stack) Push(x int) {
	s.slist = append(s.slist, x)
}

// Pop pops an int from the stack.
func (s *Stack) Pop() int {
	x := s.slist[len(s.slist)-1]
	s.slist = s.slist[0 : s.Size()-1]
	return x
}

// Size returns the size of the stack.
func (s *Stack) Size() int {
	return len(s.slist)
}
