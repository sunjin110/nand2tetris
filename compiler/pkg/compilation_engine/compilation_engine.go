package compilation_engine

import (
	"compiler/pkg/tokenizer"
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

	c.nextToken()
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

	c.nextToken()
	if !IsStatementPrefixToken(c.getToken()) {
		// statementの宣言がありませんでした
		return nil
	}

	var statementList []Statement

	for {

		statementType := c.getToken()
		switch statementType {
		case LetStatementPrefix:

		case IfStatementPrefix:

		case WhileStatementPrefix:

		case DoStatementPrefix:

		case ReturnStatementPrefix:
			// ここで必ずreturnする

			return statementList
		}

	}

}

// compileDo do文をコンパイルする
func compileDo() {

}

// compileLet let文をコンパイルする
func compileLet() {

}

// compileWhile while文をコンパイルする
func compileWhile() {

}

// compileReturn return文をコンパイルする
func compileReturn() {

}

// compileIf if文をコンパイルする。else文を伴う可能性がある
func compileIf() {

}

// compileExpression 式をコンパイルする
func compileExpression() {

}

// compileTerm termをコンパイルする
func compileTerm() {

}

// compileExpressionList コンマで分離された式のリスト(空白の可能性もある)をコンパイルする
func compileExpressionList() {

}
