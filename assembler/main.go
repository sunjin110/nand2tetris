package main

import (
	"assembler/pkg/code"
	"assembler/pkg/common"
	"assembler/pkg/common/jsonutil"
	"assembler/pkg/parser"
	"assembler/pkg/symboltable"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// .asmファイルを読み込んで、バイナリーコード(.hack)を作成する
func main() {
	log.Println("Assembler")

	// 引数を取得する
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("引数がありません")
		os.Exit(1)
	}

	// input file
	asmFileName := args[0]

	// シンボルテーブルを作成するために一度すべてを読む
	ffp, err := os.Open(asmFileName)
	if err != nil {
		fmt.Println("ファイルが見つかりません")
		panic(err)
	}
	// L命令のsymbolをmappingしているものを取得する
	symbolTableMap, isUseAddressMap := symboltable.GetSymbolTableMap(bufio.NewScanner(ffp))

	fp, err := os.Open(asmFileName)
	if err != nil {
		fmt.Println("ファイルが見つかりません")
		panic(err)
	}
	defer fp.Close()

	// output file
	hackFileName := fmt.Sprintf("%s.hack", common.GetFileNameWithoutExt(fp.Name()))
	hackFile, err := os.Create(hackFileName)
	if err != nil {
		fmt.Println("output fileが開けませんでした")
		panic(err)
	}
	defer hackFile.Close()

	scanner := bufio.NewScanner(fp)
	var lineCounter int
	addressCounter := 16
	for scanner.Scan() {

		line := scanner.Text()

		// コマンドのタイプを判別する
		commandType := parser.GetCommandType(line)
		if commandType == parser.NoneCommand {
			// NoneCommandの場合は何もしない
			continue
		}

		// debug
		if lineCounter < 8850 {
			// fmt.Println(lineCounter, "行目")
			// fmt.Println("line is ", line)
		}

		lineCounter++
		// debug

		var outLine string
		// A命令のとき
		switch commandType {
		case parser.ACommand:

			symbol := parser.GetSymbol(line, commandType)

			i, err := common.StrToUint(symbol)

			if lineCounter < 8850 {
				// fmt.Printf("symbol=[%s]\n", symbol)
			}

			if err == nil {

				// 数字の場合
				outLine = fmt.Sprintf("0%015b\n", i)
			} else {

				address, exists := symbolTableMap[symbol]

				if symbol == "ponggame.0" {
					fmt.Println("ponggame.0")
				}

				if exists {
					// もし存在する場合は、adressをそれに追加する
					outLine = fmt.Sprintf("0%015b\n", address)
				} else {
					// ない場合は新しい変数、使用できるアドレスを取得してそこに格納する
					var canUseAddress int
					canUseAddress, addressCounter = symboltable.GetCanUseAddressAndCounter(isUseAddressMap, addressCounter)
					isUseAddressMap[canUseAddress] = true
					symbolTableMap[symbol] = canUseAddress
					outLine = fmt.Sprintf("0%015b\n", canUseAddress)

					if symbol == "ponggame.0" && !exists {
						log.Println("address is ", canUseAddress)
						jsonutil.Print(symbolTableMap)
					}
				}

			}

		case parser.CCommand:
			outLine = getCBinary(line, commandType)
		case parser.LCommand:
			// LCommandのときは何もしない
			continue
		default:
			panic("想定していないcommandType")
		}

		// 書き込み
		_, err := hackFile.WriteString(outLine)
		if err != nil {
			log.Println("書き込みに失敗しました")
			panic(err)
		}

	}

}

// getCBinary C命令のbinaryを取得する
func getCBinary(line string, commandType parser.CommandType) string {

	dest, comp, jump := parser.GetCMemonic(line, commandType)

	// fmt.Println("jum is ", jump)
	// fmt.Printf("dest=[%s]\n", dest)
	// fmt.Printf("comp=[%s]\n", comp)
	// fmt.Printf("jump=[%s]\n", jump)

	destBinary := code.ConvDest(dest)
	compBinary := code.ConvComp(comp)
	jumpBinary := code.ConvJump(jump)

	return fmt.Sprintf("111%07b%03b%03b\n", compBinary, destBinary, jumpBinary)
}
