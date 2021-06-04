package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

// Dirwark ディレクトリ内のファイルを取得する
func Dirwark(dir string) ([]string, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var pathList []string
	for _, file := range files {

		if file.IsDir() {
			// 再帰処理
			childPathList, err := Dirwark(filepath.Join(dir, file.Name()))
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

// FilterExtPathList 指定した拡張子のみを抽出
func FilterExtPathList(pathList []string, ext string) []string {

	var filterPathList []string
	for _, path := range pathList {
		if strings.HasSuffix(path, ext) {
			filterPathList = append(filterPathList, path)
		}
	}
	return filterPathList
}
