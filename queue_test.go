package concurrent

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()
	_, ok := queue.Dequeue()

	// empty
	if queue.Len() != 0 {
		t.Errorf("empty queue should have size 1 but got %d", queue.Len())
	}
	if ok {
		t.Errorf("cannot dequeue from empty queue")
	}

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		queue.Enqueue(v)
	}

	if queue.Len() != len(values) {
		t.Errorf("len error, expect %v, got %v", len(values), queue.Len())
	}
	results := []int{}

	for v, ok := queue.Dequeue(); ok; v, ok = queue.Dequeue() {
		results = append(results, v.(int))
	}

	if !SameSequence(values, results) {
		t.Errorf(
			"queue content is different than result got %v expect %v",
			results,
			values,
		)
	}

	if queue.Len() != 0 {
		t.Error("expect empty queue")
	}
}
