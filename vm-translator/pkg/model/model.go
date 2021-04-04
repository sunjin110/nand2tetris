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
