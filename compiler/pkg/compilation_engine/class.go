package compilation_engine

// Class Classの構造
type Class struct {
	ClassName         string
	ClassVarDecList   []*ClassVarDec
	SubRoutineDecList []*SubRoutineDec
}

// ClassVarDec Classの変数の宣言部分
type ClassVarDec struct {
	VarKind     VariableKind // static or field
	VarType     VariableType // int or char or boolean or className
	VarNameList []string     // 複数宣言することができるので「,」で分割可能
}

// SubRoutineDec methodなどのやつ
type SubRoutineDec struct {
	RoutineKind    SubRoutineKind  // constructor or function or method
	ReturnType     VariableType    // void or int or char or boolean or className
	SubRoutineName string          // sub routineの名前
	ParameterList  []*Parameter    // 引数
	SubRoutineBody *SubRoutineBody // body
}

// SubRoutineBody .
type SubRoutineBody struct {
	VarDecList    []*VarDec // 型の宣言など
	StatementList []Statement
}

// Parameter 引数
type Parameter struct {
	ParamType VariableType
	ParamName string
}

// VarDec 宣言
type VarDec struct {
	Type     VariableType // 型
	NameList []string     // 複数宣言する場合があるため
}
