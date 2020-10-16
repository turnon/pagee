package pagee

import "testing"

func TestParseConfig(t *testing.T) {
	js := `{"from": 1, "to": 3, "interval": 1}`
	c := parseConfig([]byte(js))
	if c.enumerator["from"] != 1 || c.enumerator["to"] != 3 {
		t.Error(c, " is not ", js)
	}

	enum := c.enum()
	testDuration(t, 3, 4, func() {
		for range enum {
		}
	})
}
