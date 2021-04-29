package tokenizer_test

import (
	"compiler/pkg/tokenizer"
	"fmt"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 ./pkg/tokenizer

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("tokenizer", func() {

		g.It("GetKeyWord", func() {

			result := tokenizer.GetKeyWord("class")
			log.Println("result is ", result)

			g.Assert(result == tokenizer.KeyWordClass).IsTrue()

		})

		g.It("trim ", func() {

			hoge := "abcdefg"

			hogeTrim := hoge[1 : len(hoge)-1]

			fmt.Println("hoge trim is ", hogeTrim)

			g.Assert(hogeTrim).Eql("bcdef")
		})

	})

}
