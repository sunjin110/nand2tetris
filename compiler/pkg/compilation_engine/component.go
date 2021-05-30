package compilation_engine

import (
	"strconv"
	"unicode"
)

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

// IsExpressionListPrefixToken expression listのものかどうかを判断する
func IsExpressionListPrefixToken(token string) bool {
	return IsTermPrefixToken(token)
}

// IsTermPrefixToken termの先頭のtokenかどうかを判定する
func IsTermPrefixToken(token string) bool {

	// 空白の場合は、Error
	if token == "" {
		panic("Termのtokenが空です")
	}

	// 先頭が(の場合OK
	if rune(token[0]) == '(' {
		return true
	}

	// 先頭が[の場合OK
	if rune(token[0]) == '[' {
		return true
	}

	// 先頭が"の場合はOK
	if rune(token[0]) == '"' {
		return true
	}

	// 先頭がUnaryOkの場合OK
	if rune(token[0]) == rune(HypenUnaryOp) || rune(token[0]) == rune(TildeUnaryOp) {
		return true
	}

	// 数字に変換できるならOK
	_, err := strconv.Atoi(token)
	if err == nil {
		return true
	}

	// 変数ならOK
	return isVariableToken(token)

}

// 指定したtokenが変数(先頭がアルファベット, それ以外はアルファベットor数字orアンダースコア)
func isVariableToken(token string) bool {

	// 空白の場合は、除外
	if token == "" {
		return false
	}

	// 先頭の文字は、アルファベットまたは_でない場合はだめ
	if !unicode.IsLetter(rune(token[0])) && rune(token[0]) != '_' {
		return false
	}

	for _, s := range token {

		// 変数、数字、アンダースコアではない場合は、変数ではない
		if !unicode.IsLetter(s) && !unicode.IsNumber(s) && s != '_' {
			return false
		}
	}

	// 一番後ろが「)」の場合は、subroutineCall確定
	return token[len(token)-1] != ')'
}
