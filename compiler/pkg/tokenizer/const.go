package tokenizer

// TokenType トークンのタイプ
type TokenType string

// KeyWord キーワード
type KeyWord string

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
