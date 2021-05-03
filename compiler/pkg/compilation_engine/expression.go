package compilation_engine

// 式の宣言

// Expression 式
type Expression struct {
}

// Term .
type Term interface{}

// ValNameConstantTerm .
type ValNameConstantTerm struct {
	ValName string
}

// SubRoutineCall .
type SubRoutineCall struct {
	ClassName      *string
	VarName        *string
	SubRoutineName string
	ExpressionList []*Expression
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
	TildeUnaryOp UnaryOp = '!'
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
