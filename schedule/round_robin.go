package schedule

// RoundRobin is a scheduling algorithm that traverse elements by round cycle.
type RoundRobin[T any] struct {
	elements []T
	index    int
}

func NewRoundRobin[T any](elements []T) *RoundRobin[T] {
	return &RoundRobin[T]{
		elements: elements,
	}
}

func (r *RoundRobin[T]) Next() T {
	defer func() {
		r.index++
		if r.index == len(r.elements) {
			r.index = 0
		}
	}()

	return r.elements[r.index]
}

// Check that RoundRobin implements Scheduler
var _ Scheduler[any] = &RoundRobin[any]{}
