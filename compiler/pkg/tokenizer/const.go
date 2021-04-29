package tokenizer

// TokenType トークンのタイプ
type TokenType string

const (
	// TokenTypeKeyWord 予約後のtoken
	TokenTypeKeyWord = "KEYWORD"

	// TokenTypeSymbol 演算子などのシンボルのtoken
	TokenTypeSymbol = "SYMBOL"

	// TokenTypeIdentifier 変数名
	TokenTypeIdentifier = "IDENTIFIER"

	// TokenTypeIntConst 数字
	TokenTypeIntConst = "INT_CONST"

	// TokenTypeStringConst 文字列
	TokenTypeStringConst = "STRING_CONST"
)

// KeyWord キーワード
type KeyWord string

const (
	// KeyWordClass .
	KeyWordClass = "CLASS"

	// KeyWordMethod .
	KeyWordMethod = "METHOD"

	// KeyWordFunction .
	KeyWordFunction = "FUNCTION"

	// KeyWordConstructor .
	KeyWordConstructor = "CONSTRUCTOR"

	// KeyWordInt .
	KeyWordInt = "INT"

	// KeyWordBoolean .
	KeyWordBoolean = "BOOLEAN"

	// KeyWordChar .
	KeyWordChar = "CHAR"

	// KeyWordVoid .
	KeyWordVoid = "VOID"

	// KeyWordVar .
	KeyWordVar = "VAR"

	// KeyWordStatic .
	KeyWordStatic = "STATIC"

	// KeyWordField .
	KeyWordField = "FIELD"

	// KeyWordLet .
	KeyWordLet = "LET"

	// KeyWordDo .
	KeyWordDo = "DO"

	// KeyWordIf .
	KeyWordIf = "IF"

	// KeyWordElse .
	KeyWordElse = "ELSE"

	// KeyWordWhile .
	KeyWordWhile = "WHILE"

	// KeyWordReturn .
	KeyWordReturn = "RETURN"

	// KeyWordTrue .
	KeyWordTrue = "TRUE"

	// KeyWordFalse .
	KeyWordFalse = "FALSE"

	// KeyWordNull .
	KeyWordNull = "NULL"

	// KeyWordThis .
	KeyWordThis = "THIS"
)
