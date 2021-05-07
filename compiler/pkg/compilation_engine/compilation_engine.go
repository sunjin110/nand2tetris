package compilation_engine

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/tokenizer"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// tokenizerから入力を受け取り、構文解析された構造を出力ファイルに出力する

// CompilationEngine .
type CompilationEngine struct {
	Tknz  *tokenizer.Tokenizer
	Class *Class // クラス
}

// New CompilationEngineを作成する
func New(t *tokenizer.Tokenizer) *CompilationEngine {
	return &CompilationEngine{
		Tknz: t,
	}
}

// SyntaxError 構文Error
func (c *CompilationEngine) SyntaxError(errMsg string) {
	chk.SE(fmt.Errorf("SyntaxError: \n\tfile:%s:%d\n\terr-msg:%s\n\tnow-token: %s", c.Tknz.FileName, c.Tknz.LineNum, errMsg, c.getToken()))
}

// nextToken 次のtokenに移動
func (c *CompilationEngine) nextToken() bool {
	return c.Tknz.NextToken()
}

// getToken 次のtokenを取得する
func (c *CompilationEngine) getToken() string {
	return c.Tknz.Token
}

// Start コンパイル開始
func (c *CompilationEngine) Start() {

	// compileCalssを先に必ず呼ぶ
	c.compileClass()
}

// compileClass クラスをコンパイルする
func (c *CompilationEngine) compileClass() {

	// 一番頭、Classかどうかを判定する
	c.nextToken()
	if tokenizer.GetKeyWord(c.getToken()) != tokenizer.KeyWordClass {
		c.SyntaxError("classではありませんでした")
	}

	// クラス名を取得する
	c.nextToken()
	className := c.getToken()

	// {
	c.nextToken()
	if c.getToken() != "{" {
		c.SyntaxError("構文Error: classで「{」がありません")
	}

	c.nextToken()

	// ClassVarDecをどうにかする
	classVarDecList := c.compileClassVarDec()

	// SubRoutineList
	subRoutineDecList := c.compileSubroutine()

	// TODO } のチェック

	class := &Class{
		ClassName:         className,
		ClassVarDecList:   classVarDecList,
		SubRoutineDecList: subRoutineDecList,
	}

	c.Class = class
}

// compileClassVarDec スタティック宣言またはフィールド宣言をコンパイルする
func (c *CompilationEngine) compileClassVarDec() []*ClassVarDec {

	t := c.getToken()

	// classの宣言があるかどうかを確認

	// static or fieldという文字がない場合は、varDecが存在しないので、nilで返す
	if !IsClassVarDecPrefixToken(t) {
		return nil
	}
	var classVarDecList []*ClassVarDec
	for {
		// static or field
		varKind := c.getToken()

		// int or char or boolean or className
		c.nextToken()
		varType := c.getToken()

		// 変数名 (同時宣言しているものがあるので、それもチェックする)
		var varNameList []string
		for {
			c.nextToken()
			name := c.getToken()
			varNameList = append(varNameList, name)

			c.nextToken()
			if c.getToken() == "," {
				continue
			}

			break
		}

		// make
		classVarDec := &ClassVarDec{
			VarKind:     VariableKind(varKind),
			VarType:     VariableType(varType),
			VarNameList: varNameList,
		}

		classVarDecList = append(classVarDecList, classVarDec)

		// check
		if c.getToken() != ";" {
			c.SyntaxError(";が足りないです")
		}

		// next チェック
		// すでに他のtokenに移っている場合は、skip
		c.nextToken()
		if !IsClassVarDecPrefixToken(c.getToken()) {
			break
		}
	}

	return classVarDecList
}

// compileSubroutine メソッド、ファンクション、コンストラクタをコンパイルする
func (c *CompilationEngine) compileSubroutine() []*SubRoutineDec {

	// check
	t := c.getToken()

	// サブルーチンかどうか
	if !IsSubRoutineDecPrefixToken(t) {
		return nil
	}

	var subRoutineDecList []*SubRoutineDec
	for {

		// function or constructor or method
		subRoutineKind := c.getToken()

		// void or any type
		c.nextToken()
		returnType := c.getToken()

		// sub routine name
		c.nextToken()
		subRoutineName := c.getToken()

		// parameter
		parameterList := c.compileParameterList()

		// check {
		c.nextToken()
		if c.getToken() != "{" {
			c.SyntaxError("subRoutineで{がありませんでした")
		}

		// subRoutineBody
		// 先に型を宣言して、処理する形を遵守する
		c.nextToken()
		varDecList := c.compileVarDec()

		// statement
		statementList := c.compileStatements()

		subRoutineDec := &SubRoutineDec{
			RoutineKind:    SubRoutineKind(subRoutineKind),
			ReturnType:     VariableType(returnType),
			SubRoutineName: subRoutineName,
			ParameterList:  parameterList,
			SubRoutineBody: &SubRoutineBody{
				VarDecList:    varDecList,
				StatementList: statementList,
			},
		}

		subRoutineDecList = append(subRoutineDecList, subRoutineDec)

		// } check
		c.nextToken()
		if c.getToken() != "}" {
			c.SyntaxError("SubRoutineの「}」がありません")
		}

		// 次もsubroutineかどうかを確認
		// 違うなら、roopから外れる
		c.nextToken()
		if !IsSubRoutineDecPrefixToken(c.getToken()) {
			break
		}

	}

	return subRoutineDecList
}

// compileParameterList パラメータのリスト(空の可能性もある)をコンパイルする。カッコ"()"は含まない
func (c *CompilationEngine) compileParameterList() []*Parameter {

	c.nextToken()
	if c.getToken() != "(" {
		c.SyntaxError("引数の(がありませんでした")
	}

	var parameterList []*Parameter
	c.nextToken()
	for {

		// もし)がきたら終了
		if c.getToken() == ")" {
			break
		}

		// 引数の型
		paramType := c.getToken()

		// 引数名
		c.nextToken()
		paramName := c.getToken()

		// ではない場合は、type
		parameter := &Parameter{
			ParamType: VariableType(paramType),
			ParamName: paramName,
		}

		parameterList = append(parameterList, parameter)

		c.nextToken()
		if c.getToken() == "," {
			c.nextToken()
			continue
		}
	}

	return parameterList
}

// compileVarDec var宣言をコンパイルする
func (c *CompilationEngine) compileVarDec() []*VarDec {

	// c.nextToken()
	if !IsVarDecPrefixToken(c.getToken()) {
		// var宣言がありませんでした
		return nil
	}

	var varDecList []*VarDec
	for {

		// type
		c.nextToken()
		varType := c.getToken()

		// 変数名 (同時宣言をしているものがあるので、それもチェックする)
		var varNameList []string
		for {
			c.nextToken()
			name := c.getToken()
			varNameList = append(varNameList, name)

			c.nextToken()
			if c.getToken() == "," {
				continue
			}
			break
		}

		// make
		varDec := &VarDec{
			Type:     VariableType(varType),
			NameList: varNameList,
		}

		varDecList = append(varDecList, varDec)

		// check
		if c.getToken() != ";" {
			c.SyntaxError(";が足りないです")
		}

		// next チェック
		// すでに他のtokenに移っている場合は、skip
		c.nextToken()
		if !IsVarDecPrefixToken(c.getToken()) {
			break
		}
	}

	return varDecList
}

// compileStatements 一連の文をコンパイルする。波括弧"{}"は含まない
func (c *CompilationEngine) compileStatements() []Statement {

	if !IsStatementPrefixToken(c.getToken()) {
		// statementの宣言がありませんでした
		return nil
	}

	var statementList []Statement

	for {

		statementType := c.getToken()
		log.Println("statement type is ", statementType)
		switch statementType {
		case LetStatementPrefix:
			letStatement := c.compileLet()
			statementList = append(statementList, letStatement)
		case IfStatementPrefix:
			ifStatement := c.compileIf()
			statementList = append(statementList, ifStatement)
		case WhileStatementPrefix:
			whileStatement := c.compileWhile()
			statementList = append(statementList, whileStatement)
		case DoStatementPrefix:
			doStatement := c.compileDo()
			statementList = append(statementList, doStatement)
		case ReturnStatementPrefix:
			// ここで必ずreturnする
			returnStatement := c.compileReturn()
			statementList = append(statementList, returnStatement)
			return statementList
		}

		// この段階で、次のprefixがstatementでない場合は、終了
		log.Println("compile statementsの最後", c.getToken())
		if !IsStatementPrefixToken(c.getToken()) {
			break
		}

	}

	return statementList
}

// compileDo do文をコンパイルする
func (c *CompilationEngine) compileDo() *DoStatement {

	if c.getToken() != string(DoStatementPrefix) {
		c.SyntaxError("doのstatementではありません")
	}

	c.nextToken()
	tmp := c.getToken()

	// ClassName or varNameがあるかどうか
	var classOrVarName string
	var subRoutineName string
	c.nextToken()
	if c.getToken() == "." {
		classOrVarName = tmp

		// soubRoutineName
		c.nextToken()
		subRoutineName = c.getToken()

		c.nextToken()
	} else {
		subRoutineName = tmp
	}

	// ( check
	if c.getToken() != "(" {
		c.SyntaxError("SubRoutineCallで「(」がありません")
	}

	// expresstionList
	c.nextToken()
	expressionList := c.compileExpressionList()

	// ) check
	if c.getToken() != ")" {
		c.SyntaxError(fmt.Sprintf("SubRoutineCallで「)」がありません:%s", c.getToken()))
	}

	// ; check
	c.nextToken()
	if c.getToken() != ";" {
		c.SyntaxError("SubRoutineCallで「;」がありません")
	}

	// ;をスキップ
	c.nextToken()

	return &DoStatement{
		SubroutineCall: &SubRoutineCall{
			ClassOrVarName: classOrVarName,
			SubRoutineName: subRoutineName,
			ExpressionList: expressionList,
		},
	}
}

// compileLet let文をコンパイルする
func (c *CompilationEngine) compileLet() *LetStatement {

	if c.getToken() != string(LetStatementPrefix) {
		c.SyntaxError("letのstatementではありません")
	}

	// name
	c.nextToken()
	destVarName := c.getToken()

	// arrayかどうかを判定する
	var arrayExpression *Expression
	c.nextToken()
	if c.getToken() == "[" {
		c.nextToken()
		arrayExpression = c.compileExpression()

		if c.getToken() != "]" {
			c.SyntaxError("let式のarrayに「]」がありませんでした")
		}
	}

	// = check
	if c.getToken() != string(EqlOp) {
		c.SyntaxError("let式に=がありません")
	}

	// 代入する式を取得する
	c.nextToken()
	expression := c.compileExpression()

	// 「;」check
	if c.getToken() != ";" {
		c.SyntaxError("Let式の最後に「;」がありません")
	}

	// ;をスキップ
	c.nextToken()

	return &LetStatement{
		DestVarName:     destVarName,
		ArrayExpression: arrayExpression,
		Expression:      expression,
	}
}

// compileWhile while文をコンパイルする
func (c *CompilationEngine) compileWhile() *WhileStatement {

	if c.getToken() != string(WhileStatementPrefix) {
		c.SyntaxError("shileのstatementではありません")
	}

	c.nextToken()
	if c.getToken() != "(" {
		c.SyntaxError("while statementで「(」がありません")
	}

	c.nextToken()
	// coordinationExpressionを取得
	coordinationExpression := c.compileExpression()

	// check )
	if c.getToken() != ")" {
		c.SyntaxError("while statmentで「)」がありません")
	}

	// check {
	c.nextToken()
	if c.getToken() != "{" {
		c.SyntaxError("while statmentで「{」がありません")
	}

	// statementListを習得する
	c.nextToken()
	statementList := c.compileStatements()

	// check }
	if c.getToken() != "}" {
		c.SyntaxError("while statementで「}」がありません")
	}

	c.nextToken()

	return &WhileStatement{
		ConditionalExpression: coordinationExpression,
		StatementList:         statementList,
	}
}

// compileReturn return文をコンパイルする
func (c *CompilationEngine) compileReturn() *ReturnStatement {

	if c.getToken() != string(ReturnStatementPrefix) {
		c.SyntaxError("returnのstatementではありません")
	}

	c.nextToken()
	returnExpression := c.compileExpression()

	if c.getToken() != ";" {
		c.SyntaxError("returnに「;」が含まれていませんでした")
	}

	return &ReturnStatement{
		ReturnExpression: returnExpression,
	}
}

// compileIf if文をコンパイルする。else文を伴う可能性がある
func (c *CompilationEngine) compileIf() *IfStatement {

	if c.getToken() != string(IfStatementPrefix) {
		c.SyntaxError("ifのsatementではありません")
	}

	// ( check
	c.nextToken()
	if c.getToken() != "(" {
		c.SyntaxError("ifのstatementに「(」がありません")
	}

	// expression
	c.nextToken()
	conditionalExpression := c.compileExpression()

	// ) check
	if c.getToken() != ")" {
		c.SyntaxError("ifのstatementに「)」がありません")
	}

	// { check
	c.nextToken()
	if c.getToken() != "{" {
		c.SyntaxError("ifのstatementに「{」がありません")
	}

	// statement
	c.nextToken()
	statementList := c.compileStatements()

	// } check
	if c.getToken() != "}" {
		c.SyntaxError("ifのstatmenetに「}」がありません")
	}

	// elseがあるかどうか?
	c.nextToken()
	var elseStatementList []Statement

	if c.getToken() == "else" {
		c.nextToken()

		// { チェック
		if c.getToken() != "{" {
			c.SyntaxError("elseの次の「{」がありません")
		}

		c.nextToken()
		elseStatementList = c.compileStatements()

		// } check
		if c.getToken() != "}" {
			c.SyntaxError("elseの最後の「}」がありません")
		}

		c.nextToken()
	}

	return &IfStatement{
		ConditionalExpression: conditionalExpression,
		StatementList:         statementList,
		ElseStatementList:     elseStatementList,
	}
}

// compileExpression 式をコンパイルする
func (c *CompilationEngine) compileExpression() *Expression {

	// Termかどうかを判定する
	if !IsTermPrefixToken(c.getToken()) {
		// 式ではない場合は、スキップする
		return nil
	}

	initTerm := c.compileTerm()

	var opTermList []*OpTerm
	for {

		// Opがあることを確認する
		var op Op
		switch Op(c.getToken()[0]) {
		case PlusOp, MinusOp, AsteriskOp,
			SlashOp, AndOp, PipeOp, LessThanOp,
			GreaterThanOp, EqlOp:

			op = Op(c.getToken()[0])
		default:
			return &Expression{
				InitTerm:   initTerm,
				OpTermList: opTermList,
			}
		}

		// term
		c.nextToken()
		term := c.compileTerm()

		opTerm := &OpTerm{
			Operation: op,
			OpTerm:    term,
		}

		opTermList = append(opTermList, opTerm)

	}

}

// compileExpressionList コンマで分離された式のリスト(空白の可能性もある)をコンパイルする
func (c *CompilationEngine) compileExpressionList() []*Expression {

	// expressionListを追加する必要がない場合は早期returnする
	if !IsExpressionListPrefixToken(c.getToken()) {
		return nil
	}

	var expressionList []*Expression

	for {

		expression := c.compileExpression()
		if expression != nil {
			expressionList = append(expressionList, expression)
		}

		// もし「,」でない場合は終了
		if c.getToken() != "," {
			break
		}
		c.nextToken()
	}

	return expressionList
}

// compileTerm termをコンパイルする
// TODO nextのことをまだ、unaryとか()のやつは考えられていない
// そもそもまだ未完成
func (c *CompilationEngine) compileTerm() Term {

	token := c.getToken()

	// 数字に変換できる場合は、数字const
	if i, err := strconv.Atoi(token); err == nil {
		c.nextToken()
		return &IntegerConstTerm{
			Val: i,
		}
	}

	// 文字列
	if strings.HasPrefix(token, "\"") && strings.HasSuffix(token, "\"") {
		val := c.Tknz.GetStringVal()
		c.nextToken()
		return &StringConstTerm{
			Val: val,
		}
	}

	// unary付きの場合
	if rune(token[0]) == '-' || rune(token[0]) == '~' {

		unaryOp := UnaryOp(token[0])

		c.nextToken()
		term := c.compileTerm()

		return &UnaryOpTerm{
			UnaryOp: unaryOp,
			Term:    term,
		}
	}

	// keyword const
	if token == string(TrueKeyword) ||
		token == string(FalseKeyword) ||
		token == string(NullKeyword) ||
		token == string(ThisKeyword) {

		c.nextToken()

		return &KeyWordConstTerm{
			KeyWord: KeywordConstant(token),
		}
	}

	// 先頭が(の場合はなんかうまいことする
	if rune(token[0]) == '(' {

		c.nextToken()
		expression := c.compileExpression()

		// TODO )を調べる

		return &ExpressionTerm{
			Expression: expression,
		}
	}

	// 変数
	if isVariableToken(token) {

		valName := token

		// もし「[」がある場合は式を取得する
		var arrayExpression *Expression
		c.nextToken()
		if c.getToken() == "[" {

			c.nextToken()
			arrayExpression = c.compileExpression()
		}

		return &ValNameConstantTerm{
			ValName:         valName,
			ArrayExpression: arrayExpression,
		}
	}

	// それ以外はsubroutine
	// subRoutineCall :=
	// if subRoutineCall != nil {
	// return subRoutineCall
	// }

	// TODO SubRoutineCallとそれ以外
	return nil
}
