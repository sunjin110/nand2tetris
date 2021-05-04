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
type Statement interface {
	GetStatementType() string
}

// LetStatement 代入の文
type LetStatement struct {
	DestVarName     string      // 代入先の変数名
	ArrayExpression *Expression // 格納する変数がArrayで、どの番地に追加するかがある場合は追加
	Expression      *Expression
}

// GetStatementType .
func (*LetStatement) GetStatementType() string {
	return LetStatementPrefix
}

// IfStatement IFの文
type IfStatement struct {
	ConditionalExpression *Expression // 条件式
	StatementList         []Statement // 式
	ElseStatementList     []Statement // elseがある場合の式
}

// GetStatementType .
func (*IfStatement) GetStatementType() string {
	return IfStatementPrefix
}

// WhileStatement Whileの文
type WhileStatement struct {
	ConditionalExpression *Expression // 条件式
	StatementList         []Statement // 実行するStatement
}

// GetStatementType .
func (*WhileStatement) GetStatementType() string {
	return WhileStatementPrefix
}

// DoStatement doの文
type DoStatement struct {
	SubroutineCall *SubRoutineCall
}

// GetStatementType .
func (*DoStatement) GetStatementType() string {
	return DoStatementPrefix
}

// ReturnStatement return文
type ReturnStatement struct {
	ReturnExpression *Expression
}

// GetStatementType .
func (*ReturnStatement) GetStatementType() string {
	return ReturnStatementPrefix
}
