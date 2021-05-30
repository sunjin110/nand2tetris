package xmlwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/compilation_engine"
	"fmt"
	"os"
)

const (
	keyword    = "keyword"
	identifier = "identifier"
	symbol     = "symbol"
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
func (w *XmlWriter) writeClass() error {

	w.write("<class>")

	// class name
	w.incNest()

	// class init
	w.write(getKeywordXml("class"))
	w.write(getIdentifierXml(w.class.ClassName))
	w.write(getSymbolXml("{"))

	// class var dec
	w.writeClassVarDec(w.class.ClassVarDecList)

	// subroutine dec
	w.writeSubroutineDec(w.class.SubRoutineDecList)

	// class defer
	w.write(getSymbolXml("}"))

	w.decNest()

	w.write("</class>")

	return nil
}

// writeClassVarDec .
func (w *XmlWriter) writeClassVarDec(classVarDecList []*compilation_engine.ClassVarDec) {
	for _, classVarDec := range classVarDecList {

		w.write("<classVarDec>")
		w.incNest()

		// var kind
		w.write(getKeywordXml(string(classVarDec.VarKind)))

		// type
		// w.write(getKeywordXml(string(classVarDec.VarType)))
		w.writeVariableType(classVarDec.VarType)

		// names
		for i, varName := range classVarDec.VarNameList {
			w.write(getIdentifierXml(varName))

			// 最後でない場合は「,」を追加
			if len(classVarDec.VarNameList) != i+1 {
				w.write(getSymbolXml(","))
			}
		}

		// ;
		w.write(getSymbolXml(";"))

		w.decNest()
		w.write("</classVarDec>")
	}
}

// writeSubroutineDec .
func (w *XmlWriter) writeSubroutineDec(subRoutineDecList []*compilation_engine.SubRoutineDec) {

	for _, subRoutineDec := range subRoutineDecList {

		w.write("<subroutineDec>")
		w.incNest()

		w.write(getKeywordXml(string(subRoutineDec.RoutineKind)))

		w.writeVariableType(subRoutineDec.ReturnType)

		w.write(getIdentifierXml(subRoutineDec.SubRoutineName))

		// (
		w.write(getSymbolXml("("))

		// parameterList
		w.write("<parameterList>")
		w.incNest()
		for i, parameter := range subRoutineDec.ParameterList {

			w.writeVariableType(parameter.ParamType)
			w.write(getIdentifierXml(parameter.ParamName))

			if len(subRoutineDec.ParameterList) != i+1 {
				w.write(getSymbolXml(","))
			}

		}
		w.decNest()
		w.write("</parameterList>")

		// )
		w.write(getSymbolXml(")"))

		// subroutineBody
		w.writeSubroutineBody(subRoutineDec.SubRoutineBody)

		w.decNest()
		w.write("</subroutineDec>")

	}

}

// writeSubroutineBody .
func (w *XmlWriter) writeSubroutineBody(subRoutineBody *compilation_engine.SubRoutineBody) {

	defer func() {
		// }
		w.write(getSymbolXml("}"))

		w.decNest()
		w.write("</subroutineBody>")
	}()

	w.write("<subroutineBody>")
	w.incNest()

	// {
	w.write(getSymbolXml("{"))

	// check
	if subRoutineBody == nil {
		return
	}

	// var dec
	w.writeVarDec(subRoutineBody.VarDecList)

	// statement
	w.writeStatement(subRoutineBody.StatementList)

}

// writeVarDec .
func (w *XmlWriter) writeVarDec(varDecList []*compilation_engine.VarDec) {

	for _, varDec := range varDecList {

		w.write("<varDec>")
		w.incNest()

		w.write(getKeywordXml("var"))

		w.writeVariableType(varDec.Type)

		for i, varName := range varDec.NameList {
			w.write(getIdentifierXml(varName))
			if len(varDec.NameList) != i+1 {
				w.write(getSymbolXml(","))
			}
		}

		// ;
		w.write(getSymbolXml(";"))

		w.decNest()
		w.write("</varDec>")

	}
}

// writeStatement .
func (w *XmlWriter) writeStatement(statementList []compilation_engine.Statement) {

	w.write("<statements>")
	w.incNest()

	for _, statement := range statementList {

		switch statement.GetStatementType() {
		case compilation_engine.LetStatementPrefix:
			w.writeLetStatement(statement.(*compilation_engine.LetStatement))
		case compilation_engine.IfStatementPrefix:
			w.writeIfStatement(statement.(*compilation_engine.IfStatement))
		case compilation_engine.WhileStatementPrefix:
			w.writeWhileStatement(statement.(*compilation_engine.WhileStatement))
		case compilation_engine.DoStatementPrefix:
			w.writeDoStatement(statement.(*compilation_engine.DoStatement))
		case compilation_engine.ReturnStatementPrefix:
			w.writeReturnStatement(statement.(*compilation_engine.ReturnStatement))
		default:
			chk.SE(fmt.Errorf("writeStatement: 宣言していないstatementが渡されました:%s", statement.GetStatementType()))
		}
	}

	w.decNest()
	w.write("</statements>")
}

// writeLetStatement .
func (w *XmlWriter) writeLetStatement(letStatement *compilation_engine.LetStatement) {

	w.write("<letStatement>")
	w.incNest()

	w.write(getKeywordXml("let"))
	w.write(getIdentifierXml(letStatement.DestVarName))

	// array
	if letStatement.ArrayExpression != nil {
		w.write(getSymbolXml("["))
		w.writeExpression(letStatement.ArrayExpression)
		w.write(getSymbolXml("]"))
	}

	// =
	w.write(getSymbolXml("="))

	// expression
	w.writeExpression(letStatement.Expression)

	w.write(getSymbolXml(";"))

	w.decNest()
	w.write("</letStatement>")

}

// writeIfStatement .
func (w *XmlWriter) writeIfStatement(ifStatement *compilation_engine.IfStatement) {

	w.write("<ifStatement>")
	w.incNest()

	w.write(getKeywordXml("if"))

	// (
	w.write(getSymbolXml("("))

	if ifStatement.ConditionalExpression != nil {
		w.writeExpression(ifStatement.ConditionalExpression)
	}

	// )
	w.write(getSymbolXml(")"))

	// {
	w.write(getSymbolXml("{"))

	w.writeStatement(ifStatement.StatementList)

	// }
	w.write(getSymbolXml("}"))

	// else
	if len(ifStatement.ElseStatementList) > 0 {
		w.write(getKeywordXml("else"))
		w.write(getSymbolXml("{"))
		w.writeStatement(ifStatement.ElseStatementList)
		w.write(getSymbolXml("}"))
	}

	w.decNest()
	w.write("</ifStatement>")
}

// writeWhileStatement .
func (w *XmlWriter) writeWhileStatement(whileStatement *compilation_engine.WhileStatement) {

	w.write("<whileStatement>")
	w.incNest()

	w.write(getKeywordXml("while"))

	// (
	w.write(getSymbolXml("("))

	if whileStatement.ConditionalExpression != nil {
		w.writeExpression(whileStatement.ConditionalExpression)
	}

	// )
	w.write(getSymbolXml(")"))

	// {
	w.write(getSymbolXml("{"))

	// statement
	w.writeStatement(whileStatement.StatementList)

	// }
	w.write(getSymbolXml("}"))

	w.decNest()
	w.write("</whileStatement>")

}

// writeDoStatement .
func (w *XmlWriter) writeDoStatement(doStatement *compilation_engine.DoStatement) {

	w.write("<doStatement>")
	w.incNest()

	w.write(getKeywordXml("do"))

	w.writeSubRoutineCall(doStatement.SubroutineCall)

	w.write(getSymbolXml(";"))

	w.decNest()
	w.write("</doStatement>")
}

// writeReturnStatement .
func (w *XmlWriter) writeReturnStatement(returnStatement *compilation_engine.ReturnStatement) {

	w.write("<returnStatement>")
	w.incNest()

	w.write(getKeywordXml("return"))

	if returnStatement.ReturnExpression != nil {
		w.writeExpression(returnStatement.ReturnExpression)
	}

	w.write(getSymbolXml(";"))

	w.decNest()
	w.write("</returnStatement>")
}

// writeExpressionList .
func (w *XmlWriter) writeExpressionList(expressionList []*compilation_engine.Expression) {
	w.write("<expressionList>")
	w.incNest()

	for _, expression := range expressionList {
		w.writeExpression(expression)
	}

	w.decNest()
	w.write("</expressionList>")
}

// writeExpression .
func (w *XmlWriter) writeExpression(expression *compilation_engine.Expression) {

	w.write("<expression>")
	w.incNest()

	// term
	w.writeTerm(expression.InitTerm)

	// op term
	for _, opTerm := range expression.OpTermList {

		// operation
		w.write(getSymbolXml(string(opTerm.Operation)))

		// term
		w.writeTerm(opTerm.OpTerm)
	}

	w.decNest()
	w.write("</expression>")
}

// writeTerm TODO
func (w *XmlWriter) writeTerm(term compilation_engine.Term) {

	w.write("<term>")
	w.incNest()

	switch term.GetTermType() {

	case compilation_engine.IntegerConstType:
		w.writeIntegerConstTerm(term.(*compilation_engine.IntegerConstTerm))
	case compilation_engine.StringConstType:
		w.writeStringConstTerm(term.(*compilation_engine.StringConstTerm))
	case compilation_engine.KeyWordConstType:
		w.writeKeyWordConstTerm(term.(*compilation_engine.KeyWordConstTerm))
	case compilation_engine.ValNameConstType:
		w.writeValNameConstType(term.(*compilation_engine.ValNameConstantTerm))
	case compilation_engine.SubRoutineCallType:
		w.writeSubRoutineCall(term.(*compilation_engine.SubRoutineCall))
	case compilation_engine.ExpressionType:
		w.writeExpressionTerm(term.(*compilation_engine.ExpressionTerm))
	case compilation_engine.UnaryOpTermType:
		w.writeUnaryOpTerm(term.(*compilation_engine.UnaryOpTerm))
	default:
		chk.SE(fmt.Errorf("writeTerm:想定していないterm typeが来ました:%s", term.GetTermType()))
	}

	w.decNest()
	w.write("</term>")
}

// writeIntegerConstTerm .
func (w *XmlWriter) writeIntegerConstTerm(integerConstTerm *compilation_engine.IntegerConstTerm) {
	w.write(fmt.Sprintf("<integerConstant> %d </integerConstant>", integerConstTerm.Val))
}

// writeStringConstTerm .
func (w *XmlWriter) writeStringConstTerm(stringConstTerm *compilation_engine.StringConstTerm) {
	w.write(fmt.Sprintf("<stringConstant> %s </stringConstant>", stringConstTerm.Val))
}

// writeKeyWordConstTerm .
func (w *XmlWriter) writeKeyWordConstTerm(keyWordConstTerm *compilation_engine.KeyWordConstTerm) {
	w.write(fmt.Sprintf("<%s> %s </%s>", keyword, keyWordConstTerm.KeyWord, keyword))
}

// writeValNameConstType .
func (w *XmlWriter) writeValNameConstType(valNameConstantTerm *compilation_engine.ValNameConstantTerm) {
	w.write(fmt.Sprintf("<%s> %s </%s>", identifier, valNameConstantTerm.ValName, identifier))
}

// writeSubRoutineCall .
func (w *XmlWriter) writeSubRoutineCall(subRoutineCall *compilation_engine.SubRoutineCall) {

	// Class or ValName
	if subRoutineCall.ClassOrVarName != "" {
		w.write(getIdentifierXml(subRoutineCall.ClassOrVarName))
		w.write(getSymbolXml("."))
	}

	// subRoutineName
	w.write(getIdentifierXml(subRoutineCall.SubRoutineName))

	// (
	w.write(getSymbolXml("("))

	// expressionList
	w.writeExpressionList(subRoutineCall.ExpressionList)

	// )
	w.write(getSymbolXml(")"))
}

// writeExpressionTerm ()に包まれてるterm
func (w *XmlWriter) writeExpressionTerm(expressionTerm *compilation_engine.ExpressionTerm) {

	// (
	w.write(getSymbolXml("("))

	// expression
	w.writeExpression(expressionTerm.Expression)

	// )
	w.write(getSymbolXml(")"))
}

// writeUnaryOpTerm .
func (w *XmlWriter) writeUnaryOpTerm(unaryOpTerm *compilation_engine.UnaryOpTerm) {

	// unary
	w.write(getSymbolXml(string(unaryOpTerm.UnaryOp)))

	// term
	w.writeTerm(unaryOpTerm.Term)
}

// writeVariableType primitiveかどうかでkeywordかidentifierかが変わる
func (w *XmlWriter) writeVariableType(variableType compilation_engine.VariableType) {

	if variableType.IsPrimitive() {
		w.write(getKeywordXml(string(variableType)))
	} else {
		w.write(getIdentifierXml(string(variableType)))
	}

}

// write
func (writer *XmlWriter) write(value string) {

	// nest
	var nest string
	for i := 0; i < writer.nestDepth; i++ {
		nest += "\t"
	}

	_, err := writer.file.WriteString(fmt.Sprintf("%s%s\n", nest, value))
	chk.SE(err)
}

// createFile fileを作成する
func createFile(filePath string) *os.File {
	fp, err := os.Create(filePath)
	chk.SE(err)
	return fp
}

// nestを+1する
func (w *XmlWriter) incNest() {
	w.nestDepth++
}

// nestを-1する
func (w *XmlWriter) decNest() {
	w.nestDepth--
}

func getKeywordXml(v string) string {
	return fmt.Sprintf("<%s> %s </%s>", keyword, v, keyword)
}

func getIdentifierXml(v string) string {
	return fmt.Sprintf("<%s> %s </%s>", identifier, v, identifier)
}

func getSymbolXml(v string) string {
	return fmt.Sprintf("<%s> %s </%s>", symbol, v, symbol)
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
