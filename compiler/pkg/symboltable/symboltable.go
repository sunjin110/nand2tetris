package symboltable

import "compiler/pkg/compilation_engine"

// Engine .
type Engine struct {
	class       *compilation_engine.Class
	SymbolTable *SymbolTable
}

// SymbolTable class1つにつきのsymbol table
type SymbolTable struct {
	ClassSymbolList          []*Symbol
	SubroutineSymbolTableMap map[string]*SubroutineSymbolTable
}

// SubroutineSymbolTable subroutineのSymbolTable
type SubroutineSymbolTable struct {
	SubroutineName string
	SymbolList     []*Symbol
}

// Symbol 一つの要素
type Symbol struct {
	VarName   string // 変数名
	Type      string // int | String | boolean | ClassName
	Attribute string // static | field | argument | local
	Num       int32  // 番号
}

// New .
func New(class *compilation_engine.Class) *Engine {
	return &Engine{
		class:       class,
		SymbolTable: nil, // StartでここのsymbolTableを構築する
	}
}

// Start SymbolTable作成かいし
func (engine *Engine) Start() {
	// TODO
}
