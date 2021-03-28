package common

import (
	"encoding/binary"
	"path/filepath"
	"strconv"
)

// Uint2bytes converts uint64 to []byte
func Uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, i)
	return bytes[8-size : 8]
}

// StrToUint 10進数の文字列を数字に治す
func StrToUint(str string) uint {
	i, _ := strconv.ParseUint(str, 10, 0)
	return uint(i)
}

func GetFileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
