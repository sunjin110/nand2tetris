package xmlwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/compilation_engine"
	"fmt"
	"os"
)

// XmlWriter .
type XmlWriter struct {
	file      *os.File
	class     *compilation_engine.Class
	nestDepth int // ネストの深さ(tab)
}

// New XmlWriterを作成する
func New(filePath string, class *compilation_engine.Class) *XmlWriter {
	return &XmlWriter{
		file:      createFile(filePath),
		class:     class,
		nestDepth: 0,
	}
}

// WriteParser パーサで解析した内容を書き出す
func (writer *XmlWriter) WriteParser() {

	writer.writeClass()

}

// writeClass Classからxmlのファイルを作成する
func (writer *XmlWriter) writeClass() error {

	writer.writeFile("<class>")

	writer.writeFile("</class>")

	return nil
}

// writeFile
// <key> value </key>
func (writer *XmlWriter) writeFile(value string) {

	// nest
	var nest string
	for i := 0; i < writer.nestDepth; i++ {
		nest += "\t"
	}

	_, err := writer.file.WriteString(fmt.Sprintf("%s%s\n", nest, value))
	chk.SE(err)
}

// func writeClass(f *os.File, class *compilation_engine.Class) {

// 	// class
// 	writeLine(f, "<class>")

// 	writeLine(f, "</class>")
// }

// // 実際にfileに書き込む
// func writeLine(file *os.File, outLine string) {
// 	_, err := file.WriteString(outLine + "\n")
// 	chk.SE(err)
// }

// createFile fileを作成する
func createFile(filePath string) *os.File {
	fp, err := os.Create(filePath)
	chk.SE(err)
	return fp
}

// import (
// 	"compiler/pkg/common/chk"
// 	"compiler/pkg/tokenizer"
// 	"fmt"
// 	"os"
// )

// // XmlWriter .
// type XmlWriter struct {
// 	file *os.File
// }

// // New .
// func New(filePath string) (*XmlWriter, error) {
// 	fp, err := os.Create(filePath)
// 	chk.SE(err)

// 	write(fp, "<tokens>\n")
// 	return &XmlWriter{
// 		file: fp,
// 	}, nil
// }

// // WriteToken .
// func (x *XmlWriter) WriteToken(tokenType tokenizer.TokenType, tokenVal string) {

// 	value := tokenVal
// 	switch value {
// 	case "<":
// 		value = "&lt;"
// 	case ">":
// 		value = "&gt;"
// 	case "&":
// 		value = "&amp;"
// 	}

// 	outLine := fmt.Sprintf("<%s> %s </%s>\n", tokenType, value, tokenType)
// 	write(x.file, outLine)

// }

// // Close .
// func (x *XmlWriter) Close() {

// 	write(x.file, "</tokens>\n")

// 	x.file.Close()
// }
