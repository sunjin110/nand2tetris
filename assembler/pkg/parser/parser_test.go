package parser_test

import (
	"assembler/pkg/parser"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1  -timeout 30s -run ^Test$ assembler/pkg/parser

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("parser", func() {

		g.It("GetSymbol", func() {

			result := parser.GetSymbol("@symbol", parser.ACommand)
			g.Assert(result).Eql("symbol")

			result2 := parser.GetSymbol("(symbol2)", parser.LCommand)
			g.Assert(result2).Eql("symbol2")

		})

	})

}
