package pagee

func newFinRange(from, to int, step ...int) <-chan int {
	st := 1
	if len(step) >= 1 {
		st = step[0]
	}

	ex := func(fr int) bool {
		return fr > to
	}

	return newRange(from, st, ex)
}

func newRange(from int, step int, exceeded func(int) bool) <-chan int {
	ch := make(chan int)
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
