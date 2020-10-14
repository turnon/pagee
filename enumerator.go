package pagee

func newEnum(from, to int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- from
			if from >= to {
				close(ch)
				return
			}
			from++
		}
	}()
	return ch
}
