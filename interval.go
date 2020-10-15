package pagee

import (
	"time"
)

func intervalRead(sec time.Duration, ch <-chan interface{}) <-chan interface{} {
	wrapper := make(chan interface{})
	go func() {
		for any := range ch {
			wrapper <- any
			<-time.After(sec * time.Second)
		}
		close(wrapper)
	}()
	return wrapper
}

func intervalReadInt(sec time.Duration, ch <-chan int) <-chan interface{} {
	convertor := make(chan interface{})
	go func() {
		for i := range ch {
			convertor <- i
		}
		close(convertor)
	}()
	return intervalRead(sec, convertor)
}
