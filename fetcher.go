package pagee

import (
	"encoding/json"
	"strings"
	"text/template"
	"time"
)

type fetcher struct {
	URL      string  `json:"url"`
	Selector string  `json:"selector"`
	From     float64 `json:"from"`
	Step     float64 `json:"step"`
	To       float64 `json:"to"`
	Interval int     `json:"interval"`
}

var temp = template.New("fetcher")

func parse(js []byte) fetcher {
	var f fetcher
	json.Unmarshal(js, &f)
	return f
}

func (f fetcher) waiting() <-chan time.Time {
	if f.Interval > 0 {
		return time.NewTicker(time.Duration(f.Interval) * time.Second).C
	}

	ch := make(chan time.Time)
	close(ch)
	return ch
}

func (f fetcher) urlEnum() (<-chan string, error) {
	wait := f.waiting()
	pageNumbers := f.pageNumEnum()
	urls := make(chan string)

	t, err := temp.Parse(f.URL)
	if err != nil {
		return nil, err
	}

	go func() {
		for n := range pageNumbers {
			url, err := fillURL(t, n)
			if err != nil {
				urls <- err.Error()
				break
			}
			urls <- *url
			<-wait
		}
		close(urls)
	}()

	return urls, nil
}

func (f fetcher) pageNumEnum() <-chan interface{} {
	from := f.From
	step := f.Step
	if step == 0 {
		if f.From > f.To {
			step = -1
		} else {
			step = 1
		}
	}

	exceeded := func(fr float64) bool {
		return fr > f.To
	}
	if step < 0 {
		exceeded = func(fr float64) bool {
			return fr < f.To
		}
	}

	ch := make(chan interface{})
	go func() {
		for {
			if exceeded(from) {
				close(ch)
				return
			}
			ch <- from
			from += step
		}
	}()

	return ch
}

func fillURL(t *template.Template, n interface{}) (*string, error) {
	sb := &strings.Builder{}
	if err := t.Execute(sb, n); err != nil {
		return nil, err
	}

	str := sb.String()
	return &str, nil
}
