package pagee

import "testing"

func TestHashToRangerUsingFinRangeWithoutStep(t *testing.T) {
	hash := map[string]interface{}{
		"from": 1,
		"to":   3,
	}
	enum := hashToRanger(hash)
	testFinWithoutStep(enum, t)
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

func TestHashToRangerUsingFinRangeWithStep(t *testing.T) {
	hash := map[string]interface{}{
		"from": 1,
		"to":   6,
		"step": 2,
	}
	enum := hashToRanger(hash)
	testFinWithStep(enum, t)
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

func TestHashToRangerUsingInfRangeWithoutStep(t *testing.T) {
	hash := map[string]interface{}{
		"from": 1,
	}
	enum := hashToRanger(hash)
	testInfWithoutStep(enum, t)
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

func TestHashToRangerUsingInfRangeWithStep(t *testing.T) {
	hash := map[string]interface{}{
		"from": 1,
		"step": 2,
	}
	enum := hashToRanger(hash)
	testInfWithStep(enum, t)
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
