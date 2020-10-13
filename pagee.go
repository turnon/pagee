package pagee

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

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
