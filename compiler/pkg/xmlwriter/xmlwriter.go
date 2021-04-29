package xmlwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/tokenizer"
	"fmt"
	"os"
)

// XmlWriter .
type XmlWriter struct {
	file *os.File
}

// New .
func New(filePath string) (*XmlWriter, error) {
	fp, err := os.Create(filePath)
	chk.SE(err)

	write(fp, "<tokens>\n")
	return &XmlWriter{
		file: fp,
	}, nil
}

// WriteToken .
func (x *XmlWriter) WriteToken(tokenType tokenizer.TokenType, tokenVal string) {

	value := tokenVal
	switch value {
	case "<":
		value = "&lt;"
	case ">":
		value = "&gt;"
	case "&":
		value = "&amp;"
	}

	outLine := fmt.Sprintf("<%s> %s </%s>\n", tokenType, value, tokenType)
	write(x.file, outLine)

}

// Close .
func (x *XmlWriter) Close() {

	write(x.file, "</tokens>\n")

	x.file.Close()
}

// 実際にfileに書き込む
func write(file *os.File, outLine string) {
	_, err := file.WriteString(outLine)
	chk.SE(err)
}
