package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vm-translator/pkg/translate"
)

const (
	// VmExt 拡張子
	VmExt = ".vm"

	// AsmExt 拡張子
	AsmExt = ".asm"
)

// STEP1 スタック算術コマンドの作成

// STEP2 メモリアクセスコマンドの作成

func main() {
	log.Println("VM translator")

	// 引数を取得する
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("引数がありません")
		os.Exit(1)
	}
	dir := args[0]

	// file一覧を取得する
	pathList, err := dirwark(dir)
	if err != nil {
		fmt.Println("ディレクトリの解析でerror")
		panic(err)
	}

	// filter
	pathList = filterExtPathList(pathList, VmExt)

	// output name
	outputFileName := fmt.Sprintf("%s%s", dir, AsmExt)

	// translate
	translate.Translate(outputFileName, pathList)
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
