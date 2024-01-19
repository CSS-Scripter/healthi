package fn

import "testing"

func TestSliceForEach(t *testing.T) {
	s := Slice[int]{1, 2, 3, 4, 5}
	lastCalc := 0
	s.ForEach(func(i int) bool {
		lastCalc = i * 2
		return lastCalc < 5
	})

	if lastCalc != 6 {
		t.Fatalf("expected a last value of 6, got %d", lastCalc)
	}
}

func TestSliceMap(t *testing.T) {
	s := Slice[int]{1, 2, 3, 4, 5}
	s.Map(func(i int) int {
		return i * 2
	})

	expected := []int{2, 4, 6, 8, 10}

	for i := range s {
		if s[i] != expected[i] {
			t.Fatalf("expected %d at position %d, but got %d", expected[i], i, s[i])
		}
	}
}
