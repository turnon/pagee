package pagee

import "testing"

func TestNewEnum(t *testing.T) {
	enum := newEnum(1, 3)
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
	}

	errMsg := "newEnum is wrong"
	if arr[0] != 1 || arr[1] != 2 || arr[2] != 3 || len(arr) != 3 {
		t.Error(errMsg, arr)
	}
}
