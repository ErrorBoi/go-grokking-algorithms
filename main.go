package main

import (
	"fmt"

	bs "github.com/ErrorBoi/go-grokking-algorithms/01_BinarySearch"
)

func main() {
	testBinarySearch()
}

func testBinarySearch() {
	arr := []int{1, 2, 5, 6, 9}
	item := 6
	res := bs.BinarySearch(arr, item)
	fmt.Printf("Item %d is located at index %d\n", item, res)
}
