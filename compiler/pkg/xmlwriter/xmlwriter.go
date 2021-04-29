package xmlwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/tokenizer"
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
	return &XmlWriter{
		file: fp,
	}, nil
}

// WriteToken .
func (x *XmlWriter) WriteToken(token string, tokenType tokenizer.TokenType, keyword tokenizer.KeyWord) {

}

// Close .
func (x *XmlWriter) Close() {

	// TODO write「</tokens>」

	x.file.Close()
}
