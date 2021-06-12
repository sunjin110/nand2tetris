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

			g.It("run", func() {

				result := math.Multiply(11, 5)
				log.Println("result is ", result)
				g.Assert(result).Eql(int16(55))
			})

		})

	})

}
