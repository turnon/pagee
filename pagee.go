package pagee

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func url2doc(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

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
