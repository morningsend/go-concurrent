package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

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

func reverse(a []int) {
	if len(a) < 2 {
		return
	}
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

func SameSequence(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if a == nil && b == nil {
		return true
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestStackConcurrently(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)

	stack := NewStack()

	n := 10000

	even := make([]int, n/2, n/2)
	odd := make([]int, n/2, n/2)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even[i/2] = i
		} else {
			odd[i/2] = i
		}
	}

	pushNumbers := func(s *Stack, nums []int) {
		for _, n := range nums {
			s.Push(n)
		}
	}

	fmt.Printf("even %v, odd %v", even, odd)
	go func() {
		pushNumbers(stack, even)
		wg.Done()
	}()

	go func() {
		pushNumbers(stack, odd)
		wg.Done()
	}()

	wg.Wait()

	if stack.Len() != n {
		t.Fatal("stack lost value")
	}

	evenResult := make([]int, 0, n/2)
	oddResult := make([]int, 0, n/2)

	for i := 0; i < n; i++ {
		x, ok := stack.Pop()
		if !ok {
			t.Fatal("stack emptied early")
		}
		if b := ((x).(int)); b%2 == 0 {
			evenResult = append(evenResult, b)
		} else {
			oddResult = append(oddResult, b)
		}
	}

	reverse(evenResult)
	reverse(oddResult)

	if !SameSequence(even, evenResult) {
		t.Errorf("wrong result expect %v, got %v", even, evenResult)
	}

	if !SameSequence(odd, oddResult) {
		t.Errorf("wrong result expect %v, got %v", odd, oddResult)
	}
}
