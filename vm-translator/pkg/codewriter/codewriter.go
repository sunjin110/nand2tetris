package codewriter

import (
	"errors"
	"fmt"
	"os"
	"vm-translator/pkg/common/chk"
	"vm-translator/pkg/model"
)

const (
	add = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nM=D+M\n@SP\nM=M-1\n"
	sub = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nM=M-D\n@SP\nM=M-1\n"
	neg = "@SP\nA=M-1\nM=-M\n"
	eq  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nD=D-M\nM=-1\n@EQ_%s_%d\nD;JEQ\n@SP\nA=M-1\nA=A-1\nM=0\n(EQ_%s_%d)\n@SP\nM=M-1\n"
	gt  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nD=M-D\nM=-1\n@GT_%s_%d\nD;JGT\n@SP\nA=M-1\nA=A-1\nM=0\n(GT_%s_%d)\n@SP\nM=M-1\n"
	lt  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nD=M-D\nM=-1\n@LT_%s_%d\nD;JLT\n@SP\nA=M-1\nA=A-1\nM=0\n(LT_%s_%d)\n@SP\nM=M-1\n"
	and = ""
	or  = ""
	not = ""

	pushConstant = "@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
)

// CodeWriter .
type CodeWriter struct {
	file       *os.File
	VmFileName string // どのVMファイルを変換中か
	LabelCount int32  // ラベルをアトミックにするためのカウント
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

	switch command {
	case "add":
		asm := add
		write(c.file, asm)
	case "sub":
		asm := sub
		write(c.file, asm)
	case "neg":
		asm := neg
		write(c.file, asm)
	case "eq":
		asm := fmt.Sprintf(eq, c.VmFileName, c.LabelCount, c.VmFileName, c.LabelCount)
		c.LabelCount += 1
		write(c.file, asm)
	case "gt":
		c.LabelCount += 1
		asm := fmt.Sprintf(gt, c.VmFileName, c.LabelCount, c.VmFileName, c.LabelCount)
		write(c.file, asm)
	case "lt":
		c.LabelCount += 1
		asm := fmt.Sprintf(lt, c.VmFileName, c.LabelCount, c.VmFileName, c.LabelCount)
		write(c.file, asm)
	default:
		chk.SE(errors.New("未実装"))
	}

	// TODO
}

// WritePushPop C_PUSH, C_POPコマンドをアセンブリコードに変換し、それを書き込む
func (c *CodeWriter) WritePushPop(commandType model.CommandType, segment string, index int) {

	switch commandType {
	case model.CommandTypePush:

		switch segment {
		case model.MemorySegmentConstant:

			// push constant %d
			asm := fmt.Sprintf(pushConstant, index)
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
