package analyzer

import "compiler/pkg/common/jsonutil"

// Analyzer Xmlに変換する
func Analyzer(outputFileName string, pathList []string) {
	jsonutil.Print(pathList)
}
