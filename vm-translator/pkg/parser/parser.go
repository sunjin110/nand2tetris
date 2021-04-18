package parser

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"vm-translator/pkg/common/chk"
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

// Next 次の行に進む
// 戻りの値がfalseの場合は、何も存在しないということ
func (p *Parser) Next() bool {

	// スキャンする、もし何もなければfalseを返す
	if !p.scanner.Scan() {
		return false
	}

	command := trimeLine(p.scanner.Text())
	if command == "" {
		// もし空白の場合は次
		return p.Next()
	}

	p.Command = command

	// CommandTypeを判別
	p.CommandType = getCommandType(command)

	return true
}

// Arg1 1つ目の引数を取得する
func (p *Parser) Arg1() string {

	// C_RETURNの場合は、呼ばれないようにする
	if p.CommandType == model.CommandTypeReturn {
		chk.SE(errors.New("returnでこのmethodを使うのは想定していない"))
	}

	// 算術の場合は、コマンド自体を返す
	if p.CommandType == model.CommandTypeArithmetic {
		return strings.Split(p.Command, " ")[0]
	}

	return strings.Split(p.Command, " ")[1]
}

// Arg2 2つ目の引数を取得する
func (p *Parser) Arg2() int {

	// C_PUSH, C_POP, C_FUNCTION, C_CALLの場合のみ呼ぶ

	if p.CommandType == model.CommandTypePush ||
		p.CommandType == model.CommandTypePop ||
		p.CommandType == model.CommandTypeCall ||
		p.CommandType == model.CommandTypeFunction {
		seg := strings.Split(p.Command, " ")[2]
		i, err := strconv.Atoi(seg)
		if err != nil {
			panic(err)
		}
		return i
	}

	chk.SE(errors.New("対象のCommandTypeは対応していない"))
	return 0
}

// Close fileをcloseする
func (p *Parser) Close() {
	p.file.Close()
}

// 不要な空白や、コメントを削除する
func trimeLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "//")[0])
}

// コマンドタイプを判別する
func getCommandType(command string) model.CommandType {

	if command == "" {
		panic("コマンドが空です")
	}

	cmd := strings.Split(command, " ")[0]

	commandType, exists := model.Arg1CommandTypeMap[cmd]
	if !exists {
		log.Println("cmd is ", cmd)
		chk.SE(errors.New("定義されていないコマンドがきました"))
	}
	return commandType
}
