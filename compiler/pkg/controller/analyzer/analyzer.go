package analyzer

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/tokenizer"
	"compiler/pkg/writer/xmlwriter"
	"strings"
)

// Analyzer Xmlに変換する
func Analyzer(pathList []string) {
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

// writeXML xmlとして出力する
func writeXML(path string, c *compilation_engine.Class) {

	// xmlに変換する
	outputFileName := strings.Split(path, ".")[0] + ".xml"

	xw := xmlwriter.New(outputFileName, c)
	xw.WriteParser()
}
