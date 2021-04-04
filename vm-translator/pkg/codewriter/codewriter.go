package codewriter

import "os"

// CodeWriter .
type CodeWriter struct {
	file       *os.File
	VmFileName string // どのVMファイルを変換中か
}

// New .
func New(filePath string) (*CodeWriter, error) {
	fp, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return &CodeWriter{
		file: fp,
	}, nil
}

// SetVmFileName 新しいVMファイルの変換が開始したことを知らせる
func (c *CodeWriter) SetVmFileName(fileName string) {
	c.VmFileName = fileName
}

// WriteArithmetic 与えられた算術コマンドをアセンブリコードに変換して、それを書き込む
func (c *CodeWriter) WriteArithmetic(command string) {
	// TODO
}

// WritePushPop C_PUSH, C_POPコマンドをアセンブリコードに変換し、それを書き込む
func (c *CodeWriter) WritePushPop(command string) {
	// TODO
}

// Close .
func (c *CodeWriter) Close() {
	c.file.Close()
}
