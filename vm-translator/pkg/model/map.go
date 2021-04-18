package model

// Arg1CommandTypeMap Arg1からそのコマンドがなんのタイプなのかを紐付けるMap
var Arg1CommandTypeMap = map[string]CommandType{
	ArithmeticAdd:    CommandTypeArithmetic,
	AirthmeticSub:    CommandTypeArithmetic,
	AirthmeticNeg:    CommandTypeArithmetic,
	AirthmeticEq:     CommandTypeArithmetic,
	AirthmeticGt:     CommandTypeArithmetic,
	AirthmeticLt:     CommandTypeArithmetic,
	AirthmeticAnd:    CommandTypeArithmetic,
	AirthmeticOr:     CommandTypeArithmetic,
	AirthmeticNot:    CommandTypeArithmetic,
	MemoryAccessPush: CommandTypePush,
	MemoryAccessPop:  CommandTypePop,
	Label:            CommandTypeLabel,
	IfGoto:           CommandTypeIf,
	Goto:             CommandTypeGoto,
	Function:         CommandTypeFunction,
	Return:           CommandTypeReturn,
	// TODO いっぱいある
}
