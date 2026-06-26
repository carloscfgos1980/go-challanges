package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{name: "found middle", arr: []int{1, 3, 5, 7, 9}, target: 5, want: 2},
		{name: "found first", arr: []int{1, 3, 5, 7, 9}, target: 1, want: 0},
		{name: "found last", arr: []int{1, 3, 5, 7, 9}, target: 9, want: 4},
		{name: "not found", arr: []int{1, 3, 5, 7, 9}, target: 6, want: -1},
		{name: "empty array", arr: []int{}, target: 1, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.target)
			if got != tt.want {
				t.Fatalf("BinarySearch() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{name: "found middle", arr: []int{2, 4, 6, 8, 10}, target: 6, want: 2},
		{name: "found first", arr: []int{2, 4, 6, 8, 10}, target: 2, want: 0},
		{name: "found last", arr: []int{2, 4, 6, 8, 10}, target: 10, want: 4},
		{name: "not found", arr: []int{2, 4, 6, 8, 10}, target: 5, want: -1},
		{name: "empty array", arr: []int{}, target: 1, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearchRecursive(tt.arr, tt.target, 0, len(tt.arr)-1)
			if got != tt.want {
				t.Fatalf("BinarySearchRecursive() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFindInsertPosition(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{name: "insert in middle", arr: []int{1, 3, 5, 7}, target: 4, want: 2},
		{name: "insert at start", arr: []int{1, 3, 5, 7}, target: 0, want: 0},
		{name: "insert at end", arr: []int{1, 3, 5, 7}, target: 9, want: 4},
		{name: "existing value", arr: []int{1, 3, 5, 7}, target: 5, want: 2},
		{name: "empty array", arr: []int{}, target: 10, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindInsertPosition(tt.arr, tt.target)
			if got != tt.want {
				t.Fatalf("FindInsertPosition() = %d, want %d", got, tt.want)
			}
		})
	}
}
