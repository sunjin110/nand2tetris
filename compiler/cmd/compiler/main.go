package main

import (
	"compiler/pkg/common/fileutil"
	"compiler/pkg/controller/compiler"
	"flag"
	"fmt"
	"os"
)

// Jackコードを読み込んでvmファイルを作成する

// JackExt Jack言語の拡張子
const JackExt = ".jack"

func main() {
	fmt.Println("==== Jack Compiler ====")

	// 引数を取得する
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("引数がありません")
		os.Exit(1)
	}

	// source
	dir := args[0]

	// file一覧を習得する
	pathList, err := fileutil.Dirwark(dir)
	if err != nil {
		fmt.Println("ディレクトリの解析でerror")
		os.Exit(1)
	}

	// filter
	pathList = fileutil.FilterExtPathList(pathList, JackExt)

	// compile
	compiler.Compile(pathList)
}
