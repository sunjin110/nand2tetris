package codewriter

import (
	"errors"
	"fmt"
	"os"
	"vm-translator/pkg/common/chk"
	"vm-translator/pkg/model"
)

const (
	// init
	initSp = "@256\nD=A\n@SP\nM=D\n"

	// Airthmetic命令セット
	add = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nM=D+M\n@SP\nM=M-1\n"
	sub = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nM=M-D\n@SP\nM=M-1\n"
	neg = "@SP\nA=M-1\nM=-M\n"
	eq  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nD=D-M\nM=-1\n@EQ_%s_%d\nD;JEQ\n@SP\nA=M-1\nA=A-1\nM=0\n(EQ_%s_%d)\n@SP\nM=M-1\n"
	gt  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nD=M-D\nM=-1\n@GT_%s_%d\nD;JGT\n@SP\nA=M-1\nA=A-1\nM=0\n(GT_%s_%d)\n@SP\nM=M-1\n"
	lt  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nD=M-D\nM=-1\n@LT_%s_%d\nD;JLT\n@SP\nA=M-1\nA=A-1\nM=0\n(LT_%s_%d)\n@SP\nM=M-1\n"
	and = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nM=D&M\n@SP\nM=M-1\n"
	or  = "@SP\nA=M-1\nD=M\nM=0\nA=A-1\nM=D|M\n@SP\nM=M-1\n"
	not = "@SP\nA=M-1\nM=!M\n"

	pushConstant = "@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
	pushLocal    = "@%d\nD=A\n@LCL\nM=D+M\nA=M\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@%d\nD=A\n@LCL\nM=M-D\n"
	popLocal     = "@%d\nD=A\n@LCL\nM=D+M\n@SP\nM=M-1\nA=M\nD=M\n@LCL\nA=M\nM=D\n@%d\nD=A\n@LCL\nM=M-D\n"
	pushArg      = "@%d\nD=A\n@ARG\nM=D+M\nA=M\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@%d\nD=A\n@ARG\nM=M-D\n"
	popArg       = "@%d\nD=A\n@ARG\nM=D+M\n@SP\nM=M-1\nA=M\nD=M\n@ARG\nA=M\nM=D\n@%d\nD=A\n@ARG\nM=M-D\n"
	pushThis     = "@%d\nD=A\n@THIS\nM=D+M\nA=M\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@%d\nD=A\n@THIS\nM=M-D\n"
	popThis      = "@%d\nD=A\n@THIS\nM=D+M\n@SP\nM=M-1\nA=M\nD=M\n@THIS\nA=M\nM=D\n@%d\nD=A\n@THIS\nM=M-D\n"
	pushThat     = "@%d\nD=A\n@THAT\nM=D+M\nA=M\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@%d\nD=A\n@THAT\nM=M-D\n"
	popThat      = "@%d\nD=A\n@THAT\nM=D+M\n@SP\nM=M-1\nA=M\nD=M\n@THAT\nA=M\nM=D\n@%d\nD=A\n@THAT\nM=M-D\n"

	pushPointer = "@%s\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n" // index:0 => THIS, index:1 = THAT
	popPointer  = "@SP\nM=M-1\nA=M\nD=M\n@%s\nM=D\n"      // index:0 => THIS, index:1 = THAT
	pushTemp    = "@%d\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n" // %d = 5 + index
	popTemp     = "@SP\nM=M-1\nA=M\nD=M\n@%d\nM=D\n"      // %d = 5 + index

	pushStatic = "@Static_%s_%d\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"
	popStatic  = "@SP\nM=M-1\nA=M\nD=M\n@Static_%s_%d\nM=D\n"

	labelAsm = "(%s)\n"                                    // (LABEL_NAME)
	ifGoto   = "@SP\nA=M-1\nD=M\n@SP\nM=M-1\n@%s\nD;JNE\n" // %s is label name
	goTo     = "@%s\n0;JMP\n"                              // %s is label name

	funcNameSet = "(%s)\n"                                                   // functionのラベルを追加する
	funcInit    = "@%d\nD=A\n@LCL\nM=D+M\nA=M\nM=0\n@%d\nD=A\n@LCL\nM=M-D\n" // 引数の数だけLocalの値を0で初期化する

	returnAsm = "@LCL\nD=M\n@R13\nM=D\n@5\nD=A\n@R13\nA=M-D\nD=M\n@R14\nM=D\n@SP\nA=M-1\nD=M\n@ARG\nA=M\nM=D\n@ARG\nD=M+1\n@SP\nM=D\n@1\nD=A\n@R13\nA=M-D\nD=M\n@THAT\nM=D\n@2\nD=A\n@R13\nA=M-D\nD=M\n@THIS\nM=D\n@3\nD=A\n@R13\nA=M-D\nD=M\n@ARG\nM=D\n@4\nD=A\n@R13\nA=M-D\nD=M\n@LCL\nM=D\n@R14\nA=M\n0;JMP\n"

	// call
	callAsm = "@%s\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@LCL\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@ARG\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@THIS\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@THAT\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@5\nD=A\n@%d\nD=D+A\n@SP\nD=M-D\n@ARG\nM=D\n@SP\nD=M\n@LCL\nM=D\n@%s\n0;JMP\n(%s)\n"

	returnAddressCollLabelPattern = "RET_ADDRESS_CALL_%d" // callの戻り値として使用するlabelのパターン
	uniqueLabelPattern            = "%s$%s"               // function別で同じLabelが使用されるが、それをuniqueなlabelとして変換するためのパターン
)

// CodeWriter .
type CodeWriter struct {
	file        *os.File
	VmFileName  string // どのVMファイルを変換中か
	LabelCount  int32  // ラベルをアトミックにするためのカウント
	CallCount   int32  // callしたときに、戻り値として使用するlabelをアトミックにするためのカウント
	NowFunction string // labelをfunction別で同じものを使用しても大丈夫なようにuniqueにするためのもの
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

// WriteInit VMの初期化、出力ファイルの先頭に配置
func (c *CodeWriter) WriteInit() {
	// write(c.file, initCode)
	write(c.file, initSp)
	c.WriteCall("Sys.init", 0)
}

// WriteArithmetic 与えられた算術コマンドをアセンブリコードに変換して、それを書き込む
func (c *CodeWriter) WriteArithmetic(command string) {

	var asm string
	switch command {
	case model.ArithmeticAdd:
		asm = add
	case model.AirthmeticSub:
		asm = sub
	case model.AirthmeticNeg:
		asm = neg
	case model.AirthmeticEq:
		asm = fmt.Sprintf(eq, c.VmFileName, c.LabelCount, c.VmFileName, c.LabelCount)
		c.LabelCount += 1
	case model.AirthmeticGt:
		asm = fmt.Sprintf(gt, c.VmFileName, c.LabelCount, c.VmFileName, c.LabelCount)
		c.LabelCount += 1
	case model.AirthmeticLt:
		asm = fmt.Sprintf(lt, c.VmFileName, c.LabelCount, c.VmFileName, c.LabelCount)
		c.LabelCount += 1
	case model.AirthmeticAnd:
		asm = and
	case model.AirthmeticOr:
		asm = or
	case model.AirthmeticNot:
		asm = not
	default:
		chk.SE(errors.New("想定していないArthmeticコマンドが渡されました"))
	}

	write(c.file, asm)
}

// WritePushPop C_PUSH, C_POPコマンドをアセンブリコードに変換し、それを書き込む
func (c *CodeWriter) WritePushPop(commandType model.CommandType, segment string, index int) {

	var asm string
	switch segment {
	case model.MemorySegmentConstant: // constant
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushConstant, index)
		case model.CommandTypePop:
			chk.SE(errors.New("constantはpopできません"))
		}

	case model.MemorySegmentLocal: // local
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushLocal, index, index)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popLocal, index, index)
		}

	case model.MemorySegmentArgument: // argument
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushArg, index, index)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popArg, index, index)
		}
	case model.MemorySegmentThis: // this
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushThis, index, index)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popThis, index, index)
		}

	case model.MemorySegmentThat: // that
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushThat, index, index)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popThat, index, index)
		}

	case model.MemorySegmentPointer: // pointer

		var name string
		if index == 0 {
			name = "THIS"
		} else if index == 1 {
			name = "THAT"
		} else {
			chk.SE(errors.New("pointerは0, 1以外の参照はできません"))
		}

		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushPointer, name)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popPointer, name)
		}

	case model.MemorySegmentTemp: // temp -> index + 5
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushTemp, index+5)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popTemp, index+5)
		}

	case model.MemorySegmentStatic: // static
		switch commandType {
		case model.CommandTypePush:
			asm = fmt.Sprintf(pushStatic, c.VmFileName, index)
		case model.CommandTypePop:
			asm = fmt.Sprintf(popStatic, c.VmFileName, index)
		}

	default:
		chk.SE(errors.New("想定していないsegmentが渡されました"))
	}

	if asm == "" {
		chk.SE(errors.New("PushPopに失敗しました"))
	}

	write(c.file, asm)

}

// WriteLabel labelコマンドを行うアセンブリコードの生成
func (c *CodeWriter) WriteLabel(label string) {

	l := fmt.Sprintf(uniqueLabelPattern, c.NowFunction, label) // uniqueなラベルにするため
	asm := fmt.Sprintf(labelAsm, l)
	write(c.file, asm)
}

// WriteGoto gotoコマンドを行うアセンブリコードの生成
func (c *CodeWriter) WriteGoto(label string) {
	// 無条件でジャンプする

	l := fmt.Sprintf(uniqueLabelPattern, c.NowFunction, label) // uniqueなlabelにするため
	asm := fmt.Sprintf(goTo, l)
	write(c.file, asm)
}

// WriteIf ifコマンドを行うアセンブリコードの生成
func (c *CodeWriter) WriteIf(label string) {
	// 一つ前のstackのでーたを取得する
	// そいつが0でない場合は、Labelにループさせる

	l := fmt.Sprintf(uniqueLabelPattern, c.NowFunction, label) // uniqueなlabelにするため
	asm := fmt.Sprintf(ifGoto, l)
	write(c.file, asm)
}

// WriteCall callコマンドを行うアセンブリコードを生成
// funcName callするfunction名
// numArgs 引数の数
func (c *CodeWriter) WriteCall(funcName string, numArgs int) {

	returnAddressCollLabel := fmt.Sprintf(returnAddressCollLabelPattern, c.CallCount)
	// asm := fmt.Sprintf(callAsm, returnAddressCollLabel, numArgs+5, funcName, returnAddressCollLabel)
	asm := fmt.Sprintf(callAsm, returnAddressCollLabel, numArgs, funcName, returnAddressCollLabel)
	write(c.file, asm)

	// call countを+1する
	c.CallCount++
}

// WriteFunction functionコマンドを行うアセンブリコードを生成
func (c *CodeWriter) WriteFunction(funcName string, numLocals int) {

	// uniqueなlabelにするためにfunctionNameを上書きする
	c.NowFunction = funcName

	// functionのラベルを追加する
	asm := fmt.Sprintf(funcNameSet, funcName)

	// numLocals個のLocal変数をすべて0で初期化する
	for i := 0; i < numLocals; i++ {
		asm += fmt.Sprintf(funcInit, i, i)
	}
	write(c.file, asm)
}

// WriteReturn returnコマンドを行うアセンブリコードを生成
func (c *CodeWriter) WriteReturn() {

	// callしたところに戻る作業、
	// リターンアドレスを取得
	// 戻り地を戻ったときのSPの一番上(現在のARG:引数の0番目がそこになるからARGを使って配置する)で取れるようにする
	// その他の戻ったときの、SPとかの位置を復元する

	write(c.file, returnAsm)
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
