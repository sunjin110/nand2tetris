// code parserで分解したニーモニックをバイナリコードへ変換する

package code

// destMap ニーモニックを2進数(バイナリコード)に変換
// 読み取り専用
var destMap = map[string]byte{
	"":     0, // 000
	"null": 0, // 000 // nullと記述されても大丈夫なように
	"M":    1, // 001
	"D":    2, // 010
	"MD":   3, // 011
	"A":    4, // 100
	"AM":   5, // 101
	"AD":   6, // 110
	"AMD":  7, // 111
}

// compMap ニーモニックを2進数に変換(バイナリコード)
var compMap = map[string]byte{
	"0":   42,  // 0101010
	"1":   63,  // 0111111
	"-1":  58,  // 0111010
	"D":   12,  // 0001100
	"A":   48,  // 0110000
	"M":   112, // 1110000
	"!D":  13,  // 0001101
	"!A":  49,  // 0110001
	"!M":  113, // 1110001
	"-D":  15,  // 0001111
	"-A":  51,  // 0110011
	"-M":  115, // 1110011
	"D+1": 31,  // 0011111
	"A+1": 55,  // 0110111
	"M+1": 119, // 1110111
	"D-1": 14,  // 0001110
	"A-1": 50,  // 0110010
	"M-1": 114, // 1110010
	"D+A": 2,   // 0000010
	"D+M": 66,  // 1000010
	"D-A": 19,  // 0010011
	"D-M": 83,  // 1010011
	"A-D": 7,   // 0000111
	"M-D": 71,  // 1000111
	"D&A": 0,   // 0000000
	"D&M": 64,  // 1000000
	"D|A": 21,  // 0010101
	"D|M": 85,  // 1010101
}

// jumpMap ニーモニックを2進数に変換(バイナリコード)
// 読み取り専用
var jumpMap = map[string]byte{
	"":     0, // 000
	"null": 0, // 000
	"JGT":  1, // 001
	"JEQ":  2, // 010
	"JGE":  3, // 011
	"JLT":  4, // 100
	"JNE":  5, // 101
	"JLE":  6, // 110
	"JMP":  7, // 111
}

// ConvDest destニーモニックのバイナリコードを返す
func ConvDest(destStr string) byte {
	return destMap[destStr]
}

// ConvComp compニーモニックのバイナリコードを返す
func ConvComp(compStr string) byte {
	return compMap[compStr]
}

// ConvJump jumpニーモニックのバイナリコードを返す
func ConvJump(jumpStr string) byte {
	return jumpMap[jumpStr]
}
