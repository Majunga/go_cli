package utils

import "testing"

func TestReduce(t *testing.T) {
	t.Run("should return the same slice if the predicate is false for all elements", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		predicate := func(s string) bool {
			return false
		}

		result := Reduce(slice, predicate)

		if len(result) != 3 {
			t.Errorf("Expected the same slice, but got %v", result)
		}
	})

	t.Run("should return a slice without the element that matches the predicate", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		predicate := func(s string) bool {
			return s == "b"
		}

		result := Reduce(slice, predicate)

		if len(result) != 2 {
			t.Errorf("Expected a slice without 'b', but got %v", result)
		}
	})
}
