package compilation_engine

// 文章の構造宣言

const (
	// LetStatementPrefix .
	LetStatementPrefix = "let"

	// IfStatementPrefix .
	IfStatementPrefix = "if"

	// WhileStatementPrefix .
	WhileStatementPrefix = "while"

	// DoStatementPrefix .
	DoStatementPrefix = "do"

	// ReturnStatementPrefix .
	ReturnStatementPrefix = "return"
)

// Statement 一番大きな単位
// TODO interface化
type Statement interface {
	IsStatement() bool
}

// LetStatement 代入の文
type LetStatement struct {
	DestVarName string // 代入先の変数名
	// TODO arrayとかに対応できる構造
	Expression Expression
}

func (l *LetStatement) IsStatement() bool {
	return true
}

// IfStatement IFの文
type IfStatement struct {
	ConditionalExpression *Expression  // 条件式
	StatementList         []*Statement // 式
	ElseStatementList     []*Statement // elseがある場合の式
}

// WhileStatement Whileの文
type WhileStatement struct {
	ConditionalExpression *Expression  // 条件式
	StatementList         []*Statement // 実行するStatement
}

// DoStatement doの文
type DoStatement struct {
	SubroutineCall *SubRoutineCall
}

// ReturnStatement return文
type ReturnStatement struct {
	ReturnExpression *Expression
}
