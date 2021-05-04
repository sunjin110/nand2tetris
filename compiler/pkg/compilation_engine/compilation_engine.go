package compilation_engine

import (
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/tokenizer"
	"log"
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
		panic("classではありませんでした")
	}

	// クラス名を取得する
	c.nextToken()
	className := c.getToken()

	// {
	c.nextToken()
	if c.getToken() != "{" {
		panic("構文Error: classで「{」がありません")
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
			panic(";が足りないです")
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
			panic("subRoutineで{がありませんでした")
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
			panic("SubRoutineの「}」がありません")
		}

		// 次もsubroutineかどうかを確認
		// 違うなら、roopから外れる
		c.nextToken()
		log.Println("次のやつ", c.getToken())
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
		panic("引数の(がありませんでした")
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
			panic(";が足りないです")
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
		switch statementType {
		case LetStatementPrefix:
			letStatement := c.compileLet()
			statementList = append(statementList, letStatement)
		case IfStatementPrefix:
			ifStatement := c.compileIf()
			statementList = append(statementList, ifStatement)

			// TODO ここに到達しない
			log.Println("if statmenet is ", jsonutil.Marshal(ifStatement))

		case WhileStatementPrefix:

		case DoStatementPrefix:
			doStatement := c.compileDo()
			statementList = append(statementList, doStatement)
		case ReturnStatementPrefix:
			// ここで必ずreturnする
			returnStatement := c.compileReturn()
			statementList = append(statementList, returnStatement)
			return statementList
		}

		// 次に進める
		// この段階で、次のprefixがstatementでない場合は、Error(Returnがない)
		c.nextToken()
		log.Println("hoge is ", c.getToken())
		if !IsStatementPrefixToken(c.getToken()) {
			panic("Statementに「return」が指定されていません")
		}

	}

}

// compileDo do文をコンパイルする
func (c *CompilationEngine) compileDo() *DoStatement {

	if c.getToken() != string(DoStatementPrefix) {
		panic("doのstatementではありません")
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
		panic("SubRoutineCallで「(」がありません")
	}

	// expresstionList
	expressionList := c.compileExpressionList()

	// ) check
	if c.getToken() != ")" {
		panic("SubRoutineCallで「)」がありません")
	}

	// ; check
	c.nextToken()
	if c.getToken() != ";" {
		panic("SubRoutineCallで「;」がありません")
	}

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
		panic("letのstatementではありません")
	}

	// name
	c.nextToken()
	destVarName := c.getToken()

	// arrayかどうかを判定する
	var arrayExpression *Expression
	c.nextToken()
	if c.getToken() == "[" {
		arrayExpression = c.compileExpression()

		if c.getToken() != "]" {
			panic("let式のarrayに「]」がありませんでした")
		}
	}

	// = check
	if c.getToken() != string(EqlOp) {
		panic("let式に=がありません")
	}

	// 代入する式を取得する
	c.nextToken()
	expression := c.compileExpression()

	// 「;」check
	if c.getToken() != ";" {
		panic("Let式の最後に「;」がありません")
	}

	return &LetStatement{
		DestVarName:     destVarName,
		ArrayExpression: arrayExpression,
		Expression:      expression,
	}
}

// compileWhile while文をコンパイルする
func compileWhile() {

}

// compileReturn return文をコンパイルする
func (c *CompilationEngine) compileReturn() *ReturnStatement {

	if c.getToken() != string(ReturnStatementPrefix) {
		panic("returnのstatementではありません")
	}

	returnExpression := c.compileExpression()

	if c.getToken() != ";" {
		panic("returnに「;」が含まれていませんでした")
	}

	return &ReturnStatement{
		ReturnExpression: returnExpression,
	}
}

// compileIf if文をコンパイルする。else文を伴う可能性がある
func (c *CompilationEngine) compileIf() *IfStatement {

	if c.getToken() != string(IfStatementPrefix) {
		panic("ifのsatementではありません")
	}

	// ( check
	c.nextToken()
	if c.getToken() != "(" {
		panic("ifのstatementに「(」がありません")
	}

	// expression
	c.nextToken()
	conditionalExpression := c.compileExpression()

	// ) check
	if c.getToken() != ")" {
		panic("ifのstatementに「)」がありません")
	}

	// { check
	c.nextToken()
	if c.getToken() != "{" {
		panic("ifのstatementに「{」がありません")
	}

	// statement
	c.nextToken()
	statementList := c.compileStatements()

	// } check
	if c.getToken() != "}" {
		panic("ifのstatmenetに「}」がありません")
	}

	// elseがあるかどうか?
	c.nextToken()
	var elseStatementList []Statement
	if c.getToken() == "else" {
		elseStatementList = c.compileStatements()
	}

	return &IfStatement{
		ConditionalExpression: conditionalExpression,
		StatementList:         statementList,
		ElseStatementList:     elseStatementList,
	}
}

// compileExpression 式をコンパイルする
func (c *CompilationEngine) compileExpression() *Expression {

	// TODO
	c.nextToken()

	return &Expression{}
}

// compileExpressionList コンマで分離された式のリスト(空白の可能性もある)をコンパイルする
func (c *CompilationEngine) compileExpressionList() []*Expression {

	// TODO
	c.nextToken()

	return nil
}

// compileTerm termをコンパイルする
func compileTerm() {

}
