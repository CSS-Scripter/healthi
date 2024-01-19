package fn

type Map[T comparable, R any] map[T]R
type Entry[T comparable, R any] struct {
	Key   T
	Value R
}

func (m Map[T, R]) Keys() Slice[T] {
	s := make(Slice[T], len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	return s
}

func (m Map[T, R]) Values() Slice[R] {
	s := make(Slice[R], len(m))
	i := 0
	for _, v := range m {
		s[i] = v
		i++
	}
	return s
}

func (m Map[T, R]) Entries() Slice[Entry[T, R]] {
	s := make(Slice[Entry[T, R]], len(m))
	i := 0
	for k, v := range m {
		s[i] = Entry[T, R]{k, v}
		i++
	}
	return s
}
