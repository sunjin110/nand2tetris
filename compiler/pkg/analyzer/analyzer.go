package analyzer

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/tokenizer"
	"compiler/pkg/xmlwriter"
	"fmt"
)

// Analyzer Xmlに変換する
func Analyzer(outputFileName string, pathList []string) {
	jsonutil.Print(pathList)

	// xml writerを作成する
	xmlWriter, err := xmlwriter.New(outputFileName)
	chk.SE(err)
	defer func() {
		xmlWriter.Close()
	}()

	// コードをparseする
	for _, path := range pathList {

		fmt.Println("path is ", path)

		t, err := tokenizer.New(path)
		chk.SE(err)

		// tokenを一つずつ取得する
		for t.NextToken() {

			token := t.Token
			fmt.Println("token is ", token)

			tokenType := tokenizer.GetTokenType(token)
			switch tokenType {
			case tokenizer.TokenTypeKeyWord:
				// keyword := tokenizer.GetKeyWord(token)
				xmlWriter.WriteToken(tokenType, token)
			case tokenizer.TokenTypeSymbol:
				symbol := t.GetSymbol()
				xmlWriter.WriteToken(tokenType, string(symbol))
			case tokenizer.TokenTypeIdentifier:
				identifier := t.GetIdentifier()
				xmlWriter.WriteToken(tokenType, identifier)
			case tokenizer.TokenTypeIntConst:
				intVal := t.GetIntVal()
				xmlWriter.WriteToken(tokenType, fmt.Sprintf("%d", intVal))
			case tokenizer.TokenTypeStringConst:
				strVal := t.GetStringVal()
				xmlWriter.WriteToken(tokenType, strVal)
			default:
				panic("定義されていないtokenType")
			}

		}

	}

}
