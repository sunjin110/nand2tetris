package symboltable

import (
	"compiler/pkg/common/jsonutil"
	"compiler/pkg/compilation_engine"
	"log"
)

const (
	argument = "argument"
	variable = "var"
)

// Engine .
type Engine struct {
	class       *compilation_engine.Class
	SymbolTable *SymbolTable
}

// SymbolTable class1つにつきのsymbol table
type SymbolTable struct {
	ClassName                string
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
		ClassName:                class.ClassName,
		ClassSymbolList:          getClassSymbolList(class.ClassVarDecList),
		SubroutineSymbolTableMap: getSubroutineSymbolTableMap(class.SubRoutineDecList),
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

// getSubroutineSymbolTableMap サブルーチンリストのシンボルテーブルのMapを取得する
func getSubroutineSymbolTableMap(subRoutineDecList []*compilation_engine.SubRoutineDec) map[string]*SubroutineSymbolTable {

	if len(subRoutineDecList) == 0 {
		return nil
	}

	// key: method name
	subroutineSymbolTableMap := map[string]*SubroutineSymbolTable{}
	for _, subRoutineDec := range subRoutineDecList {
		// subroutineSymbolTableMap[subRoutineDec.SubRoutineName] = getSubroutineSymbolTable(subRoutineDec)
		subroutineSymbolTable := getSubroutineSymbolTable(subRoutineDec)
		if subroutineSymbolTable != nil {
			subroutineSymbolTableMap[subRoutineDec.SubRoutineName] = subroutineSymbolTable
		}
	}

	return subroutineSymbolTableMap
}

// getSubroutineSymbolTable サブルーチンのシンボルテーブルを取得する
func getSubroutineSymbolTable(subRoutineDec *compilation_engine.SubRoutineDec) *SubroutineSymbolTable {

	var subroutineSymbolList []*Symbol

	// numberを定義する必要があるので作成する
	// key: 属性(attribute), value: num
	numMap := map[string]int32{}

	// 先に、引数
	for _, parameter := range subRoutineDec.ParameterList {

		num := numMap[argument]
		symbol := createSymbol(parameter.ParamName, string(parameter.ParamType), argument, num)
		subroutineSymbolList = append(subroutineSymbolList, symbol)
		numMap[argument]++
	}

	// var
	for _, varDec := range subRoutineDec.SubRoutineBody.VarDecList {
		for _, varName := range varDec.NameList {
			num := numMap[variable]
			symbol := createSymbol(varName, string(varDec.Type), variable, num)
			subroutineSymbolList = append(subroutineSymbolList, symbol)
			numMap[variable]++
		}
	}

	// なければnil
	if len(subroutineSymbolList) == 0 {
		return nil
	}

	return &SubroutineSymbolTable{
		SubroutineName: subRoutineDec.SubRoutineName,
		SymbolList:     subroutineSymbolList,
	}
}
