package tokenizer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// symbolList シンボルのリスト
var symbolList []rune = []rune{'{', '}', '(', ')', '.', ',', ';', '+', '-', '*', '/', '&', '|', '<', '>', '=', '~'}

// keywordList キーワードのリスト
var keywordList []string = []string{"class", "constructor", "function", "method", "field", "static", "var", "int", "char", "boolean", "void", "true", "false", "null", "this", "let", "do", "if", "else", "while", "return"}

var symbolMap map[rune]bool
var keywordMap map[string]bool

const (
	// space
	space = ' '
)

// 使用するmapを作成する
func init() {

	// symbol
	symbolMap = map[rune]bool{}
	// symbolmapを作成する
	for _, symbol := range symbolList {
		symbolMap[symbol] = true
	}

	// keyword
	keywordMap = map[string]bool{}
	for _, keyword := range keywordList {
		keywordMap[keyword] = true
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

	var isStringConstMode bool

	var sb strings.Builder
	for _, c := range line {

		// 文字列「"」が来たとき、次の「"」がくるまで文字列として判断する
		if isStringConstMode {
			sb.WriteRune(c)

			// もし"が来た場合は、文字列モード終了
			if c == '"' {
				isStringConstMode = false
			}
			continue
		}

		// 空白の場合
		if c == space {
			// spaceが来たので区切る

			token := sb.String()
			if token != "" {
				tokenList = append(tokenList, token)
			}
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

		// もしダブルクォートが来た場合は文字列モード
		if c == '"' {
			isStringConstMode = true
		}

		// それ以外の単語なので、ただただ追加する
		sb.WriteRune(c)
	}

	return tokenList
}

// GetTokenType トークンタイプを取得する
func GetTokenType(token string) TokenType {

	// 空文字の場合は、error
	if token == "" {
		panic("空文字がtokenとして認識されていました")
	}

	// keywordトークンタイプかどうかを確認する
	if _, ok := keywordMap[token]; ok {
		return TokenTypeKeyWord
	}

	// symbolトークンタイプかどうかを確認する
	if len(token) == 1 {
		// runeにする
		r := rune(token[0])
		if _, ok := symbolMap[r]; ok {
			return TokenTypeSymbol
		}
	}

	// ユニコード文字列のタイプかどうかを確認する
	if strings.HasPrefix(token, "\"") && strings.HasSuffix(token, "\"") {
		return TokenTypeStringConst
	}

	// 数字かどうかを確認
	_, err := strconv.Atoi(token)
	if err == nil {
		// 数字に変換できたので、IntConst
		return TokenTypeIntConst
	}

	// それ以外の場合は変数
	return TokenTypeIdentifier
}

// GetKeyWord キーワードを取得する
func GetKeyWord(token string) KeyWord {

	tokenType := GetTokenType(token)
	if tokenType != TokenTypeKeyWord {
		panic("TokenType: Keyword でないので、KeyWordを取得できません")
	}

	// 予約後を大文字にしたものがkeywordのconstなので、文字をuppserしてそのまま返しています、
	// 決してらくしたいからとかじゃない

	if keywordMap[token] {
		return KeyWord(strings.ToUpper(token))
	}

	panic("存在しないkeyword")
}

// 不要な空白や、コメントを削除する
func trimeLine(line string) string {
	return strings.TrimSpace(strings.Split(line, "//")[0])
}
