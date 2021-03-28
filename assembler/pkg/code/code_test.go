package code_test

import (
	"assembler/pkg/code"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s -run ^Test$ assembler/pkg/code

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Code", func() {

		g.It("Dest", func() {

			result := code.ConvDest("M")
			log.Println("result is ", result)

		})

	})

}
