package parser

import (
	"bufio"
	"os"
	"vm-translator/pkg/model"
)

// 一つの.vmファイルに対してパースをおこなる、
// 入力コードへのアクセスをカプセル化する
// 空白文字とコメントを取り除く

// Parser .
type Parser struct {
	file        *os.File          // file
	scanner     *bufio.Scanner    // fileのscanner
	Command     string            // 現在参照しているコマンド
	CommandType model.CommandType // 現在参照しているコマンドのタイプを格納する
}

// New .
func New(filePath string) (*Parser, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return &Parser{
		file:        fp,
		scanner:     bufio.NewScanner(fp),
		Command:     "", // 初期値は空
		CommandType: "", // 初期値は空
	}, nil
}

// HasMoreCommands 次の行が存在するか?
// func (p *Parser) HasMoreCommands() bool {
// 	return false
// }

// Next 次の行に進む
// HasMoreCommandsがtrueの場合のみ呼ばれる想定
func (p *Parser) Next() bool {
	// TODO 次のCommandを取得する

	// スキャンする、もし何もなければfalseを返す
	if !p.scanner.Scan() {
		return false
	}

	// TODO 空白とか取り除く
	p.Command = p.scanner.Text()

	// TODO CommandTypeも判別しておく

	return true
}

// Arg1 1つ目の引数を取得する
func (p *Parser) Arg1() string {

	// C_ARTHMETICの場合, add, subなどが返される

	// C_RETURNの場合は、呼ばれないようにする

	return "TODO"
}

// Arg2 2つ目の引数を取得する
func (p *Parser) Arg2() int {

	// C_PUSH, C_POP, C_FUNCTION, C_CALLの場合のみ呼ぶ

	return 0
}

// Close fileをcloseする
func (p *Parser) Close() {
	p.file.Close()
}
