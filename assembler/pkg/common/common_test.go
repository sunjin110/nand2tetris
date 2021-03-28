package common_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s -run ^Test$ assembler/pkg/common

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("common", func() {
		g.It("Uint2bytes", func() {

			// result := common.Uint2bytes(2, 16)
			// log.Println("result is ", result)

		})

		g.It("aaa", func() {

			// fmt.Printf("%b", "")
			// parseUint("0101010")
			// parseUint("0111111")
			// parseUint("0111010")
			// parseUint("0001100")
			// parseUint("0110000")
			// parseUint("1110000")
			// parseUint("0001101")
			// parseUint("0110001")
			// parseUint("1110001")
			// parseUint("0001111")
			// parseUint("0110011")
			// parseUint("1110011")
			// parseUint("0011111")
			// parseUint("0110111")
			// parseUint("1110111")
			// parseUint("0001110")
			// parseUint("0110010")
			// parseUint("1110010")
			// parseUint("0000010")
			// parseUint("1000010")
			// parseUint("0010011")
			// parseUint("1010011")
			// parseUint("0000111")
			// parseUint("1000111")
			// parseUint("0000000")
			// parseUint("1000000")
			// parseUint("0010101")
			// parseUint("1010101")
			parseUint("0000000000001000")

		})

		g.It("StrToUint", func() {

			// result := common.StrToUint("111")
			// log.Println("result is ", result)

		})
	})

}

func parseUint(s2 string) {
	v, _ := strconv.ParseUint(s2, 2, 0)
	fmt.Println(v)
	// return string(v)
}
