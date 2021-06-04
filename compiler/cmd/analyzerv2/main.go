package main

import (
	"compiler/pkg/analyzer"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// セットアップや他のmoduleの呼び出しをする

// 1. 入力ファイルのXxx.jackから、Xxx.vmを生成する

const (
	// JackExt Jack言語の拡張子
	JackExt = ".jack"

	// VMExt VM言語の拡張子
	VMExt = ".vm"
)

func main() {
	fmt.Println("==== Jack Analyzer V2 ====")

	// 引数を習得する
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("引数がありません")
		os.Exit(1)
	}

	// source
	dir := args[0]

	// file一覧を習得する
	pathList, err := dirwark(dir)
	if err != nil {
		fmt.Println("ディレクトリの解析でerror")
		os.Exit(1)
	}

	// filter
	pathList = filterExtPathList(pathList, JackExt)

	analyzer.Analyzer(pathList)

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
			// 再起処理
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
