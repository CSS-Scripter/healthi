package fn

type Slice[T any] []T

func (s Slice[T]) ForEach(fn func(T) bool) {
	for _, i := range s {
		if cont := fn(i); !cont {
			break
		}
	}
}

func (s Slice[T]) Map(fn func(T) T) {
	for i, item := range s {
		s[i] = fn(item)
	}
}
