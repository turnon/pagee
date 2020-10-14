package pagee

import "testing"

func TestUrl2docWithRighturl(t *testing.T) {
	url := "http://ifeve.com/"
	doc, err := url2doc(url)
	if err != nil {
		t.Error(url + " should be accessible")
	}
	if doc == nil {
		t.Error(url + " should be parse-able")
	}
}

func TestUrl2docWithWrongurl(t *testing.T) {
	url := "http://ifeve2.com/"
	doc, err := url2doc(url)
	if err == nil {
		t.Error(url + " should not be accessible")
	}
	if doc != nil {
		t.Error(url + " should not be parse-able")
	}
}

func TestFillURL(t *testing.T) {
	urlTemp := "http://ifeve.com/page/{{.}}/"
	str, err := fillURL(urlTemp, 2)

	if err != nil {
		t.Error(urlTemp+" should be filled", err)
	}

	expected := "http://ifeve.com/page/2/"
	if *str != expected {
		t.Error(urlTemp + " should turn into " + expected + ", but now it is " + *str)
	}
}
