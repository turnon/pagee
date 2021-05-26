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
	LimitPages int
	pagesCount int
}

type Element struct {
	*colly.HTMLElement
}

func (w *Walk) Start(fn func(e *Element)) error {
	if w.Uri == "" || w.Next == "" || w.Item == "" {
		return errors.New("uri/href/item not defined")
	}

	c := colly.NewCollector()

	c.OnHTML(w.Item, func(e *colly.HTMLElement) {
		if w.reachLimit() {
			return
		}
		fn(&Element{e})
		w.itemsCount += 1
	})

	c.OnHTML(w.Next, func(e *colly.HTMLElement) {
		if w.reachLimit() {
			return
		}
		w.pagesCount += 1
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(w.Uri)

	return nil
}

func (w *Walk) reachLimit() bool {
	return (w.LimitItems != 0 && w.itemsCount >= w.LimitItems) || (w.LimitPages != 0 && w.pagesCount >= w.LimitPages)
}
