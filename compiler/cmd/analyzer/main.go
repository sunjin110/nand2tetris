package main

import (
	"compiler/pkg/analyzer"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// セットアップや他moduleの呼び出しを行う

// 1. 入力ファイルのXxx.jackから、JackTokenizerを生成する
// 2. Xxx.xmlという名前の出力ファイルを作り、それに書き込みを行う準備をする
// 3. 入力である JackTokenizerを出力ファイルへコンパイするすために、CompilationEngineを使用する

const (
	// JackExt Jack言語の拡張子
	JackExt = ".jack"

	// XMLMode xmlを出力するmode
	XMLMode = "xml"

	// VMMode vmを出力するmode
	VMMode = "vm"
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
	pathList, err := dirwark(dir)
	if err != nil {
		fmt.Println("ディレクトリの解析でerror")
		os.Exit(1)
	}

	// filter
	pathList = filterExtPathList(pathList, JackExt)

	mode := args[1]
	switch mode {
	case XMLMode:
		analyzer.AnalyzerToXML(pathList)
	case VMMode:
		analyzer.AnalyzerToVM(pathList)
	default:
		fmt.Println("第2引数にmode[vm|xml]を指定してください")
		os.Exit(1)
	}
}

// dirwark ディレクトリ内のファイルを取得する
func dirwark(dir string) ([]string, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var pathList []string
	for _, file := range files {

		if file.IsDir() {
			// 再帰処理
			childPathList, err := dirwark(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}

			pathList = append(pathList, childPathList...)
		} else {
			pathList = append(pathList, filepath.Join(dir, file.Name()))
		}
	}

	return pathList, nil
}

// filterExtPathList 指定した拡張子のみを抽出
func filterExtPathList(pathList []string, ext string) []string {

	var filterPathList []string
	for _, path := range pathList {
		if strings.HasSuffix(path, ext) {
			filterPathList = append(filterPathList, path)
		}
	}
	return filterPathList
}
