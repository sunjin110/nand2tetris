// parser アセンブリコマンドをフィールドとシンボルに分解する

package parser

import (
	"strings"
)

type CommandType int32

const (
	// NoneCommand 無視する対象
	// コメント「//」やただの改行
	NoneCommand CommandType = 0

	// ACommand @Xxxを意味する、Xxxはシンボルか10進数
	ACommand CommandType = 1

	// CCommanddest=comp;jumpを意味する
	CCommand CommandType = 2

	// LCommand 疑似コマンド(Xxx)を意味する、Xxxはシンボル
	LCommand CommandType = 3
)

// GetCommandType コマンドのタイプを取得する
func GetCommandType(line string) CommandType {

	// 空白や「//」の場合は、無効な行
	if line == "" || strings.HasPrefix(line, "//") {
		return NoneCommand
	}

	// 無駄な空白を取る
	// command := strings.TrimSpace(line)
	command := trimeLine(line)

	// @で始まる命令は、A命令
	if strings.HasPrefix(command, "@") {
		return ACommand
	}

	if strings.HasPrefix(command, "(") && strings.HasSuffix(command, ")") {
		return LCommand
	}

	return CCommand
}

// GetSymbol 現コマンドのシンボルを取得する
// @Xxxxまたは、(Xxx)のXxxを返す
// Xxxはシンボルまたは、10進数の数字である
func GetSymbol(line string, commandType CommandType) string {

	if commandType == NoneCommand || commandType == CCommand {
		panic("想定されないコマンドが渡されました")
	}

	command := trimeLine(line)

	if commandType == ACommand {
		// @symbol -> symbol
		return command[1:]
	}

	// 前と後ろを削除する
	// (symbol) -> symbol
	return command[1 : len(command)-1]
}

// GetCMemonic C命令のニーモニックを
// dest, comp, jumpで分割して返す
// dest=comp;jump
func GetCMemonic(line string, commandType CommandType) (string, string, string) {

	if commandType != CCommand {
		panic("C命令以外は想定していません")
	}

	// 無駄な空白とかを削除する
	command := trimeLine(line)

	eqlIndex := strings.Index(command, "=")
	semicolonIndex := strings.Index(command, ";")

	// dest
	var dest string
	if eqlIndex > 0 {
		dest = command[:eqlIndex]
	}

	// jump
	var jump string
	if semicolonIndex != -1 {
		jump = command[semicolonIndex+1:]
		jump = jump[:3]
	}

	// comp
	compPreIndex := 0
	compSufIndex := len(command)
	if eqlIndex > 0 {
		compPreIndex = eqlIndex + 1
	}
	if semicolonIndex > 0 {
		compSufIndex = semicolonIndex
	}
	comp := command[compPreIndex:compSufIndex]

	return strings.TrimSpace(dest), strings.TrimSpace(comp), strings.TrimSpace(jump)
}

// 不要な空白や、コメントを削除する
func trimeLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "//")[0])
}
