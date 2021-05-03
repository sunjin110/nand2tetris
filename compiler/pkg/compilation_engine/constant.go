package compilation_engine

// VariableKind 変数種類
type VariableKind string

const (
	// StaticVariableKind スタティック変数、全てのObjectで共通で利用される
	StaticVariableKind VariableKind = "static"

	// FieldVariableKind フィールド変数、クラスインスタンスごとにそれぞれ別の変数
	FieldVariableKind VariableKind = "field"

	// LocalVariableKind ローカル変数、Method内でのみ使用される
	LocalVariableKind VariableKind = "local"

	// ParamVariableKind 引数の変数
	ParamVariableKind VariableKind = "param"
)

// VariableType 型
type VariableType string

const (
	// IntType .
	IntType VariableType = "int"

	// CharType .
	CharType VariableType = "char"

	// BooleanType .
	BooleanType VariableType = "boolean"

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
