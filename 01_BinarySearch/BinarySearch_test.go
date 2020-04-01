package _1_BinarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 5, 6, 9}

	t.Run("Item was found by index 3", func(t *testing.T) {
		item := 6
		if res := BinarySearch(arr, item); res != 3 {
			t.Errorf("Item was found by index %d, want 3", res)
		}
	})
	t.Run("Item was found by index 0", func(t *testing.T) {
		item := 1
		if res := BinarySearch(arr, item); res != 0 {
			t.Errorf("Item was found by index %d, want 0", res)
		}
	})
	t.Run("Item was not found", func(t *testing.T) {
		item := 3
		if res := BinarySearch(arr, item); res != -1 {
			t.Errorf("Item was found by index %d, want -1 (Not Found)", res)
		}
	})
}
