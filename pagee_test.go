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
