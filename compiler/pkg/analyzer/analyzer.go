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

// // Analyzer Xmlに変換する
// func Analyzer(outputFileName string, pathList []string) {
// 	jsonutil.Print(pathList)

// 	// コードをparseする
// 	for _, path := range pathList {

// 		fmt.Println("path is ", path)

// 		t, err := tokenizer.New(path)
// 		chk.SE(err)

// 		// 1行ずつ読んでいく
// 		for t.NextLine() {

// 			line := t.Line

// 			tokenList := tokenizer.CreateTokenList(line)
// 			log.Println("token list is ", jsonutil.Marshal(tokenList))
// 			for _, token := range tokenList {
// 				fmt.Println("token is ", token)
// 				tokenType := tokenizer.GetTokenType(token)
// 				log.Println("tokenType is ", tokenType)
// 			}

// 		}

// 	}

// }
