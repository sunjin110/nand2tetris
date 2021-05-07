package tokenizer

import (
	"bufio"
	"compiler/pkg/common/chk"
	"os"
	"strconv"
	"strings"
)

// symbolList シンボルのリスト
var symbolList []rune = []rune{'{', '}', '(', ')', '[', ']', '.', ',', ';', '+', '-', '*', '/', '&', '|', '<', '>', '=', '~'}

// keywordList キーワードのリスト
var keywordList []string = []string{"class", "constructor", "function", "method", "field", "static", "var", "int",
	"char", "boolean", "void", "true", "false", "null", "this", "let", "do", "if", "else", "while", "return"}

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

// Tokenizer Jack言語をToken単位で分割する機構
type Tokenizer struct {
	file                   *os.File       // file(.jack)
	FileName               string         //
	scanner                *bufio.Scanner // fileのscanner
	line                   string         // 現在の行
	LineNum                int            // 行数番号(エラー時に使用する)
	Token                  string         // 現在のToken
	nowLineTokenList       []string       // 現在の行のtokenList
	nowTokenLineIndex      int            // 現在のtokenが現在の行の何個目か？
	isMultiLineCommentMode bool           // /* */ というコメント内かどうかを判定するもの
}

// New .
func New(filePath string) (*Tokenizer, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return &Tokenizer{
		file:     fp,
		FileName: filePath,
		scanner:  bufio.NewScanner(fp),
	}, nil
}

// NextToken 次のTokenに進む
func (t *Tokenizer) NextToken() bool {

	nextTokenIdx := t.nowTokenLineIndex + 1

	// もしLineが空の場合は、NextLineを呼ぶ
	// また、現在のTokenの長さよりも、呼ぼうとしているidxの長さがあふれる場合は、次
	// この時点でfalseの場合は、次のtokenは存在しない
	if t.line == "" || nextTokenIdx >= len(t.nowLineTokenList) {
		if !t.nextLine() {
			return false
		}

		// lineが取れたので、そこからtokenのリストを生成する
		tokenList := t.createTokenList(t.line)
		if len(tokenList) == 0 {
			return t.NextToken()
		}

		t.nowLineTokenList = tokenList
		t.nowTokenLineIndex = 0

		// 行の初回なので、先頭をそのままTokenにセットして終わる
		t.Token = tokenList[0]
		return true
	}

	// 次のリストのTokenを取得する
	t.Token = t.nowLineTokenList[nextTokenIdx]
	t.nowTokenLineIndex = nextTokenIdx
	return true
}

// GetSymbol 現在のTokenがsymbolの場合、どのsymbolかを習得する
// tokenTypeがsymbol以外の場合はError
func (t *Tokenizer) GetSymbol() rune {

	tokenType := GetTokenType(t.Token)
	if tokenType != TokenTypeSymbol {
		panic("GetSymbolはTokenTypeがSymbol以外の場合は取得できません")
	}

	return rune(t.Token[0])
}

// GetIdentifier 現在のTokenがidentifierの場合の、変数名を習得する
func (t *Tokenizer) GetIdentifier() string {

	tokenType := GetTokenType(t.Token)
	if tokenType != TokenTypeIdentifier {
		panic("GetIdentifierはTokenTypeがIdentifier以外の場合は取得できません")
	}

	return t.Token
}

// GetIntVal 現在のTokenがIntConstの場合の値を習得する
func (t *Tokenizer) GetIntVal() int {

	tokenType := GetTokenType(t.Token)
	if tokenType != TokenTypeIntConst {
		panic("GetIntValはTokenTypeがIntConst以外の場合は取得できません")
	}

	i, err := strconv.Atoi(t.Token)
	chk.SE(err)
	return i
}

// GetStringVal 現在のTokenがStringConstの場合の値を習得する
func (t *Tokenizer) GetStringVal() string {

	tokenType := GetTokenType(t.Token)
	if tokenType != TokenTypeStringConst {
		panic("GetStringValはTokenTypeがStringConst以外の場合は取得できません")
	}

	// 両端を取り除く
	return t.Token[1 : len(t.Token)-1]
}

// nextLine 次の行に進む
func (t *Tokenizer) nextLine() bool {

	t.LineNum += 1

	// スキャンする、もし何もなければfalseを返す
	if !t.scanner.Scan() {
		return false
	}

	line := trimeLine(t.scanner.Text())
	if line == "" {
		// もし空白の場合は次
		return t.nextLine()
	}

	t.line = line
	return true
}

// createTokenList ListからTokenリストを取得する
func (t *Tokenizer) createTokenList(line string) []string {

	// 文字を一文字ずつ解析していく
	var tokenList []string

	// 文字列「"」で包まれている間は、空白やその他のsymbolが来ても無効化する必要がある
	var isStringConstMode bool

	var sb strings.Builder

	for idx, c := range line {

		// もし複数行コメントモードの場合、*or/以外は受け付けない
		if t.isMultiLineCommentMode {

			// もし「/」が来た場合、一つ前の単語が「*」かどうかを確認もしそうなら
			// 複数行コメントを終了する
			if c == '/' && idx > 0 {
				beforeC := rune(line[idx-1])
				if beforeC == '*' {
					// 「*/」が検出されたので、複数行コメントモードを解除する
					t.isMultiLineCommentMode = false
				}
			}
			continue
		}

		// 文字列「"」が来たとき、次の「"」がくるまで文字列として判断する
		// 複数行コメントモードではないことを前提とする /* */
		if isStringConstMode {
			sb.WriteRune(c)

			// もし"が来た場合は、文字列モード終了
			if c == '"' {
				isStringConstMode = false
			}
			continue
		}

		// もし、「/」が来たときにその次に「*」が来ているかどうかを先読みする
		// 「/」がその一列の最後の文字の場合は検証する必要はない
		// もし「*」が来る場合は、
		if c == '/' && len(line) > idx-1 {
			nextC := rune(line[idx+1])

			if nextC == '*' {
				// 「/*」が検出されたので、複数行コメントモードに入る
				t.isMultiLineCommentMode = true

				// 入る前にsbの中身がある場合は追加する
				beforeToken := sb.String()
				if beforeToken != "" {
					tokenList = append(tokenList, beforeToken)
					sb.Reset()
				}

				continue
			}
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
