package compilation_engine

// 式の宣言

const (
	// IntegerConstType
	IntegerConstType = "integerConstType"

	// StringConstType .
	StringConstType = "stringConstType"

	// KeyWordConstType .
	KeyWordConstType = "keyWordConstType"

	// ValNameConstType .
	ValNameConstType = "valNameConstType"

	// SubRoutineCollType .
	SubRoutineCallType = "subRoutineCollType"

	// ExpressionType .
	ExpressionType = "expressionType"

	// UnaryOpTermType .
	UnaryOpTermType = "unaryOpTermType"
)

// Expression 式
type Expression struct {
	InitTerm   Term      // 必ず必要な式
	OpTermList []*OpTerm // Optionで付属するTermたち
}

// OpTerm 符号とセットでもつterm
type OpTerm struct {
	Operation Op // 符号
	OpTerm    Term
}

// Term 式のinterface
type Term interface {
	GetTermType() string
}

// IntegerConstTerm intの値を持つだけのterm
type IntegerConstTerm struct {
	Val int // value
}

// GetTermType .
func (*IntegerConstTerm) GetTermType() string {
	return IntegerConstType
}

// StringConstTerm stringの値を持つだけのterm
type StringConstTerm struct {
	Val string // value
}

// GetTermType .
func (*StringConstTerm) GetTermType() string {
	return StringConstType
}

// KeyWordConstTerm
type KeyWordConstTerm struct {
	KeyWord KeywordConstant
}

// GetTermType .
func (*KeyWordConstTerm) GetTermType() string {
	return KeyWordConstType
}

// ValNameConstantTerm 変数名のみを持つ
type ValNameConstantTerm struct {
	ValName         string
	ArrayExpression *Expression // arrayの場合はこれがついてる
}

// GetTermType .
func (*ValNameConstantTerm) GetTermType() string {
	return ValNameConstType
}

// SubRoutineCall termでもあるんやで
type SubRoutineCall struct {
	ClassOrVarName string // 空白の場合は指定なし
	SubRoutineName string
	ExpressionList []*Expression
}

// GetTermType .
func (*SubRoutineCall) GetTermType() string {
	return SubRoutineCallType
}

// ExpressionTerm ()に包まれているterm
type ExpressionTerm struct {
	Expression *Expression
}

// GetTermType .
func (*ExpressionTerm) GetTermType() string {
	return ExpressionType
}

// UnaryOpTerm 終端ではない
type UnaryOpTerm struct {
	UnaryOp UnaryOp
	Term    Term
}

// GetTermType .
func (*UnaryOpTerm) GetTermType() string {
	return UnaryOpTermType
}

// Op 演算子
type Op rune

const (
	// PlusOp .
	PlusOp = '+'

	// MinusOp .
	MinusOp = '-'

	// AsteriskOp .
	AsteriskOp = '*'

	// SlashOp .
	SlashOp = '/'

	// AndOp .
	AndOp = '&'

	// PipeOp .
	PipeOp = '|'

	// LessThanOp .
	LessThanOp = '<'

	// GreaterThanOp .
	GreaterThanOp = '>'

	// EqlOp .
	EqlOp = '='
)

// UnaryOp .
type UnaryOp rune

const (
	// HypenUnaryOp .
	HypenUnaryOp UnaryOp = '-'

	// TildeUnaryOp .
	TildeUnaryOp UnaryOp = '~'
)

// KeywordConstant .
type KeywordConstant string

const (
	// TrueKeyword .
	TrueKeyword = "true"

	// FalseKeyword .
	FalseKeyword = "false"

	// NullKeyword .
	NullKeyword = "null"

	// ThisKeyword .
	ThisKeyword = "this"
)
