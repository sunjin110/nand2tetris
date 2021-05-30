package analyzer

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/tokenizer"
	"compiler/pkg/xmlwriter"
	"log"
	"strings"
)

// Analyzer Xmlに変換する
func Analyzer(outputFileName string, pathList []string) {
	jsonutil.Print(pathList)

	for _, path := range pathList {

		t, err := tokenizer.New(path)
		chk.SE(err)

		compilationEngine := compilation_engine.New(t)

		compilationEngine.Start()

		c := compilationEngine.Class

		// writeXml
		writeXml(path, c)

	}

}

// writeXml xmlとして出力する
func writeXml(path string, c *compilation_engine.Class) {

	// xmlに変換する
	outputFileName := strings.Split(path, ".")[0] + ".xml"
	log.Println("output file name is ", outputFileName)
	// xmlwriter.WriteFromClass(outputFileName+"sunjin.xml", c)

	xw := xmlwriter.New(outputFileName, c)
	xw.WriteParser()
}

// // Analyzer Xmlに変換する
// func Analyzer(outputFileName string, pathList []string) {
// 	jsonutil.Print(pathList)

// 	// Tokenizerテスト用のxxxT.xmlのファイルを作成するxmlWriterを作成
// 	// outputFileNameSplit := strings.Split(outputFileName, ".")
// 	// outputFileTName := fmt.Sprintf("%sT.%s", outputFileNameSplit[0], outputFileNameSplit[1])
// 	// xmlWriterT, err := xmlwriter.New(outputFileTName)
// 	// chk.SE(err)
// 	// defer func() {
// 	// 	xmlWriterT.Close()
// 	// }()

// 	// コードをparseする
// 	for _, path := range pathList {

// 		fmt.Println("path is ", path)

// 		// Tokenizerテスト用のxxxT.xmlのファイルを作成する
// 		xmlWriterT, err := xmlwriter.New(strings.Split(path, ".")[0] + "T.xml")
// 		chk.SE(err)

// 		t, err := tokenizer.New(path)
// 		chk.SE(err)

// 		// tokenを一つずつ取得する
// 		for t.NextToken() {

// 			token := t.Token

// 			tokenType := tokenizer.GetTokenType(token)
// 			switch tokenType {
// 			case tokenizer.TokenTypeKeyWord:
// 				xmlWriterT.WriteToken(tokenType, token)
// 			case tokenizer.TokenTypeSymbol:
// 				symbol := t.GetSymbol()
// 				xmlWriterT.WriteToken(tokenType, string(symbol))
// 			case tokenizer.TokenTypeIdentifier:
// 				identifier := t.GetIdentifier()
// 				xmlWriterT.WriteToken(tokenType, identifier)
// 			case tokenizer.TokenTypeIntConst:
// 				intVal := t.GetIntVal()
// 				xmlWriterT.WriteToken(tokenType, fmt.Sprintf("%d", intVal))
// 			case tokenizer.TokenTypeStringConst:
// 				strVal := t.GetStringVal()
// 				xmlWriterT.WriteToken(tokenType, strVal)
// 			default:
// 				panic("定義されていないtokenType")
// 			}
// 		}

// 		// Tokenizerテスト用のxxxT.xmlのファイルを閉じる
// 		xmlWriterT.Close()

// 	}

// }
