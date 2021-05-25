package pagee

import (
	"errors"

	"github.com/gocolly/colly/v2"
)

type Walk struct {
	Uri  string
	Next string
	Item string
}

func (w *Walk) Start(fn func(e *colly.HTMLElement)) error {
	if w.Uri == "" || w.Next == "" || w.Item == "" {
		return errors.New("uri/href/item not defined")
	}

	c := colly.NewCollector()

	c.OnHTML(w.Item, fn)

	c.OnHTML(w.Next, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(w.Uri)

	return nil
}
