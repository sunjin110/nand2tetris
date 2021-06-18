package output_test

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s jackos/pkg/output

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("output", func() {
		g.It("test", func() {
			// fmt.Printf("%b\n", int16(255))
			// fmt.Println(-1 & 255)
			// fmt.Printf("%b\n")
			fmt.Printf("%b\n", 0&1)
			fmt.Printf("%b\n", 1&1)
			fmt.Printf("%b\n", 2&1)
			fmt.Printf("%b\n", 3&1)
			fmt.Printf("%d", 102&1)

		})
	})

}
