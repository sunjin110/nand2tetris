package tokenizer

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// var symbolMap map[rune]bool = map[rune]bool{
// 	'{': true,
// 	'}': true,
// 	'(': true,
// 	''
// }

// const symbolList = [1, 2]
var (
// symbolList = ['{', ]
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
// TODO 文字列に空白が含まれるため、そいつを別Tokenとして扱わないようにする必要がある
func CreateTokenList(line string) []string {
	// return strings.Split(line, " ")

	// 適切に分解していく必要がある

	// 文字を一文字ずつ解析していく?
	log.Println("line is ", line)

	var tokenList []string

	var sb strings.Builder
	for _, c := range line {

		// もし空白の場合は
		log.Println("c is ", string(c), " : ", c)

		switch c {
		case 32: // space
			tokenList = append(tokenList, sb.String())
			sb.Reset() // sbのリセット
		case 59, // ;
			46, // .
			40, // (
			41: // )

			beforeToken := sb.String()
			if beforeToken != "" {
				tokenList = append(tokenList, sb.String())
			}

			tokenList = append(tokenList, string(c))
			sb.Reset()
		// case 46: // .
		// 	tokenList = append(tokenList, sb.String())
		// 	tokenList =
		default:
			sb.WriteRune(c)
		}

	}

	// tokenList = append(tokenList, sb.String())

	return tokenList

}

// 不要な空白や、コメントを削除する
func trimeLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "//")[0])
}
