package tokenizer

import (
	"bufio"
	"os"
	"strings"
)

// symbolList シンボルのリスト
var symbolList []rune = []rune{'{', '}', '(', ')', '.', ',', ';', '+', '-', '*', '/', '&', '|', '<', '>', '=', '~'}

var symbolMap map[rune]bool

const (
	// space
	space = ' '
)

func init() {

	symbolMap = map[rune]bool{}
	// symbolmapを作成する
	for _, symbol := range symbolList {
		symbolMap[symbol] = true
	}

}

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

	// P233のページで分ける

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
func CreateTokenList(line string) []string {
	// 適切に分解していく必要がある

	// 文字を一文字ずつ解析していく?

	var tokenList []string

	var sb strings.Builder
	for _, c := range line {

		// 空白の場合
		if c == space {
			// spaceが来たので区切る
			tokenList = append(tokenList, sb.String())
			sb.Reset() // StringBuilderのリセット
			continue
		}

		// symbolかどうかを判定する
		if symbolMap[c] {
			// 1つ前の塊を1つ
			beforeToken := sb.String()
			if beforeToken != "" {
				tokenList = append(tokenList, sb.String())
				sb.Reset()
			}

			// 今回のシンボルを1つのtokenとして追加する
			tokenList = append(tokenList, string(c))
			continue
		}

		// それ以外の単語なので、ただただ追加する
		sb.WriteRune(c)
	}

	return tokenList
}

// 不要な空白や、コメントを削除する
func trimeLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "//")[0])
}
