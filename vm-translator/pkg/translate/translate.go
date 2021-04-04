package translate

import (
	"fmt"
	"vm-translator/pkg/codewriter"
	"vm-translator/pkg/common/chk"
	"vm-translator/pkg/common/jsonutil"
	"vm-translator/pkg/parser"
)

// Translate 実際に変換する
func Translate(outputFileName string, pathList []string) {

	jsonutil.Print(pathList)

	// code writerを作成
	codeWriter, err := codewriter.New(outputFileName)
	chk.SE(err)
	defer func() {
		codeWriter.Close()
	}()

	// コードをparseしていく
	for _, path := range pathList {
		parser, err := parser.New(path)
		chk.SE(err)

		// 解析をしていく
		for parser.Next() {
			fmt.Println("command is ", parser.Command)
		}

	}

}
