package symboltable

import (
	"compiler/pkg/common/chk"
	"compiler/pkg/compilation_engine"
	"fmt"
)

// Engine .
type Engine struct {
	class       *compilation_engine.Class
	SymbolTable *SymbolTable
}

// SymbolTable class1つにつきのsymbol table
type SymbolTable struct {
	ClassName                string
	ClassSymbolMap           map[string]*Symbol                // key:varName
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
		ClassSymbolMap:           getClassSymbolMap(class.ClassName, class.ClassVarDecList),
		SubroutineSymbolTableMap: getSubroutineSymbolTableMap(class.ClassName, class.SubRoutineDecList),
	}
}

// getClassSymbolMap クラスのスコープにおけるシンボルテーブルを作成する
func getClassSymbolMap(className string, classVarDecList []*compilation_engine.ClassVarDec) map[string]*Symbol {

	if len(classVarDecList) == 0 {
		return nil
	}

	// numberを定義する必要があるので
	// key: 属性(attribute), value: num
	numMap := map[string]int32{}

	classSymbolMap := map[string]*Symbol{}
	for _, classVarDec := range classVarDecList {
		for _, varName := range classVarDec.VarNameList {

			// 番号を取得
			num := numMap[string(classVarDec.VarKind)]
			symbol := createSymbol(varName, string(classVarDec.VarType), string(classVarDec.VarKind), num)

			// すでにその変数が存在するかどうかを確認する
			if _, ok := classSymbolMap[symbol.VarName]; ok {
				chk.SE(fmt.Errorf("class:%sで変数%sが複数宣言されています", className, symbol.VarName))
			}

			classSymbolMap[symbol.VarName] = symbol

			// 番号を1incrementする
			numMap[string(classVarDec.VarKind)]++
		}
	}

	return classSymbolMap
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
	argAttri := string(compilation_engine.ArgumentVariableKind)
	for _, parameter := range subRoutineDec.ParameterList {

		num := numMap[argAttri]
		symbol := createSymbol(parameter.ParamName, string(parameter.ParamType), argAttri, num)

		if _, ok := subroutineSymbolMap[symbol.VarName]; ok {
			chk.SE(fmt.Errorf("class:%s subRoutine:%s 内で変数%sが複数宣言されています", className, subRoutineDec.SubRoutineName, symbol.VarName))
		}

		subroutineSymbolMap[symbol.VarName] = symbol
		numMap[argAttri]++
	}

	// var
	varAttri := string(compilation_engine.LocalVariableKind)
	for _, varDec := range subRoutineDec.SubRoutineBody.VarDecList {
		for _, varName := range varDec.NameList {
			num := numMap[varAttri]
			symbol := createSymbol(varName, string(varDec.Type), varAttri, num)

			if _, ok := subroutineSymbolMap[symbol.VarName]; ok {
				chk.SE(fmt.Errorf("class:%s subRoutine:%s 内で変数%sが複数宣言されています", className, subRoutineDec.SubRoutineName, symbol.VarName))
			}

			subroutineSymbolMap[symbol.VarName] = symbol
			numMap[varAttri]++
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
