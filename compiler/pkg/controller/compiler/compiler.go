package compiler

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/symboltable"
	"compiler/pkg/tokenizer"
	"compiler/pkg/writer/vmwriter"
	"log"
	"strings"
)

// .vmを生成する
// CompilationEngine, SymbolTable, VMWriterを用意て出力ファイルへ書き込みをする

// Compile Vmに変換する
func Compile(pathList []string) {
	jsonutil.Print(pathList)

	for _, path := range pathList {

		t, err := tokenizer.New(path)
		chk.SE(err)

		// 構文解析Engine
		compilationEngine := compilation_engine.New(t)
		compilationEngine.Start()

		// SymbolTableEngine
		symbolTableEngine := symboltable.New(compilationEngine.Class)
		symbolTableEngine.Start()

		log.Println("symboltable is ", jsonutil.Marshal(symbolTableEngine.SymbolTable))

		outputFileName := getOutputFileName(path)

		// vm writerを作成
		vmwriter := vmwriter.New(outputFileName, compilationEngine.Class, symbolTableEngine.SymbolTable)

		// vmを出力
		vmwriter.WriteVM()
		vmwriter.FileClose()
	}

}

// outputFileName 出力するファイル名を取得する
func getOutputFileName(path string) string {
	return strings.Split(path, ".")[0] + ".vm"
}
