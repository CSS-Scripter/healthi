package fn

import "testing"

func Test_MapKeys(t *testing.T) {
	m := Map[string, int]{
		"hello": 12,
		"world": 1,
	}

	keys := m.Keys()

	if len(keys) != 2 {
		t.Fatalf("expected a length of 2, got %d", len(keys))
	}

	if keys[0] != "hello" {
		t.Fatalf("expected the first key to be 'hello', got %s", keys[0])
	}

	if keys[1] != "world" {
		t.Fatalf("expected the second key to be 'world', got %s", keys[1])
	}
}

func Test_MapValues(t *testing.T) {
	m := Map[string, int]{
		"hello": 12,
		"world": 1,
	}

	values := m.Values()

	if len(values) != 2 {
		t.Fatalf("expected a length of 2, got %d", len(values))
	}

	if values[0] != 12 {
		t.Fatalf("expected the first value to be 12, got %d", values[0])
	}

	if values[1] != 1 {
		t.Fatalf("expected the second value to be 1, got %d", values[1])
	}
}

func Test_MapEntries(t *testing.T) {
	m := Map[string, int]{
		"hello": 12,
		"world": 1,
	}

	entries := m.Entries()

	if len(entries) != 2 {
		t.Fatalf("expected a length of 2, got %d", len(entries))
	}

	if entries[0].Key != "hello" {
		t.Fatalf("expected the first key to be 'hello', got %s", entries[0].Key)
	}

	if entries[0].Value != 12 {
		t.Fatalf("expected the first value to be 12, got %d", entries[0].Value)
	}

	if entries[1].Key != "world" {
		t.Fatalf("expected the second key to be 'world', got %s", entries[1].Key)
	}

	if entries[1].Value != 1 {
		t.Fatalf("expected the second value to be 1, got %d", entries[1].Value)
	}
}
