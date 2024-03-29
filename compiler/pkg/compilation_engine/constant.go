package compilation_engine

// VariableKind 変数種類
type VariableKind string

const (
	// StaticVariableKind スタティック変数、全てのObjectで共通で利用される
	StaticVariableKind VariableKind = "static"

	// FieldVariableKind フィールド変数、クラスインスタンスごとにそれぞれ別の変数
	FieldVariableKind VariableKind = "field"

	// LocalVariableKind ローカル変数、Method内でのみ使用される
	LocalVariableKind VariableKind = "var"

	// ArgumentVariableKind 引数の変数
	ArgumentVariableKind VariableKind = "argument"
)

// VariableType 型
type VariableType string

// IsPrimitive この型がprimitiveかどうか？
func (vt VariableType) IsPrimitive() bool {

	switch vt {
	case IntType, CharType, BooleanType, VoidType:
		return true
	}
	return false
}

const (
	// IntType .
	IntType VariableType = "int"

	// CharType .
	CharType VariableType = "char"

	// BooleanType .
	BooleanType VariableType = "boolean"

	// VoidType .
	VoidType VariableType = "void"

	// ClassNameType 独自で宣言したクラス型
	ClassNameType VariableType = "className"
)

// SubRoutineKind methodなどのタイプ
type SubRoutineKind string

const (
	// Constructor 初期化のやつ
	Constructor SubRoutineKind = "constructor"

	// Function 関数(static method)
	Function SubRoutineKind = "function"

	// Method メソッド
	Method SubRoutineKind = "method"
)
