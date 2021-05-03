package compilation_engine

// IsClassVarDecPrefixTokne Classのfield宣言の先頭かどうかを判定する
func IsClassVarDecPrefixTokne(token string) bool {
	return token == string(StaticVariableKind) || token == string(FieldVariableKind)
}
