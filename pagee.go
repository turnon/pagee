package pagee

import (
	"errors"

	"github.com/gocolly/colly/v2"
)

type Next struct {
	Uri  string
	Href string
	Item string
}

func (n *Next) Start(fn func(e *colly.HTMLElement)) error {
	if n.Uri == "" || n.Href == "" || n.Item == "" {
		return errors.New("uri/href/item not defined")
	}

	c := colly.NewCollector()

	c.OnHTML(n.Item, fn)

	c.OnHTML(n.Href, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(n.Uri)

	return nil
}
