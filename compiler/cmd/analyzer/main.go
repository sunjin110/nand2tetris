package main

import (
	"compiler/pkg/common/fileutil"
	"compiler/pkg/controller/analyzer"
	"flag"
	"fmt"
	"os"
)

// セットアップや他moduleの呼び出しを行う

// 1. 入力ファイルのXxx.jackから、JackTokenizerを生成する
// 2. Xxx.xmlという名前の出力ファイルを作り、それに書き込みを行う準備をする
// 3. 入力である JackTokenizerを出力ファイルへコンパイするすために、CompilationEngineを使用する

const (
	// JackExt Jack言語の拡張子
	JackExt = ".jack"
)

func main() {
	fmt.Println("==== Jack Analyzer ====")

	// 引数を取得する
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("引数がありません")
		os.Exit(1)
	}

	// source
	dir := args[0]

	// file一覧を取得する
	pathList, err := fileutil.Dirwark(dir)
	if err != nil {
		fmt.Println("ディレクトリの解析でerror")
		os.Exit(1)
	}

	// filter
	pathList = fileutil.FilterExtPathList(pathList, JackExt)

	analyzer.Analyzer(pathList)
}
