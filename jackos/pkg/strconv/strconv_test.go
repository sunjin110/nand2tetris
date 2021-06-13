package strconv_test

import (
	"testing"

	"jackos/pkg/strconv"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s jackos/pkg/strconv

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("stronv", func() {

		g.Describe("IntToString", func() {

			g.It("10", func() {
				str := strconv.IntToString(10)
				g.Assert(str).Eql("10")
			})

			g.It("3141592", func() {

				str := strconv.IntToString(3141592)
				g.Assert(str).Eql("3141592")

			})
		})

		g.Describe("StringToInt", func() {

			g.It("1", func() {
				n := strconv.StringToInt("1")
				g.Assert(n).Eql(1)
			})
			g.It("3141592", func() {
				n := strconv.StringToInt("3141592")
				g.Assert(n).Eql(3141592)
			})

		})

	})

}
