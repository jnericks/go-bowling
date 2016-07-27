package bowling

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBowlingGame(t *testing.T) {

	Convey("Given empty game", t, func() {
		game := NewGame()
		Convey("It should have no frames", func() {
			So(len(game.frames), ShouldEqual, 0)
		})
		Convey("It should score 0", func() {
			So(game.Score(), ShouldEqual, 0)
		})
	})

	Convey("Given complete game with no spares", t, func() {
		game := NewGame()
		game.RollMany(9, 0, 8, 0, 7, 0, 6, 0, 5, 0, 4, 0, 3, 0, 2, 0, 1, 0, 9, 0)
		Convey("It should have 10 frames", func() {
			So(len(game.frames), ShouldEqual, 10)
		})
		Convey("It should score as a simple sum", func() {
			So(game.Score(), ShouldEqual, 54)
		})
	})

	Convey("Given a perfect game", t, func() {
		game := NewGame()
		game.RollMany(10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10)
		Convey("It should have 10 frames", func() {
			So(len(game.frames), ShouldEqual, 10)
		})
		Convey("It should have a score of 300", func() {
			So(game.Score(), ShouldEqual, 300)
		})
	})

	Convey("Given a strike rolled after a spare", t, func() {
		game := NewGame()
		game.RollMany(9, 1, 10)
		Convey("It should have 2 frames", func() {
			So(len(game.frames), ShouldEqual, 2)
		})
		Convey("It should have a score of 30", func() {
			So(game.Score(), ShouldEqual, 30)
		})
	})

	Convey("Given a spare rolled after a strike", t, func() {
		game := NewGame()
		game.RollMany(10, 9, 1)
		Convey("It should have 2 frames", func() {
			So(len(game.frames), ShouldEqual, 2)
		})
		Convey("It should have a score of 30", func() {
			So(game.Score(), ShouldEqual, 30)
		})
	})

	Convey("Given complete game with a mix of open, spare and strike frames", t, func() {
		game := NewGame()
		game.RollMany(1, 4, 4, 5, 6, 4, 5, 5, 10, 0, 1, 7, 3, 6, 4, 10, 2, 8, 6)
		Convey("It should have 10 frames", func() {
			So(len(game.frames), ShouldEqual, 10)
		})
		Convey("It should have correct score", func() {
			So(game.Score(), ShouldEqual, 133)
		})
	})
}
