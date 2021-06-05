package vmwriter

import (
	"compiler/pkg/common/fileutil"
	"compiler/pkg/compilation_engine"
	"compiler/pkg/symboltable"
	"os"
)

// VMWriter .
type VMWriter struct {
	file        *os.File
	class       *compilation_engine.Class
	symbolTable *symboltable.SymbolTable
}

// New VMWriterを作成する
// TODO
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
}

// writePush pushコマンドを書く
// TODO
// segment -> CONST, ARG, LOCAL, STATIC, THIS, THAT, POINTER, TMEP
// index is 整数
func (writer *VMWriter) writePush(segment string, index int32) {
}

// writePop popコマンド
// TODO
// segment -> CONST, ARG, LOCAL, STATIC, THIS, THAT, POINTER, TMEP
// index is 整数
func (writer *VMWriter) writePop(segment string, index int32) {
}

// writeArithmetic 算術コマンドを書く
// TODO
// command -> ADD, SUB, NEG, EQ, GT, LT, AND, OR, NOT
func (writer *VMWriter) writeArithmetic(command string) {
}

// writeLable labelコマンドを書く
// TODO
func (writer *VMWriter) writeLabel(label string) {
}

// writeGoto gotoコマンドを書く
// TODO
func (writer *VMWriter) writeGoto(lable string) {
}

// writeIf if-gotoコマンドをかく
// TODO
func (writer *VMWriter) writeIf(label string) {
}

// writeCall callコマンドを書く
// TODO
func (writer *VMWriter) writeCall(name string, nArgs int32) {
}

// writeFunction functionコマンドを書く
// TODO
func (writer *VMWriter) writeFunction(name string, nLocals int32) {
}

// writeReturn returnコマンドを書く
// TODO
func (writer *VMWriter) writeReturn() {
}
