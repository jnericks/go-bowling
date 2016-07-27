package bowling

type frame struct {
	prev   *frame
	number int
	pins   []int
	next   *frame
}

func (f *frame) numRolls() int {
	return len(f.pins)
}

func (f *frame) canRoll(pins int) bool {

	// Defensive checks
	if pins > 10 || f.number > 10 {
		return false
	}

	// 10th frame
	if f.number == 10 {

		if f.numRolls() >= 3 {
			return false
		}

		if f.numRolls() == 2 && sum(f.pins...)%10 == 0 {
			return true
		}

		if f.numRolls() == 0 {
			return true
		}

		return f.pins[0] == 10 || pins+f.pins[0] <= 10
	}

	if f.number < 10 && f.numRolls() >= 2 {
		return false
	}

	if f.numRolls() == 0 {
		return true
	}

	if f.isStrike() {
		return false
	}

	return pins+f.pins[0] <= 10
}

func (f *frame) roll(pins int) {
	f.pins = append(f.pins, pins)
}

func (f *frame) isStrike() bool {
	return f.numRolls() == 1 && f.pins[0] == 10
}

func (f *frame) isSpare() bool {
	return f.numRolls() == 2 && f.pins[0]+f.pins[1] == 10
}

func (f *frame) score() int {
	if f.numRolls() == 0 {
		return 0
	}

	include := 0
	if f.number < 10 {
		if f.isStrike() {
			include = 2
		} else if f.isSpare() {
			include = 1
		}
	}

	score := sum(f.pins...)
	if f.next != nil && include > 0 {
		pins := f.next.pins
		if f.next.next != nil {
			pins = append(pins, f.next.next.pins...)
		}
		limit, _ := min(len(pins), include)
		for _, p := range pins[:limit] {
			score += p
		}
	}

	return score
}
