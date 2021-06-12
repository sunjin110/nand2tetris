package vmwriter

const (
	whileStartLabelPattern = "WHILE_START_%d"
	whileEndLabelPattern   = "WHILE_END_%d"
	ifTrueLabelPattern     = "IF_TRUE_%d"
	ifFalseLabelPattern    = "IF_FALSE_%d"
	ifEndLabelPattern      = "IF_END_%d"
)

const (
	segmentConst    = "constant"
	segmentTemp     = "temp"
	segmentLocal    = "local"
	segmentPointer  = "pointer"
	segmentArgument = "argument"
	segmentThis     = "this"
	// TODO more segement type

)

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
