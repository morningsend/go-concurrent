package concurrent

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	if _, ok := s.Pop(); ok {
		t.Errorf("expect empty stack")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if l := s.Len(); l != 3 {
		t.Errorf("expect len to be %v but got %v", 3, l)
	}
	if x, _ := s.Pop(); x != 3 {
		t.Errorf("Pop expect %v but got %v", 3, x)
	}

	if x, _ := s.Pop(); x != 2 {
		t.Errorf("Pop expect %v but got %v", 3, x)
	}
	if x, _ := s.Pop(); x != 1 {
		t.Errorf("Pop expect %v but got %v", 3, x)
	}

	if _, ok := s.Pop(); ok {
		t.Errorf("expect empty stack")
	}

}
