package symboltable

import (
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"log"
)

// Engine .
type Engine struct {
	class       *compilation_engine.Class
	SymbolTable *SymbolTable
}

// SymbolTable class1つにつきのsymbol table
type SymbolTable struct {
	ClassSymbolList          []*Symbol
	SubroutineSymbolTableMap map[string]*SubroutineSymbolTable // key:SubroutineName
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
	log.Println("=== Make SymbolTable !!! ===")

	engine.SymbolTable = getSymbolTable(engine.class)

	// logger
	log.Println("symbol table is ", jsonutil.Marshal(engine.SymbolTable))

}

// getSymbolTable .
func getSymbolTable(class *compilation_engine.Class) *SymbolTable {
	return &SymbolTable{
		ClassSymbolList: getClassSymbolList(class.ClassVarDecList),
	}
}

// getClassSymbolList クラスのスコープにおけるシンボルテーブルを作成する
func getClassSymbolList(classVarDecList []*compilation_engine.ClassVarDec) []*Symbol {

	if len(classVarDecList) == 0 {
		return nil
	}

	// numberを定義する必要があるので
	// key: 属性(attribute), value: num
	numMap := map[string]int32{}

	var classSymbolList []*Symbol
	for _, classVarDec := range classVarDecList {
		for _, varName := range classVarDec.VarNameList {

			// 番号を取得
			num := numMap[string(classVarDec.VarKind)]
			symbol := createSymbol(varName, string(classVarDec.VarType), string(classVarDec.VarKind), num)
			classSymbolList = append(classSymbolList, symbol)

			// 番号を1incrementする
			numMap[string(classVarDec.VarKind)]++
		}
	}

	return classSymbolList
}
