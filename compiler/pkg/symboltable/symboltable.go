package symboltable

// SymbolTable .
type SymbolTable struct {
	ClassName string
	RowList   []*Row // シンボルテーブルのあれ
	// TODO その中のscopeのものを設定できる必要がある

}

// Row 一つのデータのまとまり
type Row struct {
	VarName   string // 変数名
	Type      string // int | String | boolean | ClassName
	Attribute string // static | field | argument | local . TODO 変数名変えてもいいかも
	Num       int32  // 番号
	// ???
	SymbolTable *SymbolTable // これが
}

// // SymbolTable .
// type SymbolTable struct {
// 	VarName string // 変数名
// 	Type    string // int | String | boolean | char

// 	Num     int32  // 番号
// }

// // Scope .
