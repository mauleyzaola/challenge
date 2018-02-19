package storage

import (
	"testing"

	"github.com/mauleyzaola/challenge/operations"
)

func TestMemoryOperations(t *testing.T) {
	var storage operations.Storage
	storage = &Memory{}
	storage.Init()

	id, err := storage.Create()
	if err != nil {
		t.Error("unexpected error:", err)
	}

	basket, err := storage.Load("xx")
	if err == nil {
		t.Error("expected error but got nil instead")
	}
	basket, err = storage.Load(id)
	if err != nil {
		t.Error(err)
	}
	if basket == nil {
		t.Error("expected basket not to be nil but got nil instead")
	}
	list := storage.List()
	expected := 1
	if len(list) != expected {
		t.Errorf("expected list of baskets to be:%d but got instead:%d", expected, len(list))
	}

	if err = storage.Remove(id); err != nil {
		t.Error(err)
	}

	list = storage.List()
	expected = 0
	if len(list) != expected {
		t.Errorf("expected list of baskets to be:%d but got instead:%d", expected, len(list))
	}
}
