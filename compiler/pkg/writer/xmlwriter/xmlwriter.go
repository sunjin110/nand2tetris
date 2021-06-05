package xmlwriter

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/common/fileutil"
	"compiler/pkg/compilation_engine"
	"fmt"
	"os"
)

const (
	keyword    = "keyword"
	identifier = "identifier"
	symbol     = "symbol"
)

// XMLWriter .
type XMLWriter struct {
	file      *os.File
	class     *compilation_engine.Class
	nestDepth int // ネストの深さ(tab)
}

// New XMLWriterを作成する
func New(filePath string, class *compilation_engine.Class) *XMLWriter {
	return &XMLWriter{
		file:      fileutil.CreateFile(filePath),
		class:     class,
		nestDepth: 0,
	}
}

// WriteParser パーサで解析した内容を書き出す
func (writer *XMLWriter) WriteParser() {
	writer.writeClass()
}

// writeClass Classからxmlのファイルを作成する
func (writer *XMLWriter) writeClass() error {

	writer.write("<class>")

	// class name
	writer.incNest()

	// class init
	writer.write(getKeywordXML("class"))
	writer.write(getIdentifierXML(writer.class.ClassName))
	writer.write(getSymbolXML("{"))

	// class var dec
	writer.writeClassVarDec(writer.class.ClassVarDecList)

	// subroutine dec
	writer.writeSubroutineDec(writer.class.SubRoutineDecList)

	// class defer
	writer.write(getSymbolXML("}"))

	writer.decNest()

	writer.write("</class>")

	return nil
}

// writeClassVarDec .
func (writer *XMLWriter) writeClassVarDec(classVarDecList []*compilation_engine.ClassVarDec) {
	for _, classVarDec := range classVarDecList {

		writer.write("<classVarDec>")
		writer.incNest()

		// var kind
		writer.write(getKeywordXML(string(classVarDec.VarKind)))

		// type
		// writer.write(getKeywordXML(string(classVarDec.VarType)))
		writer.writeVariableType(classVarDec.VarType)

		// names
		for i, varName := range classVarDec.VarNameList {
			writer.write(getIdentifierXML(varName))

			// 最後でない場合は「,」を追加
			if len(classVarDec.VarNameList) != i+1 {
				writer.write(getSymbolXML(","))
			}
		}

		// ;
		writer.write(getSymbolXML(";"))

		writer.decNest()
		writer.write("</classVarDec>")
	}
}

// writeSubroutineDec .
func (writer *XMLWriter) writeSubroutineDec(subRoutineDecList []*compilation_engine.SubRoutineDec) {

	for _, subRoutineDec := range subRoutineDecList {

		writer.write("<subroutineDec>")
		writer.incNest()

		writer.write(getKeywordXML(string(subRoutineDec.RoutineKind)))

		writer.writeVariableType(subRoutineDec.ReturnType)

		writer.write(getIdentifierXML(subRoutineDec.SubRoutineName))

		// (
		writer.write(getSymbolXML("("))

		// parameterList
		writer.write("<parameterList>")
		writer.incNest()
		for i, parameter := range subRoutineDec.ParameterList {

			writer.writeVariableType(parameter.ParamType)
			writer.write(getIdentifierXML(parameter.ParamName))

			if len(subRoutineDec.ParameterList) != i+1 {
				writer.write(getSymbolXML(","))
			}

		}
		writer.decNest()
		writer.write("</parameterList>")

		// )
		writer.write(getSymbolXML(")"))

		// subroutineBody
		writer.writeSubroutineBody(subRoutineDec.SubRoutineBody)

		writer.decNest()
		writer.write("</subroutineDec>")

	}

}

// writeSubroutineBody .
func (writer *XMLWriter) writeSubroutineBody(subRoutineBody *compilation_engine.SubRoutineBody) {

	defer func() {
		// }
		writer.write(getSymbolXML("}"))

		writer.decNest()
		writer.write("</subroutineBody>")
	}()

	writer.write("<subroutineBody>")
	writer.incNest()

	// {
	writer.write(getSymbolXML("{"))

	// check
	if subRoutineBody == nil {
		return
	}

	// var dec
	writer.writeVarDec(subRoutineBody.VarDecList)

	// statement
	writer.writeStatement(subRoutineBody.StatementList)

}

// writeVarDec .
func (writer *XMLWriter) writeVarDec(varDecList []*compilation_engine.VarDec) {

	for _, varDec := range varDecList {

		writer.write("<varDec>")
		writer.incNest()

		writer.write(getKeywordXML("var"))

		writer.writeVariableType(varDec.Type)

		for i, varName := range varDec.NameList {
			writer.write(getIdentifierXML(varName))
			if len(varDec.NameList) != i+1 {
				writer.write(getSymbolXML(","))
			}
		}

		// ;
		writer.write(getSymbolXML(";"))

		writer.decNest()
		writer.write("</varDec>")

	}
}

// writeStatement .
func (writer *XMLWriter) writeStatement(statementList []compilation_engine.Statement) {

	writer.write("<statements>")
	writer.incNest()

	for _, statement := range statementList {

		switch statement.GetStatementType() {
		case compilation_engine.LetStatementPrefix:
			writer.writeLetStatement(statement.(*compilation_engine.LetStatement))
		case compilation_engine.IfStatementPrefix:
			writer.writeIfStatement(statement.(*compilation_engine.IfStatement))
		case compilation_engine.WhileStatementPrefix:
			writer.writeWhileStatement(statement.(*compilation_engine.WhileStatement))
		case compilation_engine.DoStatementPrefix:
			writer.writeDoStatement(statement.(*compilation_engine.DoStatement))
		case compilation_engine.ReturnStatementPrefix:
			writer.writeReturnStatement(statement.(*compilation_engine.ReturnStatement))
		default:
			chk.SE(fmt.Errorf("writeStatement: 宣言していないstatementが渡されました:%s", statement.GetStatementType()))
		}
	}

	writer.decNest()
	writer.write("</statements>")
}

// writeLetStatement .
func (writer *XMLWriter) writeLetStatement(letStatement *compilation_engine.LetStatement) {

	writer.write("<letStatement>")
	writer.incNest()

	writer.write(getKeywordXML("let"))
	writer.write(getIdentifierXML(letStatement.DestVarName))

	// array
	if letStatement.ArrayExpression != nil {
		writer.write(getSymbolXML("["))
		writer.writeExpression(letStatement.ArrayExpression)
		writer.write(getSymbolXML("]"))
	}

	// =
	writer.write(getSymbolXML("="))

	// expression
	writer.writeExpression(letStatement.Expression)

	writer.write(getSymbolXML(";"))

	writer.decNest()
	writer.write("</letStatement>")

}

// writeIfStatement .
func (writer *XMLWriter) writeIfStatement(ifStatement *compilation_engine.IfStatement) {

	writer.write("<ifStatement>")
	writer.incNest()

	writer.write(getKeywordXML("if"))

	// (
	writer.write(getSymbolXML("("))

	if ifStatement.ConditionalExpression != nil {
		writer.writeExpression(ifStatement.ConditionalExpression)
	}

	// )
	writer.write(getSymbolXML(")"))

	// {
	writer.write(getSymbolXML("{"))

	writer.writeStatement(ifStatement.StatementList)

	// }
	writer.write(getSymbolXML("}"))

	// else
	if len(ifStatement.ElseStatementList) > 0 {
		writer.write(getKeywordXML("else"))
		writer.write(getSymbolXML("{"))
		writer.writeStatement(ifStatement.ElseStatementList)
		writer.write(getSymbolXML("}"))
	}

	writer.decNest()
	writer.write("</ifStatement>")
}

// writeWhileStatement .
func (writer *XMLWriter) writeWhileStatement(whileStatement *compilation_engine.WhileStatement) {

	writer.write("<whileStatement>")
	writer.incNest()

	writer.write(getKeywordXML("while"))

	// (
	writer.write(getSymbolXML("("))

	if whileStatement.ConditionalExpression != nil {
		writer.writeExpression(whileStatement.ConditionalExpression)
	}

	// )
	writer.write(getSymbolXML(")"))

	// {
	writer.write(getSymbolXML("{"))

	// statement
	writer.writeStatement(whileStatement.StatementList)

	// }
	writer.write(getSymbolXML("}"))

	writer.decNest()
	writer.write("</whileStatement>")

}

// writeDoStatement .
func (writer *XMLWriter) writeDoStatement(doStatement *compilation_engine.DoStatement) {

	writer.write("<doStatement>")
	writer.incNest()

	writer.write(getKeywordXML("do"))

	writer.writeSubRoutineCall(doStatement.SubroutineCall)

	writer.write(getSymbolXML(";"))

	writer.decNest()
	writer.write("</doStatement>")
}

// writeReturnStatement .
func (writer *XMLWriter) writeReturnStatement(returnStatement *compilation_engine.ReturnStatement) {

	writer.write("<returnStatement>")
	writer.incNest()

	writer.write(getKeywordXML("return"))

	if returnStatement.ReturnExpression != nil {
		writer.writeExpression(returnStatement.ReturnExpression)
	}

	writer.write(getSymbolXML(";"))

	writer.decNest()
	writer.write("</returnStatement>")
}

// writeExpressionList .
func (writer *XMLWriter) writeExpressionList(expressionList []*compilation_engine.Expression) {
	writer.write("<expressionList>")
	writer.incNest()

	for i, expression := range expressionList {
		writer.writeExpression(expression)

		if len(expressionList) != i+1 {
			writer.write(getSymbolXML(","))
		}
	}

	writer.decNest()
	writer.write("</expressionList>")
}

// writeExpression .
func (writer *XMLWriter) writeExpression(expression *compilation_engine.Expression) {

	writer.write("<expression>")
	writer.incNest()

	// term
	writer.writeTerm(expression.InitTerm)

	// op term
	for _, opTerm := range expression.OpTermList {

		// operation
		writer.write(getSymbolXML(string(opTerm.Operation)))

		// term
		writer.writeTerm(opTerm.OpTerm)
	}

	writer.decNest()
	writer.write("</expression>")
}

// writeTerm
func (writer *XMLWriter) writeTerm(term compilation_engine.Term) {

	writer.write("<term>")
	writer.incNest()

	switch term.GetTermType() {

	case compilation_engine.IntegerConstType:
		writer.writeIntegerConstTerm(term.(*compilation_engine.IntegerConstTerm))
	case compilation_engine.StringConstType:
		writer.writeStringConstTerm(term.(*compilation_engine.StringConstTerm))
	case compilation_engine.KeyWordConstType:
		writer.writeKeyWordConstTerm(term.(*compilation_engine.KeyWordConstTerm))
	case compilation_engine.ValNameConstType:
		writer.writeValNameConstType(term.(*compilation_engine.ValNameConstantTerm))
	case compilation_engine.SubRoutineCallType:
		writer.writeSubRoutineCall(term.(*compilation_engine.SubRoutineCall))
	case compilation_engine.ExpressionType:
		writer.writeExpressionTerm(term.(*compilation_engine.ExpressionTerm))
	case compilation_engine.UnaryOpTermType:
		writer.writeUnaryOpTerm(term.(*compilation_engine.UnaryOpTerm))
	default:
		chk.SE(fmt.Errorf("writeTerm:想定していないterm typeが来ました:%s", term.GetTermType()))
	}

	writer.decNest()
	writer.write("</term>")
}

// writeIntegerConstTerm .
func (writer *XMLWriter) writeIntegerConstTerm(integerConstTerm *compilation_engine.IntegerConstTerm) {
	writer.write(fmt.Sprintf("<integerConstant> %d </integerConstant>", integerConstTerm.Val))
}

// writeStringConstTerm .
func (writer *XMLWriter) writeStringConstTerm(stringConstTerm *compilation_engine.StringConstTerm) {
	writer.write(fmt.Sprintf("<stringConstant> %s </stringConstant>", stringConstTerm.Val))
}

// writeKeyWordConstTerm .
func (writer *XMLWriter) writeKeyWordConstTerm(keyWordConstTerm *compilation_engine.KeyWordConstTerm) {
	writer.write(fmt.Sprintf("<%s> %s </%s>", keyword, keyWordConstTerm.KeyWord, keyword))
}

// writeValNameConstType .
func (writer *XMLWriter) writeValNameConstType(valNameConstantTerm *compilation_engine.ValNameConstantTerm) {

	// identifier
	writer.write(fmt.Sprintf("<%s> %s </%s>", identifier, valNameConstantTerm.ValName, identifier))

	// ある場合はなんか上手いことする
	if valNameConstantTerm.ArrayExpression != nil {
		writer.write(getSymbolXML("["))
		writer.writeExpression(valNameConstantTerm.ArrayExpression)
		writer.write(getSymbolXML("]"))
	}
}

// writeSubRoutineCall .
func (writer *XMLWriter) writeSubRoutineCall(subRoutineCall *compilation_engine.SubRoutineCall) {

	// Class or ValName
	if subRoutineCall.ClassOrVarName != "" {
		writer.write(getIdentifierXML(subRoutineCall.ClassOrVarName))
		writer.write(getSymbolXML("."))
	}

	// subRoutineName
	writer.write(getIdentifierXML(subRoutineCall.SubRoutineName))

	// (
	writer.write(getSymbolXML("("))

	// expressionList
	writer.writeExpressionList(subRoutineCall.ExpressionList)

	// )
	writer.write(getSymbolXML(")"))
}

// writeExpressionTerm ()に包まれてるterm
func (writer *XMLWriter) writeExpressionTerm(expressionTerm *compilation_engine.ExpressionTerm) {

	// (
	writer.write(getSymbolXML("("))

	// expression
	writer.writeExpression(expressionTerm.Expression)

	// )
	writer.write(getSymbolXML(")"))
}

// writeUnaryOpTerm .
func (writer *XMLWriter) writeUnaryOpTerm(unaryOpTerm *compilation_engine.UnaryOpTerm) {

	// unary
	writer.write(getSymbolXML(string(unaryOpTerm.UnaryOp)))

	// term
	writer.writeTerm(unaryOpTerm.Term)
}

// writeVariableType primitiveかどうかでkeywordかidentifierかが変わる
func (writer *XMLWriter) writeVariableType(variableType compilation_engine.VariableType) {

	if variableType.IsPrimitive() {
		writer.write(getKeywordXML(string(variableType)))
	} else {
		writer.write(getIdentifierXML(string(variableType)))
	}

}

// write
func (writer *XMLWriter) write(value string) {

	// nest
	var nest string
	for i := 0; i < writer.nestDepth; i++ {
		// nest += "\t"
		nest += "  "
	}

	_, err := writer.file.WriteString(fmt.Sprintf("%s%s\n", nest, value))
	chk.SE(err)
}

// nestを+1する
func (writer *XMLWriter) incNest() {
	writer.nestDepth++
}

// nestを-1する
func (writer *XMLWriter) decNest() {
	writer.nestDepth--
}

func getKeywordXML(v string) string {
	return fmt.Sprintf("<%s> %s </%s>", keyword, v, keyword)
}

func getIdentifierXML(v string) string {
	return fmt.Sprintf("<%s> %s </%s>", identifier, v, identifier)
}

func getSymbolXML(v string) string {
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
