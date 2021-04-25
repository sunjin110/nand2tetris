package tokenizer

import (
	"bufio"
	"os"
	"strings"
)

// Jack言語をparseする機構

// Parser .
type Tokenizer struct {
	file    *os.File       // file(.jack)
	scanner *bufio.Scanner // fileのscanner
	Line    string         // 現在の行
}

// New .
func New(filePath string) (*Tokenizer, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return &Tokenizer{
		file:    fp,
		scanner: bufio.NewScanner(fp),
	}, nil
}

// NextToken 次のTokenに進む
func (t *Tokenizer) NextToken() bool {

	return false
}

// NextLine 次の行に進む
func (t *Tokenizer) NextLine() bool {

	// スキャンする、もし何もなければfalseを返す
	if !t.scanner.Scan() {
		return false
	}

	line := trimeLine(t.scanner.Text())
	if line == "" {
		// もし空白の場合は次
		return t.NextLine()
	}

	t.Line = line
	return true
}

// CreateTokenList ListからTokenリストを取得する
// TODO 文字列に空白が含まれるため、そいつを別Tokenとして扱わないようにする必要がある
func CreateTokenList(line string) []string {
	return strings.Split(line, " ")
}

// 不要な空白や、コメントを削除する
func trimeLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "//")[0])
}
