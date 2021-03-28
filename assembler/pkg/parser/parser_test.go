package parser_test

import (
	"assembler/pkg/parser"
	"testing"

	"github.com/franela/goblin"
)

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("parser", func() {

		g.It("GetSymbol", func() {

			result := parser.GetSymbol("@symbol", parser.ACommand)
			g.Assert(result).Eql("symbol")

		})

	})

}
