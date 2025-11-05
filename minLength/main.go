package main

func main() {

}

func minLength(s string) int {
	charSt := Stack{
		make([]rune, 0, len(s)),
	}

	for _, c := range s {
		if charSt.IsEmpty() {
			charSt.Push(c)
		} else if (charSt.Peek() == 'C' && c == 'D') || (charSt.Peek() == 'A' && c == 'B') {
			charSt.Pop()
		} else {
			charSt.Push(c)
		}
	}

	return charSt.Len()
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(val rune) {
	s.items = append(s.items, val)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() rune {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return -1
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Len() int {
	return len(s.items)
}
