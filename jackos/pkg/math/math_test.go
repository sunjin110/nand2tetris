package math_test

import (
	"jackos/pkg/math"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s jackos/pkg/math

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Math Test", func() {

		g.Describe("Multiply", func() {
			g.It("11 x 5", func() {
				result := math.Multiply(11, 5)
				log.Println("result is ", result)
				g.Assert(result).Eql(int16(55))
			})

			g.It("11 x -5", func() {
				g.Assert(math.Multiply(11, -5)).Eql(int16(-55))
			})
		})

		g.Describe("Devide", func() {
			g.It("10 / 2", func() {
				g.Assert(math.Devide(10, 2)).Eql(int16(5))
			})
			g.It("11 / 2", func() {
				g.Assert(math.Devide(11, 2)).Eql(int16(5))
			})
		})

		g.Describe("Pow", func() {

			g.It("10の2乗", func() {
				g.Assert(math.Pow(10, 2)).Eql(int16(100))
			})

		})

		g.Describe("Sqrt", func() {

			g.It("√4 == 2", func() {
				g.Assert(math.Sqrt(4)).Eql(int16(2))
			})

			g.It("√9 == 3", func() {
				g.Assert(math.Sqrt(9)).Eql(int16(3))
			})

			g.It("√10 == 3", func() {
				g.Assert(math.Sqrt(10)).Eql(int16(3))
			})

		})

	})

}
