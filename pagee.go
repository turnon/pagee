package pagee

import (
	"errors"

	"github.com/gocolly/colly/v2"
)

type Walk struct {
	Uri  string
	Next string
	Item string

	LimitItems int
	itemsCount int
}

func (w *Walk) Start(fn func(e *colly.HTMLElement)) error {
	if w.Uri == "" || w.Next == "" || w.Item == "" {
		return errors.New("uri/href/item not defined")
	}

	c := colly.NewCollector()

	c.OnHTML(w.Item, func(e *colly.HTMLElement) {
		if w.reachLimit() {
			return
		}
		fn(e)
		w.itemsCount += 1
	})

	c.OnHTML(w.Next, func(e *colly.HTMLElement) {
		if w.reachLimit() {
			return
		}
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(w.Uri)

	return nil
}

func (w *Walk) reachLimit() bool {
	if w.LimitItems == 0 {
		return false
	}
	return w.itemsCount >= w.LimitItems
}
