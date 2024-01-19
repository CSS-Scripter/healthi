package fn

type Map[T comparable, R any] map[T]R
type Entry[T comparable, R any] struct {
	Key   T
	Value R
}

func (m Map[T, R]) Keys() Slice[T] {
	s := Slice[T]{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

func (m Map[T, R]) Values() Slice[R] {
	s := Slice[R]{}
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func (m Map[T, R]) Entries() Slice[Entry[T, R]] {
	s := Slice[Entry[T, R]]{}
	for k, v := range m {
		s = append(s, Entry[T, R]{k, v})
	}
	return s
}
