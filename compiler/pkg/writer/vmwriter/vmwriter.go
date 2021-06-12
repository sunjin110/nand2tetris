package vmwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/fileutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/symboltable"
	"fmt"
	"os"
)

// VMWriter .
type VMWriter struct {
	file        *os.File
	class       *compilation_engine.Class
	symbolTable *symboltable.SymbolTable
}

// New VMWriterを作成する
func New(filePath string, class *compilation_engine.Class, symbolTable *symboltable.SymbolTable) *VMWriter {
	return &VMWriter{
		file:        fileutil.CreateFile(filePath),
		class:       class,
		symbolTable: symbolTable,
	}
}

// FileClose ファイルの書き込みを終了する
func (writer *VMWriter) FileClose() {
	writer.file.Close()
}

// WriteVM VMファイルを書く
func (writer *VMWriter) WriteVM() {

	// TODO
	//

}

// writePush pushコマンドを書く
// segment -> CONST, ARG, LOCAL, STATIC, THIS, THAT, POINTER, TMEP
// index is 整数
func (writer *VMWriter) writePush(segment string, index int32) {
	writer.write(fmt.Sprintf("push %s %d", segment, index))
}

// writePop popコマンド
// segment -> CONST, ARG, LOCAL, STATIC, THIS, THAT, POINTER, TMEP
// index is 整数
func (writer *VMWriter) writePop(segment string, index int32) {
	writer.write(fmt.Sprintf("pop %s %d", segment, index))
}

// writeArithmetic 算術コマンドを書く
// command -> ADD, SUB, NEG, EQ, GT, LT, AND, OR, NOT
func (writer *VMWriter) writeArithmetic(command string) {
	writer.write(command)
}

// writeLable labelコマンドを書く
// TODO
func (writer *VMWriter) writeLabel(label string) {
	writer.write(fmt.Sprintf("label %s", label))
}

// writeGoto gotoコマンドを書く
func (writer *VMWriter) writeGoto(label string) {
	writer.write(fmt.Sprintf("goto %s", label))
}

// writeIf if-gotoコマンドをかく
func (writer *VMWriter) writeIf(label string) {
	writer.write(fmt.Sprintf("if-goto %s", label))
}

// writeCall callコマンドを書く
func (writer *VMWriter) writeCall(name string, nArgs int32) {
	writer.write(fmt.Sprintf("call %s %d", name, nArgs))

}

// writeFunction functionコマンドを書く
func (writer *VMWriter) writeFunction(name string, nLocals int32) {
	writer.write(fmt.Sprintf("function %s %d", name, nLocals))
}

// writeReturn returnコマンドを書く
func (writer *VMWriter) writeReturn() {
	writer.write("return")
}

// write
func (writer *VMWriter) write(value string) {
	_, err := writer.file.WriteString(fmt.Sprintf("%s\n", value))
	chk.SE(err)
}
