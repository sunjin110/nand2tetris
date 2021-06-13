package memory_test

import (
	"jackos/pkg/memory/simplememory"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s jackos/pkg/memory

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("memory", func() {

		g.Describe("simple memory", func() {

			g.It("flow", func() {

				m := simplememory.New()
				m.Alloc(10)

			})

		})

	})

}
