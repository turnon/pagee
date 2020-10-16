package pagee

import (
	"testing"
)

func TestJSONFinRangeWithoutStep(t *testing.T) {
	it := newIter([]byte(`{"from": 1, "to": 3}`))
	testFinWithoutStep(it.toRange(), t)
}

func TestNewFinRangeWithoutStep(t *testing.T) {
	enum := newFinRange(1, 3)
	testFinWithoutStep(enum, t)
}

func testFinWithoutStep(enum <-chan int, t *testing.T) {
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Error(arr)
	}
}

func TestJSONFinRangeWithStep(t *testing.T) {
	it := newIter([]byte(`{"from": 1, "step": 2, "to": 6}`))
	testFinWithStep(it.toRange(), t)
}

func TestNewFinRangeWithStep(t *testing.T) {
	enum := newFinRange(1, 6, 2)
	testFinWithStep(enum, t)
}

func testFinWithStep(enum <-chan int, t *testing.T) {
	arr := []int{}
	for n := range enum {
		arr = append(arr, n)
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 3 || arr[2] != 5 {
		t.Error(arr)
	}
}

func TestJSONInfRangeWithoutStep(t *testing.T) {
	it := newIter([]byte(`{"from": 1}`))
	testInfWithoutStep(it.toRange(), t)
}

func TestNewInfRangeWithoutStep(t *testing.T) {
	enum := newInfRange(1)
	testInfWithoutStep(enum, t)
}

func testInfWithoutStep(enum <-chan int, t *testing.T) {
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

func TestJSONInfRangeInfRangeWithStep(t *testing.T) {
	it := newIter([]byte(`{"from": 1, "step": 2}`))
	testInfWithStep(it.toRange(), t)
}

func TestNewInfRangeWithStep(t *testing.T) {
	enum := newInfRange(1, 2)
	testInfWithStep(enum, t)
}

func testInfWithStep(enum <-chan int, t *testing.T) {
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
