package vmwriter

// JackOSで標準で使用できるFunctionなどの定義

// const (
// 	// MathMultiply .
// 	MathMultiply = "Math.multiply"
// )

// getMathMultiply 掛け算の関数
func getMathMultiply() (string, int32) {
	return "Math.multiply", 2
}

// getMathDivide 引き算の関数
func getMathDivide() (string, int32) {
	return "Math.divide", 2
}

// getMemoryAlloc memory allocate
func getMemoryAlloc() (string, int32) {
	return "Memory.alloc", 1
}
