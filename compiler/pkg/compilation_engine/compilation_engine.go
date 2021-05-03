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
	if !IsClassVarDecPrefixTokne(t) {
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
		// すでに他のtokenに写っている場合は、skip
		c.nextToken()
		checkStr := c.getToken()
		if !IsClassVarDecPrefixTokne(checkStr) {
			break
		}
	}

	return classVarDecList
}

// compileSubroutine メソッド、ファンクション、コンストラクタをコンパイルする
func (c *CompilationEngine) compileSubroutine() []*SubRoutineDec {
	return nil
}

// compileParameterList パラメータのリスト(空の可能性もある)をコンパイルする。カッコ"()"は含まない
func compileParameterList() {

}

// compileVarDec var宣言をコンパイルする
func compileVarDec() {

}

// compileStatements 一連の文をコンパイルする。波括弧"{}"は含まない
func compileStatements() {

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
