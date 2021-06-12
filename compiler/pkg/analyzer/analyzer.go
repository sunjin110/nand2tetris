package analyzer

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/tokenizer"
	"compiler/pkg/vmwriter"
	"compiler/pkg/xmlwriter"
	"strings"
)

// AnalyzerToXML Xmlに変換する
func AnalyzerToXML(pathList []string) {
	jsonutil.Print(pathList)

	for _, path := range pathList {

		t, err := tokenizer.New(path)
		chk.SE(err)

		compilationEngine := compilation_engine.New(t)

		compilationEngine.Start()

		c := compilationEngine.Class

		// writeXml
		writeXML(path, c)
	}
}

// AnalyzerToVM Vmに変換する
func AnalyzerToVM(pathList []string) {
	jsonutil.Print(pathList)

	for _, path := range pathList {

		t, err := tokenizer.New(path)
		chk.SE(err)

		compilationEngine := compilation_engine.New(t)

		compilationEngine.Start()

		c := compilationEngine.Class

		// writeVM
		writeVM(path, c)
	}
}

// writeXML xmlとして出力する
func writeXML(path string, c *compilation_engine.Class) {

	// xmlに変換する
	outputFileName := strings.Split(path, ".")[0] + ".xml"

	xw := xmlwriter.New(outputFileName, c)
	xw.WriteParser()
}

// writeVM vmとして出力する
func writeVM(path string, c *compilation_engine.Class) {

	// vmに変換する
	outputFileName := strings.Split(path, ".")[0] + ".vm"

	vmw := vmwriter.New(outputFileName, c)
	vmw.Write()
}
