package pagee

func hashToRanger(hash map[string]interface{}) <-chan int {
	from := hash["from"]

	step, stepExists := hash["step"]
	if !stepExists {
		step = 1
	}

	to, toExists := hash["to"]
	if toExists {
		return newFinRange(from.(int), to.(int), step.(int))
	}

	return newInfRange(from.(int), step.(int))
}

func newFinRange(from, to int, step ...int) <-chan int {
	st := takeStep(step)

	ex := func(fr int) bool {
		return fr > to
	}

	return newRange(from, st, ex)
}

func newInfRange(from int, step ...int) <-chan int {
	st := takeStep(step)

	ex := func(_ int) bool {
		return false
	}

	return newRange(from, st, ex)
}

func takeStep(args []int) int {
	if len(args) >= 1 {
		return args[0]
	}
	return 1
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
