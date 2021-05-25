package pagee

import (
	"testing"

	"github.com/gocolly/colly/v2"
)

func TestNextStart(t *testing.T) {
	n := Next{
		Uri:  "https://gocn.vip/topics/excellent",
		Href: ".pagination .next a",
		Item: ".topics .topic .title a",
	}

	elements := []string{}
	err := n.Start(func(e *colly.HTMLElement) {
		elements = append(elements, e.Attr("href"))
	})

	t.Log(len(elements))
	t.Log(elements)

	if err != nil {
		t.Error(err)
	}
}
