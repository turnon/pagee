package pagee

import (
	"strings"
	"text/template"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var temp = template.New("a")

func uri2doc(uri string) (*goquery.Document, error) {
	resp, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func fillURL(uri string, n interface{}) (*string, error) {
	t, err := temp.Parse(uri)
	if err != nil {
		return nil, err
	}

	sb := &strings.Builder{}
	if err = t.Execute(sb, n); err != nil {
		return nil, err
	}

	str := sb.String()
	return &str, nil
}
