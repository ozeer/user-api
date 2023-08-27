package utils

import (
	"fmt"
	"strconv"
)

func Uint32ToString(num uint32) string {
	return strconv.FormatUint(uint64(num), 10)
}

func StringToUint32(str string) uint32 {
	uint32Value, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return uint32(uint32Value)
}
