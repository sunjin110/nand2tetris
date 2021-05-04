package compilation_engine

// IsClassVarDecPrefixTokne Classのfield宣言の先頭かどうかを判定する
func IsClassVarDecPrefixTokne(token string) bool {
	return token == string(StaticVariableKind) || token == string(FieldVariableKind)
}

// IsSubRoutineDecPrefixToken SubRoutineの先頭かどうかを判定する
func IsSubRoutineDecPrefixToken(token string) bool {
	t := SubRoutineKind(token)
	return t == Constructor || t == Function || t == Method
}
