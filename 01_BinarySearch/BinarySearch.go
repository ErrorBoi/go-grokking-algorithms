package _1_BinarySearch

import (
	"testing"
)

func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 5, 6, 9}
	t.Run("Item was found by index 3", func(t *testing.T) {
		item := 6
		if res := BinarySearch(arr, item); res != 3 {
			t.Errorf("Item was found by index %d, want 3", res)
		}
	})
}
