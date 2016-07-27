package bowling

import "log"

// Game of bowling
type Game struct {
	frames []*frame
}

// Roll number of pins knocked down
func (game *Game) Roll(pins int) {
	if game.frames == nil || len(game.frames) <= 0 {
		game.addFrame(pins)
		return
	}

	if game.currentFrame().canRoll(pins) {
		game.currentFrame().roll(pins)
	} else {
		game.addFrame(pins)
	}
}

// RollMany pins
func (game *Game) RollMany(pins ...int) {
	for i := 0; i < len(pins); i++ {
		game.Roll(pins[i])
	}
}

// Score a bowling game
func (game *Game) Score() int {
	score := 0
	for i := 0; i < len(game.frames); i++ {
		f := game.frames[i]
		score += f.score()
		log.Printf("Frame %d, Score %d\n", f.number, score)
	}
	return score
}

func (game *Game) currentFrame() *frame {
	if len(game.frames) <= 0 {
		return nil
	}
	return game.frames[len(game.frames)-1]
}

func (game *Game) addFrame(pins int) {

	if pins > 10 {
		return
	}

	prev := game.currentFrame()

	number := 1
	if prev != nil {
		number += prev.number
	}

	f := frame{
		prev:   prev,
		number: number,
		pins:   []int{pins},
	}

	if prev != nil {
		prev.next = &f
	}

	game.frames = append(game.frames, &f)
}
