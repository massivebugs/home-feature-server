package util

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		items: map[T]bool{},
	}
}

func (s *Set[T]) Add(e T) {
	s.items[e] = true
}

func (s *Set[T]) Remove(e T) {
	delete(s.items, e)
}

func (s *Set[T]) Exists(e T) bool {
	_, ok := s.items[e]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

// Returns a slice containing the items added to Set.
// Items are NOT in insertion order, but random.
func (s *Set[T]) ToSlice() []T {
	sl := make([]T, len(s.items))

	idx := 0
	for k := range s.items {
		sl[idx] = k
		idx++
	}

	return sl
}
