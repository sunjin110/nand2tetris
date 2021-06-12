package compiler

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/tokenizer"
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

		// TODO SymbolTable module

		// VMWriter Module

	}

}
