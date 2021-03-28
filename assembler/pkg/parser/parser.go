package parser

import "strings"

// HasMoreCommonds 入力にまだコマンドが存在するか?
// func HasMoreCommonds() bool

// Advance 入力から次のコマンドを読み、それを現在のコマンドにする
// このルーチンはHasMoreCommandsがtrueの場合のみ呼ぶようにする
// goだと必要ないかも
// func Advance()

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

	// @で始まる命令は、A命令
	if strings.HasPrefix(line, "@") {
		return ACommand
	}

	if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
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

	if commandType == ACommand {
		return line[1:]
	}

	// 前と後ろを削除する
	return line[1 : len(line)-1]
}
