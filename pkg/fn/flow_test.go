package fn

import (
	"errors"
	"testing"
)

func TestFlow(t *testing.T) {
	value := 2
	errorMessage := "Whoops"
	flow := Flow{Slice[func() error]{
		func() error { value += 2; return nil },
		func() error { value += 6; return nil },
		func() error { return errors.New(errorMessage) },
		func() error { value = value * 2; return nil },
	}}

	err := flow.Run()
	if err == nil {
		t.Fatal("expected error to be defined, but got nil")
	}

	if err.Error() != errorMessage {
		t.Fatalf("expected error to have message %s but got %s", errorMessage, err.Error())
	}

	if value != 10 {
		t.Fatalf("expected outcome value of 10 but got %d", value)
	}
}
