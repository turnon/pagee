package pagee

import (
	"encoding/json"
	"time"
)

type m map[string]int

type cfg struct {
	enumerator
}

func parseConfig(bytes []byte) cfg {
	var c cfg
	json.Unmarshal(bytes, &c.enumerator)
	return c
}

func (c cfg) enum() <-chan interface{} {
	r := intChanToInterfaceChan(c.enumerator.toRange())
	interval, intervalDefinded := c.enumerator["interval"]

	if !intervalDefinded {
		return r
	}

	return intervalRead(time.Duration(interval), r)
}

func intChanToInterfaceChan(ch <-chan int) <-chan interface{} {
	convertor := make(chan interface{})
	go func() {
		for i := range ch {
			convertor <- i
		}
		close(convertor)
	}()

	return convertor
}
