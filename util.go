package bowling

func sum(x ...int) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func min(x ...int) (min int, err error) {
	if x == nil || len(x) <= 0 {
		err = &emptyArrayError{"Call to min expects non-empty int variargs array"}
		return
	}

	min = x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}

	return
}
