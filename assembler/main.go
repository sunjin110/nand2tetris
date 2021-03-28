package main

import (
	"assembler/pkg/code"
	"assembler/pkg/common"
	"assembler/pkg/parser"
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

	asmFileName := args[0]
	fp, err := os.Open(asmFileName)
	if err != nil {
		fmt.Println("ファイルが見つかりません")
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()

		// コマンドのタイプを判別する
		commandType := parser.GetCommandType(line)
		if commandType == parser.NoneCommand {
			// NoneCommandの場合は何もしない
			continue
		}

		// A命令のとき || L命令のとき
		switch commandType {
		case parser.ACommand, parser.LCommand:
			symbol := parser.GetSymbol(line, commandType)

			// 数字の場合
			i := common.StrToUint(symbol)
			fmt.Printf("0%015b\n", i)

		case parser.CCommand:
			dest, comp, jump := parser.GetCMemonic(line, commandType)
			// log.Println(dest, comp, jump)

			destBinary := code.ConvDest(dest)
			compBinary := code.ConvComp(comp)
			jumpBinary := code.ConvJump(jump)

			// TODO output file
			fmt.Printf("111%07b%03b%03b\n", compBinary, destBinary, jumpBinary)
		default:
			panic("想定していないcommandType")
		}

	}

}
