package pagee

import (
	"math"
	"testing"
)

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

func TestDocEnum(t *testing.T) {
	f := fetcher{URL: "http://ifeve.com/page/{{.}}/", From: 1, To: 3}
	enum, err := f.docEnum()
	if err != nil {
		t.Error(err)
		return
	}

	docCount := 0
	for range enum {
		docCount++
	}

	if docCount != 3 {
		t.Error(docCount)
	}
}

func TestItemEnum(t *testing.T) {
	f := fetcher{URL: "http://ifeve.com/page/{{.}}/", From: 1, To: 3, Selector: ".post_wrap"}
	enum, err := f.ItemEnum()
	if err != nil {
		t.Error(err)
		return
	}

	itemCount := 0
	for range enum {
		itemCount++
	}

	if itemCount != 45 {
		t.Error(itemCount)
	}
}

func TestUntilEndItemEnum(t *testing.T) {
	f := fetcher{URL: "http://ifeve.com/page/{{.}}/", From: 111, To: math.Inf(1), Selector: ".post_wrap .title a"}
	enum, err := f.ItemEnum()
	if err != nil {
		t.Error(err)
		return
	}

	itemCount := 0
	for range enum {
		itemCount++
	}

	// today last page is 115
	if itemCount < 60 {
		t.Error(itemCount)
	}
}
