package codewriter

import (
	"errors"
	"fmt"
	"os"
	"vm-translator/pkg/common/chk"
	"vm-translator/pkg/model"
)

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
func (c *CodeWriter) WritePushPop(commandType model.CommandType, segment string, index int) {

	switch commandType {
	case model.CommandTypePush:

		switch segment {
		case model.MemorySegmentConstant:

			// write(c.file, )

			// @1 // 1はpush conatant 1の定数
			// D=A
			// @256(SP)
			// M=D
			// TODO SPの値を+1する

			asm := fmt.Sprintf(`
			@%d
			D=A
			@SP
			M=D
			// TODO SPの値を+1する
			`, index)
			write(c.file, asm)

		default:
			chk.SE(errors.New("未実装"))
		}

	case model.CommandTypePop:
	default:
		chk.SE(errors.New("対応していません"))
	}

	// TODO
}

// Close .
func (c *CodeWriter) Close() {
	c.file.Close()
}

// 実際にfileに書き込む
func write(file *os.File, outLine string) {
	_, err := file.WriteString(outLine)
	chk.SE(err)
}
