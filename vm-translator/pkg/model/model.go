package model

// 定数とか、他のstructなんかを取り扱う

// CommandType コマンドタイプ
type CommandType string

const (
	// CommandTypeArithmetic 算術コマンド
	CommandTypeArithmetic CommandType = "C_ARITHMETIC"

	// CommandTypePush .
	CommandTypePush CommandType = "C_PUSH"

	// CommandTypePop .
	CommandTypePop CommandType = "C_POP"

	// CommandTypeLabel .
	CommandTypeLabel CommandType = "C_LABEL"

	// CommandTypeGoto .
	CommandTypeGoto CommandType = "C_GOTO"

	// CommandTypeIf .
	CommandTypeIf CommandType = "C_IF"

	// CommandTypeFunction .
	CommandTypeFunction CommandType = "C_FUNCTION"

	// CommandTypeReturn .
	CommandTypeReturn CommandType = "C_RETURN"

	// CommandTypeCall .
	CommandTypeCall CommandType = "C_CALL"
)

// ArithmeticCommand 算術コマンド

const (
	// ArithmeticAdd 足し算 x+y
	ArithmeticAdd = "add"

	// AirthmeticSub 引き算 x-y
	AirthmeticSub = "sub"

	// AirthmeticNeg -y
	AirthmeticNeg = "neg"

	// AirthmeticEq x = y であればtrue, else is false
	AirthmeticEq = "eq"

	// AirthmeticGt x > y であればtrue
	AirthmeticGt = "gt"

	// AirthmeticLt x < y であればtrue
	AirthmeticLt = "lt"

	// AirthmeticAnd x AND y
	AirthmeticAnd = "and"

	// AirthmeticOr x OR y
	AirthmeticOr = "or"

	// AirthmeticNot x NOT y
	AirthmeticNot = "not"
)

// MemoryAccessCommand メモリにアクセスするためのもの

const (
	// MemoryAccess .
	MemoryAccessPush = "push"

	// Pop .
	MemoryAccessPop = "pop"
)

// MemorySegment メモリを分割したものの定義

const (
	// MemorySegmentArgument 関数の引数を格納
	MemorySegmentArgument = "argument"

	// MemorySegmentLocal 関数のローカル変数を格納
	MemorySegmentLocal = "local"

	// MemorySegmentStatic スタティック変数を格納する、スタティック変数は、同じ.vmファイルの全ての関数で共有される
	MemorySegmentStatic = "static"

	// MemorySegmentConstant 0から32767までの範囲を全ての定数値を持つ擬似セグメント
	MemorySegmentConstant = "constant"

	// MemorySegmentThis 汎用セグメント
	MemorySegmentThis = "this"

	// MemorySegmentThat 汎用セグメント
	MemorySegmentThat = "that"

	// MemorySegmentPointer thisとthatセグメントのベースアドレス(参照)を持つ2つの要素からなるセグメント
	MemorySegmentPointer = "pointer"

	// MemorySegmentTemp 固定された8つの要素からなるセグメント。一時的な変数を格納するために用いる
	MemorySegmentTemp = "temp"
)

const (
	// Label .
	Label = "label"

	// IfGoto .
	IfGoto = "if-goto"

	// Goto .
	Goto = "goto"

	// Function .
	Function = "function"

	// Return .
	Return = "return"

	// Call .
	Call = "call"
)
