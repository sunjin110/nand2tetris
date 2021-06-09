package symboltable

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/compilation_engine"
	"fmt"
)

const (
	argument = "argument"
	variable = "local"
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
	// SymbolList     []*Symbol
	SymbolMap map[string]*Symbol // key: VarName, value: Symbol, varNameはuniqueである必要がある
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

// Start SymbolTable作成開始
func (engine *Engine) Start() {
	engine.SymbolTable = getSymbolTable(engine.class)
}

// getSymbolTable .
func getSymbolTable(class *compilation_engine.Class) *SymbolTable {
	return &SymbolTable{
		ClassName:                class.ClassName,
		ClassSymbolList:          getClassSymbolList(class.ClassVarDecList),
		SubroutineSymbolTableMap: getSubroutineSymbolTableMap(class.ClassName, class.SubRoutineDecList),
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
func getSubroutineSymbolTableMap(className string, subRoutineDecList []*compilation_engine.SubRoutineDec) map[string]*SubroutineSymbolTable {

	if len(subRoutineDecList) == 0 {
		return nil
	}

	// key: method name
	subroutineSymbolTableMap := map[string]*SubroutineSymbolTable{}
	for _, subRoutineDec := range subRoutineDecList {
		subroutineSymbolTable := getSubroutineSymbolTable(className, subRoutineDec)
		if subroutineSymbolTable != nil {
			subroutineSymbolTableMap[subRoutineDec.SubRoutineName] = subroutineSymbolTable
		}
	}

	return subroutineSymbolTableMap
}

// getSubroutineSymbolTable サブルーチンのシンボルテーブルを取得する
func getSubroutineSymbolTable(className string, subRoutineDec *compilation_engine.SubRoutineDec) *SubroutineSymbolTable {

	// numberを定義する必要があるので作成する
	// key: 属性(attribute), value: num
	numMap := map[string]int32{}

	// varNameでuniqueである必要がある
	subroutineSymbolMap := map[string]*Symbol{}

	// 先に、引数
	for _, parameter := range subRoutineDec.ParameterList {

		num := numMap[argument]
		symbol := createSymbol(parameter.ParamName, string(parameter.ParamType), argument, num)

		if _, ok := subroutineSymbolMap[symbol.VarName]; ok {
			chk.SE(fmt.Errorf("Class:%s SubRoutine:%s 内で変数%sが複数宣言されています", className, subRoutineDec.SubRoutineName, symbol.VarName))
		}

		subroutineSymbolMap[symbol.VarName] = symbol
		numMap[argument]++
	}

	// var
	for _, varDec := range subRoutineDec.SubRoutineBody.VarDecList {
		for _, varName := range varDec.NameList {
			num := numMap[variable]
			symbol := createSymbol(varName, string(varDec.Type), variable, num)

			if _, ok := subroutineSymbolMap[symbol.VarName]; ok {
				chk.SE(fmt.Errorf("Class:%s SubRoutine:%s 内で変数%sが複数宣言されています", className, subRoutineDec.SubRoutineName, symbol.VarName))
			}

			subroutineSymbolMap[symbol.VarName] = symbol
			numMap[variable]++
		}
	}

	// なければnil
	if len(subroutineSymbolMap) == 0 {
		return nil
	}

	return &SubroutineSymbolTable{
		SubroutineName: subRoutineDec.SubRoutineName,
		SymbolMap:      subroutineSymbolMap,
	}
}
