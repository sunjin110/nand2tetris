package analyzer

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/tokenizer"
	"fmt"
)

// Analyzer Xmlに変換する
func Analyzer(outputFileName string, pathList []string) {
	jsonutil.Print(pathList)

	// コードをparseする
	for _, path := range pathList {

		fmt.Println("path is ", path)

		t, err := tokenizer.New(path)
		chk.SE(err)

		// 1行ずつ読んでいく
		for t.NextLine() {

			line := t.Line

			tokenList := tokenizer.CreateTokenList(line)
			for _, token := range tokenList {
				fmt.Println("token is ", token)
			}

		}

	}

}
