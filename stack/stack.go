package stack

type StackItem struct {
	item interface{}
	next *StackItem
}

type Stack struct {
	top *StackItem
	depth uint64
}

func New() *Stack {
	s := new(Stack)
	s.depth = 0
	return s
}

func (s *Stack) Push(item interface{}) {
	stackItem := StackItem{item, s.top}
	s.top = &stackItem
	s.depth++
}

func (s *Stack) Pop() (item interface{}) {
	if s.depth > 0 {
		item := s.top.item
		s.top = s.top.next
		s.depth--
		return item
	}

	return nil
}