package compilation_engine

// IsClassVarDecPrefixToken Classのfield宣言の先頭かどうかを判定する
func IsClassVarDecPrefixToken(token string) bool {
	return token == string(StaticVariableKind) || token == string(FieldVariableKind)
}

// IsSubRoutineDecPrefixToken SubRoutineの先頭かどうかを判定する
func IsSubRoutineDecPrefixToken(token string) bool {
	t := SubRoutineKind(token)
	return t == Constructor || t == Function || t == Method
}

// IsVarDecPrefixToken varDecの先頭のtokenかどうかを確認する
func IsVarDecPrefixToken(token string) bool {
	return token == string(LocalVariableKind)
}

// IsStatementPrefixToken statementの先頭のtokenかどうかを判定する
func IsStatementPrefixToken(token string) bool {
	return token == LetStatementPrefix ||
		token == IfStatementPrefix ||
		token == WhileStatementPrefix ||
		token == DoStatementPrefix ||
		token == ReturnStatementPrefix
}
