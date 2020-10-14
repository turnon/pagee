package pagee

import "testing"

func TestNewFinRangeWithoutStep(t *testing.T) {
	enum := newFinRange(1, 3)
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Error(arr)
	}
}

func TestNewFinRangeWithStep(t *testing.T) {
	enum := newFinRange(1, 6, 2)
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 3 || arr[2] != 5 {
		t.Error(arr)
	}
}

func TestNewInfRangeWithoutStep(t *testing.T) {
	enum := newInfRange(1)
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
		if len(arr) >= 3 {
			break
		}
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Error(arr)
	}
}

func TestNewInfRangeWithStep(t *testing.T) {
	enum := newInfRange(1, 2)
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
		if len(arr) >= 3 {
			break
		}
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 3 || arr[2] != 5 {
		t.Error(arr)
	}
}
