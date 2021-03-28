package symboltable

import (
	"assembler/pkg/parser"
	"bufio"
	"fmt"
)

// SymbolTableMap シンボルテーブル
type SymbolTableMap map[string]int

// GetSymbolTableMap アセンブラからsymbolTableを作成する
// L命令のものはAdressは現在のread中の行数を読む
// 定義済みシンボルもここで追加しておく
// ２つ目の引数には、symbolTableMapですでに割り当てられているアドレスかどうかを調べるためにつかう
func GetSymbolTableMap(scanner *bufio.Scanner) (SymbolTableMap, map[int]bool) {

	// 定義済みシンボルを追加する
	symbolTableMap := SymbolTableMap{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	isUseAddressMap := map[int]bool{}
	for _, address := range symbolTableMap {
		isUseAddressMap[address] = true
	}

	var addressCnt int
	for scanner.Scan() {

		line := scanner.Text()

		// コマンドタイプを取得する
		commandType := parser.GetCommandType(line)

		// A命令, C命令の場合はカウントを1UPする
		if commandType == parser.ACommand || commandType == parser.CCommand {
			addressCnt++
		}

		// L命令以外はここで処理終了
		if commandType != parser.LCommand {
			continue
		}

		// symbol名を取得する
		symbol := parser.GetSymbol(line, commandType)

		// すでにそのsymbol名が使われているかどうかを確認する
		beforeAddress, exists := symbolTableMap[symbol]
		if exists {
			fmt.Printf("symbolが重複しています、symbol:%s, address: %d and %d\n", symbol, beforeAddress, addressCnt)
			panic("symbol duplicated error")
		}

		// 追加する
		symbolTableMap[symbol] = addressCnt
		isUseAddressMap[addressCnt] = true

	}

	return symbolTableMap, isUseAddressMap
}

// GetCanUseAddressAndCounter 使用できるAddressを探索する、
// ２つ目の引数に、スキップしたカウンタを取得する
func GetCanUseAddressAndCounter(isUseAddressMap map[int]bool, addressCounter int) (int, int) {

	for {

		if !isUseAddressMap[addressCounter] {
			// 使用していないアドレスを見つけた場合、
			// そのアドレスを使用済みにして、返す
			// カウンタも1プラスする
			// isUseAddressMap[addressCounter] = true
			return addressCounter, addressCounter + 1
		}

		// そのアドレスが使用されている場合、カウンタを1プラスして次に備える
		addressCounter++
	}

}

// // GetCanUseAddressAndCounter 使用できるAddressを探索する、
// // ２つ目の引数に、スキップしたカウンタを取得する
// func GetCanUseAddressAndCounter(isUseAddressMap map[int]bool, addressCounter int) (int, int) {

// 	for {

// 		if !isUseAddressMap[addressCounter] {
// 			// 使用していないアドレスを見つけた場合、
// 			// そのアドレスを使用済みにして、返す
// 			// カウンタも1プラスする
// 			// isUseAddressMap[addressCounter] = true
// 			return addressCounter, addressCounter + 1
// 		}

// 		// そのアドレスが使用されている場合、カウンタを1プラスして次に備える
// 		addressCounter++
// 	}

// }

// // AddEntry テーブルにkey:symbol, value:addressのペアを追加する
// func AddEntry(symbolTableMap SymbolTableMap, symbol string, address int) SymbolTableMap {
// 	symbolTableMap[symbol] = address
// 	return symbolTableMap
// }

// // Contains 与えられたsymbolを含むか?
// func Contains(symbolTableMap SymbolTableMap, symbol string) bool {
// 	_, exists := symbolTableMap[symbol]
// 	return exists
// }

// // GetAddress アドレスを取得する
// func GetAddress(symbolTableMap SymbolTableMap, symbol string) int {
// 	return symbolTableMap[symbol]
// }
