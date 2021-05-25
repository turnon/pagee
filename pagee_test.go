package pagee

import (
	"testing"

	"github.com/gocolly/colly/v2"
)

const (
	uri  = "https://gocn.vip/topics/excellent"
	next = ".pagination .next a"
	item = ".topics .topic .title a"
)

func TestStart(t *testing.T) {
	w := Walk{
		Uri:  uri,
		Next: next,
		Item: item,
	}

	elements := []string{}
	err := w.Start(func(e *colly.HTMLElement) {
		elements = append(elements, e.Attr("href"))
	})

	t.Log(len(elements))
	t.Log(elements[:60])

	if err != nil {
		t.Error(err)
	}
}

func TestLimitItems(t *testing.T) {
	w := Walk{
		Uri:        uri,
		Next:       next,
		Item:       item,
		LimitItems: 61,
	}

	elements := []string{}
	err := w.Start(func(e *colly.HTMLElement) {
		elements = append(elements, e.Attr("href"))
	})

	count := len(elements)
	if count != 61 {
		t.Errorf("limit is not working -> %d", count)
	}

	t.Log(elements[:count])

	if err != nil {
		t.Error(err)
	}
}
