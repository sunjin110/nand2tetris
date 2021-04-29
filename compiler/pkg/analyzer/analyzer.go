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

		// tokenを一つずつ取得する
		for t.NextToken() {

			token := t.Token
			fmt.Println("token is ", token)

		}

	}

}
