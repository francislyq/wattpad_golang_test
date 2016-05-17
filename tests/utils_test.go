package tests

import (
	"testing"
	"wattpad/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMath(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 2
		y := 3

		Convey("When the integer is incremented", func() {
			score := utils.CalOffensiveScore(x, y)
			So(score, ShouldEqual, 7)
		})
	})
}
