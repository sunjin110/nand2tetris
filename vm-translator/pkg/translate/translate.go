package translate

import (
	"fmt"
	"vm-translator/pkg/codewriter"
	"vm-translator/pkg/common/chk"
	"vm-translator/pkg/common/jsonutil"
	"vm-translator/pkg/model"
	"vm-translator/pkg/parser"
)

// Translate 実際に変換する
func Translate(outputFileName string, pathList []string) {

	jsonutil.Print(pathList)

	// code writerを作成
	codeWriter, err := codewriter.New(outputFileName)
	chk.SE(err)
	defer func() {
		// 書き込みが完了したら閉じる
		codeWriter.Close()
	}()

	// コードをparseしていく
	for i, path := range pathList {
		parser, err := parser.New(path)
		codeWriter.SetVmFileName(fmt.Sprintf("%d", i))
		// /を_に変換しています、file名に「_」が入らない前提で実装!!!
		// file名に「_」が含まれている場合区別がつかなくなってしまう可能性がある
		// その場合はおとなしくiをuniqueなものとして使う
		// codeWriter.SetVmFileName(strings.ReplaceAll(path, "/", "_"))
		chk.SE(err)

		// 解析をしていく
		for parser.Next() {
			switch parser.CommandType {
			case model.CommandTypeArithmetic:
				codeWriter.WriteArithmetic(parser.Command)
			case model.CommandTypePop, model.CommandTypePush:
				codeWriter.WritePushPop(parser.CommandType, parser.Arg1(), parser.Arg2())
			default:
				panic("まだこれ以外は対応していません")
			}

		}

		// 読み込みが完了したら閉じる
		parser.Close()
	}

}
