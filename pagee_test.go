package pagee

import "testing"

func TestUri2docWithRightUri(t *testing.T) {
	uri := "http://ifeve.com/"
	doc, err := uri2doc(uri)
	if err != nil {
		t.Error(uri + " should be accessible")
	}
	if doc == nil {
		t.Error(uri + " should be parse-able")
	}
}

func TestUri2docWithWrongUri(t *testing.T) {
	uri := "http://ifeve2.com/"
	doc, err := uri2doc(uri)
	if err == nil {
		t.Error(uri + " should not be accessible")
	}
	if doc != nil {
		t.Error(uri + " should not be parse-able")
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
