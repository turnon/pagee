package pagee

import (
	"math"
	"testing"
)

// func TestFillURL(t *testing.T) {
// 	urlTemp := "http://ifeve.com/page/{{.}}/"
// 	str, err := fillURL(urlTemp, 2)

// 	if err != nil {
// 		t.Error(urlTemp+" should be filled", err)
// 	}

// 	expected := "http://ifeve.com/page/2/"
// 	if *str != expected {
// 		t.Error(urlTemp + " should turn into " + expected + ", but now it is " + *str)
// 	}
// }

func TestPageNumEnum(t *testing.T) {
	f := fetcher{From: 1, Step: 2, To: 6}
	enum := f.pageNumEnum()

	arr := []float64{}
	for n := range enum {
		arr = append(arr, n.(float64))
	}

	if len(arr) != 3 || arr[0] != 1 || arr[1] != 3 || arr[2] != 5 {
		t.Error(arr)
	}
}

func TestInfinityPageNumEnum(t *testing.T) {
	f := fetcher{From: 10, To: math.Inf(-1)}
	enum := f.pageNumEnum()

	arr := []float64{}
	for n := range enum {
		arr = append(arr, n.(float64))
		if len(arr) >= 4 {
			break
		}
	}

	if len(arr) != 4 || arr[0] != 10 || arr[1] != 9 || arr[2] != 8 || arr[3] != 7 {
		t.Error(arr)
	}
}

func TestURLEnum(t *testing.T) {
	f := fetcher{URL: "http://ifeve.com/page/{{.}}/", From: 1, To: math.Inf(1)}
	enum, _ := f.urlEnum()

	arr := []string{}
	for url := range enum {
		arr = append(arr, url)
		if len(arr) >= 3 {
			break
		}
	}

	if arr[0] != "http://ifeve.com/page/1/" || arr[1] != "http://ifeve.com/page/2/" || arr[2] != "http://ifeve.com/page/3/" {
		t.Error(arr)
	}
}

func TestInterval(t *testing.T) {
	f := fetcher{URL: "http://ifeve.com/page/{{.}}/", From: 1, To: 3, Interval: 1}
	enum, _ := f.urlEnum()

	testDuration(t, 3, 4, func() {
		for range enum {
		}
	})
}
