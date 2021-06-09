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
	file           *os.File
	class          *compilation_engine.Class
	symbolTable    *symboltable.SymbolTable
	subRoutineName string // 現在のsubRoutineの名前
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
	className := writer.class.ClassName
	subRoutineDecList := writer.class.SubRoutineDecList
	for _, subRoutineDec := range subRoutineDecList {
		writer.writeSubRoutine(className, subRoutineDec)
	}
}

// writeSubRoutine .
func (writer *VMWriter) writeSubRoutine(className string, subRoutineDec *compilation_engine.SubRoutineDec) {

	// function Main.main 2
	writer.writeFunction(fmt.Sprintf("%s.%s", className, subRoutineDec.SubRoutineName), int32(len(subRoutineDec.ParameterList)))

	// setSubRoutineName
	writer.subRoutineName = subRoutineDec.SubRoutineName

	for _, statement := range subRoutineDec.SubRoutineBody.StatementList {
		writer.writeStatement(statement)
	}
}

// writeStatement .
func (writer *VMWriter) writeStatement(statement compilation_engine.Statement) {

	switch statement.GetStatementType() {
	case compilation_engine.LetStatementPrefix:
		writer.writeLetStatement(statement.(*compilation_engine.LetStatement))
	case compilation_engine.DoStatementPrefix:
		writer.writeDoStatement(statement.(*compilation_engine.DoStatement))
	case compilation_engine.ReturnStatementPrefix:
		writer.writeReturnStatement(statement.(*compilation_engine.ReturnStatement))
	default:
		chk.SE(fmt.Errorf("writeStatement: 宣言していないstatementが渡されました:%s", statement.GetStatementType()))
	}

}

// writeLetStatement .
func (writer *VMWriter) writeLetStatement(letStatement *compilation_engine.LetStatement) {
	// 先にexpressionを処理
	writer.writeExpression(letStatement.Expression)

	// TODO arrayExpressionを考慮して実装する

}

// writeDoStatement .
func (writer *VMWriter) writeDoStatement(doStatement *compilation_engine.DoStatement) {
	writer.writeSubroutineCall(doStatement.SubroutineCall)

	// do statementは戻り値を使用しないため、tempにstackにある戻り値をtmpにpopしてしまう
	// pop temp 0
	writer.writePop(segmentTemp, 0)
}

// writeReturnStatement .
func (writer *VMWriter) writeReturnStatement(returnStatement *compilation_engine.ReturnStatement) {
	if returnStatement.ReturnExpression != nil {
		writer.writeExpression(returnStatement.ReturnExpression)
	} else {
		// returnがない場合は、無理やり0をpushする
		writer.writePush(segmentConst, 0)
	}

	writer.writeReturn()
}

// writeSubroutineCall subroutineCallのvmを記述する
func (writer *VMWriter) writeSubroutineCall(subroutineCall *compilation_engine.SubRoutineCall) {

	// TODO 現在、funcitonしか対応していない、methodやconstructが来ても対応できるようにする??

	// ()内の計算式を習得する、そんで書く
	writer.writeExpressionList(subroutineCall.ExpressionList)

	name := subroutineCall.SubRoutineName
	if subroutineCall.ClassOrVarName != "" {
		name = fmt.Sprintf("%s.%s", subroutineCall.ClassOrVarName, subroutineCall.SubRoutineName)
	}

	writer.writeCall(name, int32(len(subroutineCall.ExpressionList)))
}

// writeExpressionList .
func (writer *VMWriter) writeExpressionList(expressionList []*compilation_engine.Expression) {
	for _, expression := range expressionList {
		writer.writeExpression(expression)
	}
}

// writeExpression 演算はStackマシンなので、逆ポーランド記法で出力する
func (writer *VMWriter) writeExpression(expression *compilation_engine.Expression) {

	// init term
	writer.writeTerm(expression.InitTerm)

	// opTermList
	for _, opTerm := range expression.OpTermList {

		//term
		writer.writeTerm(opTerm.OpTerm)

		// opTerm.Operation
		writer.writeOperation(opTerm.Operation)
	}
}

// writeTerm .
func (writer *VMWriter) writeTerm(term compilation_engine.Term) {
	switch term.GetTermType() {
	case compilation_engine.IntegerConstType:
		writer.writeIntegerConstTerm(term.(*compilation_engine.IntegerConstTerm))
	case compilation_engine.ExpressionType:
		writer.writeExpressionTerm(term.(*compilation_engine.ExpressionTerm))
	case compilation_engine.ValNameConstType:
		writer.writeValNameConstTerm(term.(*compilation_engine.ValNameConstantTerm))
	case compilation_engine.SubRoutineCallType:
		writer.writeSubroutineCall(term.(*compilation_engine.SubRoutineCall))
	case compilation_engine.UnaryOpTermType:
		writer.writeUnaryOpTerm(term.(*compilation_engine.UnaryOpTerm))
		// TODO more case
	default:
		chk.SE(fmt.Errorf("writeTerm:想定していないterm typeが来ました:%s", term.GetTermType()))

	}
}

// writeIntegerConstTerm .
func (writer *VMWriter) writeIntegerConstTerm(integerConstTerm *compilation_engine.IntegerConstTerm) {
	// constantの場合は実数を入れればいいので、indexは実数を入れる
	writer.writePush(segmentConst, int32(integerConstTerm.Val))
}

// writeExpressionTerm ex: (1 + 2)
func (writer *VMWriter) writeExpressionTerm(expressionTerm *compilation_engine.ExpressionTerm) {
	writer.writeExpression(expressionTerm.Expression)
}

// writeValNameConstTerm .
func (writer *VMWriter) writeValNameConstTerm(valNameConstTerm *compilation_engine.ValNameConstantTerm) {

	// 変数のsymbol情報を習得する
	subroutineSymbolTable := writer.getCurrentSubroutineSymbolTable()
	symbol := subroutineSymbolTable.SymbolMap[valNameConstTerm.ValName]

	// LocalにPopする
	// TODO 本当にLocalだけで大丈夫か？を確認する
	writer.writePop(segmentLocal, symbol.Num)
}

// writeUnaryOpTerm .
func (writer *VMWriter) writeUnaryOpTerm(unaryOpTerm *compilation_engine.UnaryOpTerm) {
	writer.writeTerm(unaryOpTerm.Term)
	switch unaryOpTerm.UnaryOp {
	case compilation_engine.HypenUnaryOp: // -
		writer.writeArithmetic(AirthmeticNeg)
	case compilation_engine.TildeUnaryOp: // ~
		writer.writeArithmetic(AirthmeticNot)
	default:
		chk.SE(fmt.Errorf("writeUnaryOpTerm:想定していないUnaryOpが来ました:%v", unaryOpTerm.UnaryOp))
	}
}

// writeOperation .
func (writer *VMWriter) writeOperation(op compilation_engine.Op) {

	switch op {
	case compilation_engine.PlusOp: // +
		writer.writeArithmetic(ArithmeticAdd)
	case compilation_engine.MinusOp: // -
		writer.writeArithmetic(AirthmeticSub)
	case compilation_engine.AsteriskOp: // *
		// call Match.multiply 2
		name, nArgs := getMathMultiply()
		writer.writeCall(name, nArgs)
	case compilation_engine.SlashOp: // /
		// call Math.divide 2
		name, nArgs := getMathDivide()
		writer.writeCall(name, nArgs)
	case compilation_engine.AndOp: // &
		writer.writeArithmetic(AirthmeticAnd)
	case compilation_engine.PipeOp: // |
		writer.writeArithmetic(AirthmeticOr)
	case compilation_engine.LessThanOp: // <
		writer.writeArithmetic(AirthmeticLt)
	case compilation_engine.GreaterThanOp: // >
		writer.writeArithmetic(AirthmeticGt)
	case compilation_engine.EqlOp: // =
		writer.writeArithmetic(AirthmeticEq)
	default:
		chk.SE(fmt.Errorf("writeOperation:想定しないoperationが渡されました:%v", op))
	}
}

// getCurrentSubroutineSymbolTable 現在のSubroutineのSymbolTableを習得する
func (writer *VMWriter) getCurrentSubroutineSymbolTable() *symboltable.SubroutineSymbolTable {
	return writer.symbolTable.SubroutineSymbolTableMap[writer.subRoutineName]
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
