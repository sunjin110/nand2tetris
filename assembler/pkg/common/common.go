package common

import (
	"encoding/binary"
	"strconv"
)

// Uint2bytes converts uint64 to []byte
func Uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, i)
	return bytes[8-size : 8]
}

func StrToUint(str string) uint {
	i, _ := strconv.ParseUint(str, 8, 0)
	return uint(i)
}
